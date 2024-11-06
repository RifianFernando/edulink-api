package helper

import (
	"github.com/edulink-api/connections"
	"github.com/edulink-api/lib"
	"github.com/edulink-api/models"
)

func Authenticate(email string, password string) (models.User, string) {
	// Find user by email using GORM
	var user = models.User{
		UserEmail: email,
	}
	user, err := user.GetUser()

	if err != nil {
		return models.User{}, ""
	}

	// Check password using bcrypt
	err = lib.CompareHash(user.UserPassword, password)
	if err != nil {
		return models.User{}, ""
	}

	userType := GetUserTypeByUID(user)
	if userType == "" {
		return models.User{}, ""
	}

	return user, userType
}

func GetUserTypeByUID(user models.User) string {
	// find admin first rather than staff and teacher because admin has the highest priority and if the user is both admin and staff or teacher, the user will be considered as admin
	var admins []models.Admin
	connections.DB.Where(models.Admin{
		UserID: user.UserID,
	}).First(&admins)

	if len(admins) > 0 {
		return "admin"
	}

	var staff []models.Staff
	connections.DB.Where(models.Staff{
		UserID: user.UserID,
	}).First(&staff)

	if len(staff) > 0 {
		return "staff"
	}

	var teachers []models.Teacher
	connections.DB.Where(models.Teacher{
		UserID: user.UserID,
	}).Find(&teachers)

	if len(teachers) > 0 {
		return "teacher"
	}

	return ""
}

func GetUserByEmail(email string) (models.User, error) {
	var user = models.User{
		UserEmail: email,
	}
	user, err := user.GetUser()

	return user, err
}

func UpdateUserPassword(user models.User) error {
	user.UserPassword = lib.HashPassword(user.UserPassword)
	err := user.UpdatePassword()
	if err != nil {
		return err
	}

	return err
}
