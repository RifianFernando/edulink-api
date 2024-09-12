package controllers

import (
	"errors"
	"fmt"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/models"
	"gorm.io/gorm"
)

func Authenticate(email string, password string) (int, error) {
	var user models.User

	// Find user by email using GORM
	result := connections.DB.Where("user_email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("invalid credentials")
		}
		return 0, result.Error // Return other potential DB errors
	}

	// Check password using bcrypt
	err := lib.CompareHash(user.UserPassword, password)
	if err != nil {
		return 0, fmt.Errorf("invalid credentials")
	}

	return int(user.UserID), nil // Return the user ID on successful authentication
}
