package helper

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/models"
)

func Authenticate(email string, password string) (models.User, string) {
	// Find user by email using GORM
	var user = models.User{
		UserEmail: email,
	}
	user, err := user.GetUser();

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
	var teachers []models.Teacher
	connections.DB.Where(models.Teacher{
		UserID: user.UserID,
	}).Find(&teachers)

	if len(teachers) > 0 {
		return "teacher"
	}

	var staff []models.Staff
	connections.DB.Where(models.Staff{
		UserID: user.UserID,
	}).First(&staff)

	if len(staff) > 0 {
		return "staff"
	}

	var admins []models.Admin
	connections.DB.Where(models.Admin{
		UserID: user.UserID,
	}).First(&admins)

	if len(admins) > 0 {
		return "admin"
	}

	return ""
}
