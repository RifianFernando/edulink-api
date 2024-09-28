package helper

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/models"
)

type userDetailToken struct {
	UserID    int64
	UserName  string
	User_type string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SESSION_KEY")

func CustomTimeDay(days int) time.Time {
	return time.Now().Local().Add(time.Hour * time.Duration(24*days))
}

func GenerateToken(user models.User, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &userDetailToken{
		UserID:    user.UserID,
		UserName:  user.UserName,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: CustomTimeDay(1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshClaims := &userDetailToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: CustomTimeDay(2).Unix(),
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

func UpdateSessionTable(
	token string,
	refreshToken string,
	userID int64,
	IpAddress string,
	UserAgent string,
) error {
	var sessions models.Session
	// result := connections.DB.First(&session, userID)
	result := connections.DB.Where(models.Session{
		UserID: userID,
	}).Find(&sessions)

	token = lib.HashToken(token)
	refreshToken = lib.HashToken(refreshToken)
	if result.RowsAffected == 0 {
		session := models.Session{
			UserID:       userID,
			SessionToken: token,
			RefreshToken: refreshToken,
			IPAddress:    IpAddress,
			UserAgent:    UserAgent,
			ExpiresAt:    CustomTimeDay(1),
		}
		// Save the session to the database
		if err := connections.DB.Create(&session).Error; err != nil {
			return err
		}
		return nil
	} else {
		// Update the session in the database using where
		if err := connections.DB.Model(&sessions).Where("user_id = ?", userID).Updates(models.Session{
			SessionToken: token,
			RefreshToken: refreshToken,
			IPAddress:    IpAddress,
			UserAgent:    UserAgent,
			ExpiresAt:    CustomTimeDay(1),
		}).Error; err != nil {
			return err
		}

		return nil
	}
}

func ValidateToken(
	signedToken string,
) (
	claims *userDetailToken,
	msg string,
) {
	var invalidToken = "The token is invalid"

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

	var session models.Session
	if err := connections.DB.Where(models.Session{
		UserID: claims.UserID,
	}).First(&session).Error; err != nil {
		msg = invalidToken

		return
	}

	if !lib.VerifyToken(signedToken, session.SessionToken) {
		msg = invalidToken

		return
	}

	return claims, msg
}

func DeleteToken(
	signedToken string,
) (
	isDeleted bool,
	msg string,
) {
	var invalidToken = "The token is invalid"

	token, err := jwt.ParseWithClaims(
		signedToken,
		&userDetailToken{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()

		return false, msg
	}

	claims, ok := token.Claims.(*userDetailToken)
	if !ok {
		msg = invalidToken

		return false, msg
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "The token is expired"

		return false, msg
	}

	var session models.Session
	if err := connections.DB.Where(models.Session{
		UserID: claims.UserID,
	}).First(&session).Error; err != nil {
		msg = invalidToken

		return false, msg
	}

	if !lib.VerifyToken(signedToken, session.SessionToken) {
		msg = invalidToken

		return false, msg
	}

	// bulk delete
	if err := connections.DB.Unscoped().Delete(&session).Error; err != nil {
		msg = err.Error()

		return false, msg
	}

	return true, msg
}
