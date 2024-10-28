package helper

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/edulink-api/lib"
	"github.com/edulink-api/models"
	"github.com/golang-jwt/jwt"
)

type userDetailToken struct {
	UserID    int64
	UserName  string
	User_type string
	TokenType string
	jwt.StandardClaims
}

type resetPassDetailToken struct {
	UserID    int64
	UserEmail string
	TokenType string
	jwt.StandardClaims
}

var invalidToken = "the token is invalid"

var SECRET_KEY string = os.Getenv("APP_KEY")

func CustomTimeDay(days int) time.Time {
	return time.Now().Local().Add(time.Hour * time.Duration(24*days))
}

func GenerateToken(user models.User, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &userDetailToken{
		UserID:    user.UserID,
		UserName:  user.UserName,
		User_type: userType,
		TokenType: "access_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: CustomTimeDay(1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshClaims := &userDetailToken{
		UserID:    user.UserID,
		UserName:  user.UserName,
		User_type: userType,
		TokenType: "refresh_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: CustomTimeDay(7).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, nil
}

func GenerateResetPasswordToken(
	userID int64,
	email string,
) (string, error) {
	claims := &resetPassDetailToken{
		UserID:    userID,
		UserEmail: email,
		TokenType: "reset_password",
		StandardClaims: jwt.StandardClaims{
			// 5 minutes
			ExpiresAt: time.Now().Local().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	resetToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return resetToken, nil
}

func UpdateSession(refreshToken string, userID int64, ipAddress string, userAgent string) (newToken string, newRefreshToken string, err error) {
	// Validate the refresh token
	claims, msg := ValidateToken(refreshToken, "refresh_token")
	if msg != "" || claims.UserID != userID {
		return "", "", errors.New("invalid token")
	}

	// Retrieve the session from the database
	var session = models.Session{UserID: userID}
	session = session.GetSession()
	if session.SessionID == 0 {
		return "", "", errors.New("the refresh token does not exist")
	}

	// Generate new access token and refresh token
	user := models.User{UserID: userID, UserName: claims.UserName}
	newToken, newRefreshToken, err = GenerateToken(user, GetUserTypeByUID(user))
	if err != nil {
		return "", "", err
	}

	// Update session with new refresh token
	session.RefreshToken = lib.HashToken(newRefreshToken)
	session.IPAddress = ipAddress
	session.UserAgent = userAgent
	session.ExpiresAt = CustomTimeDay(7) // Extend session expiry

	err = session.UpdateSession()
	if err != nil {
		return "", "", err
	}

	return newToken, newRefreshToken, nil
}

func InsertSession(
	refreshToken string,
	userID int64,
	ipAddress string,
	userAgent string,
) string {
	session := models.Session{
		UserID: userID,
	}

	// Check if the session already exists for the user
	exists, err := session.SessionExists()
	if err != nil {
		return err.Error()
	}

	if exists {
		return "the refresh token already exists"
	}

	// Insert new session
	session.RefreshToken = lib.HashToken(refreshToken)
	session.IPAddress = ipAddress
	session.UserAgent = userAgent
	session.ExpiresAt = CustomTimeDay(7) // Set session expiry to match refresh token

	err = session.InsertSession()

	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}

func ValidateToken(
	signedToken string,
	tokenType string,
) (
	claims *userDetailToken,
	msg string,
) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&userDetailToken{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()

		return
	}

	claims, ok := token.Claims.(*userDetailToken)
	if !ok {
		msg = invalidToken

		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "The token is expired"

		return
	}

	if claims.TokenType != tokenType {
		msg = invalidToken

		return
	}

	return claims, msg
}

func ValidateRefreshToken(
	signedToken string,
) (
	claims *userDetailToken,
	msg string,
) {
	claims, msg = ValidateToken(signedToken, "refresh_token")
	if msg != "" {
		return nil, msg
	}

	// check database
	var session = models.Session{
		UserID: claims.UserID,
	}
	session = session.GetSession()
	if session.SessionID == 0 {
		return nil, "the refresh token does not exist"
	}

	// Validate the stored refresh token with the one passed
	if !lib.VerifyToken(signedToken, session.RefreshToken) {
		return nil, "the refresh token is invalid"
	}

	return claims, msg
}

func ValidateResetPasswordToken(
	signedToken string,
) (
	claims *resetPassDetailToken,
	msg string,
) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&resetPassDetailToken{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()

		return
	}

	claims, ok := token.Claims.(*resetPassDetailToken)
	if !ok {
		msg = invalidToken

		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "the token is expired"

		return
	}

	if claims.TokenType != "reset_password" {
		msg = invalidToken

		return
	}

	return claims, msg
}

func DeleteToken(
	accessToken string,
	refreshToken string,
) (
	isDeleted bool,
	msg string,
) {
	// validate the token
	claims, msgAccess := ValidateToken(accessToken, "access_token")
	claimsRefresh, msgRefresh := ValidateRefreshToken(refreshToken)
	if msgAccess != "" || msgRefresh != "" {
		msg = invalidToken

		return false, msg
	}

	if (claims.UserID != claimsRefresh.UserID) || (claims.UserName != claimsRefresh.UserName) {
		msg = "access token and refresh token are not match"

		return false, msg
	}

	var session = models.Session{UserID: claims.UserID}
	session = session.GetSession()

	if session.SessionID == 0 {
		msg = invalidToken
		return false, msg
	}

	if !lib.VerifyToken(refreshToken, session.RefreshToken) {
		msg = invalidToken

		return false, msg
	}

	// bulk delete
	err := session.DeleteSession()
	if err != nil {
		msg = err.Error()

		return false, msg
	}

	return true, msg
}

func GetAccessTokenFromHeader(
	authHeader string,
) (
	accessToken string,
	msg string,
) {
	if authHeader == "" {
		msg = "Authorization header is empty"

		return "", msg
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		msg = "Invalid Authorization header format"

		return "", msg
	}

	accessToken = parts[1]

	return accessToken, msg
}
