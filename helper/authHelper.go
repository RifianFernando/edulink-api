package helper

import (
	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/models"
	"github.com/edulink-api/lib"
)

func Authenticate(email string, password string) (models.User, []string) {
	// Find user by email using GORM
	var user = models.User{
		UserEmail: email,
	}
	user, err := user.GetUser()

	if err != nil {
		return models.User{}, []string{}
	}

	// Check password using bcrypt
	err = lib.CompareHash(user.UserPassword, password)
	if err != nil {
		return models.User{}, []string{}
	}

	userType := GetUserTypeByUID(user)
	if len(userType) == 0 {
		return models.User{}, []string{}
	}

	return user, userType
}

func GetUserTypeByUID(user models.User) []string {
	// find admin first rather than staff and teacher because admin has the highest priority and if the user is both admin and staff or teacher, the user will be considered as admin
	var admins models.Admin
	connections.DB.Where(models.Admin{
		UserID: user.UserID,
	}).First(&admins)

	var roles []string

	if admins.AdminID != 0 {
		roles = append(roles, "admin")
	}

	var staff models.Staff
	connections.DB.Where(models.Staff{
		UserID: user.UserID,
	}).First(&staff)

	if staff.StaffId != 0 {
		roles = append(roles, "staff")
	}

	var teachers models.Teacher
	connections.DB.Where(models.Teacher{
		UserID: user.UserID,
	}).Find(&teachers)

	if teachers.TeacherID != 0 {
		roles = append(roles, "teacher")
	}

	// check if the user is a homeroom teacher
	var className models.ClassName
	className.TeacherID = teachers.TeacherID
	classes, err := className.GetHomeRoomTeacherByTeacherID()
	if err == nil && len(classes) > 0 {
		// if the user is a homeroom teacher, add the role to the roles and remove the teacher role
		for _, className := range roles {
			if className == "teacher" {
				roles = roles[:len(roles)-1]
				break
			}
		}
		roles = append(roles, "homeroom_teacher")
	}

	return roles
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

func GetUserTypeByPrivilege(user models.User) string {

	var userType string
	for _, userType = range GetUserTypeByUID(user) {
		if userType == "admin" {
			userType = "admin"
			break
		} else if userType == "staff" {
			userType = "staff"
			break
		} else if userType == "homeroom_teacher" {
			userType = "homeroom_teacher"
			break
		} else if userType == "teacher" {
			userType = "teacher"
			break
		}
	}

	return userType
}
