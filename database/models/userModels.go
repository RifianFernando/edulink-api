package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
)

type User struct {
	UserID           int64     `gorm:"primaryKey" json:"id"`
	UserName         string    `json:"name" binding:"required"`
	UserGender       string    `json:"gender" binding:"required,oneof=Male Female"`
	UserPlaceOfBirth string    `json:"place_of_birth" binding:"required"`
	UserDateOfBirth  time.Time `json:"date_of_birth"`
	UserReligion     string    `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	UserAddress      string    `json:"address" binding:"required" validate:"required,min=10,max=200"`
	UserPhoneNum     string    `json:"num_phone" binding:"required,e164"`
	UserEmail        string    `json:"email" binding:"required,email"`
	UserPassword     string    `json:"password" binding:"required"`
	lib.BaseModel
}

func (User) TableName() string {
	return lib.GenerateTableName(lib.Public, "users")
}

func (user *User) GetUser() (User, error) {
	// get result by model
	result := connections.DB.Where(&user).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return *user, nil
}

// Create user
func (user *User) CreateUser() error {
	result := connections.DB.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return fmt.Errorf("user already exist")
		}
		return result.Error
	}
	return nil
}

func (user *User) UpdatePassword() error {
	result := connections.DB.Model(&user).Update("user_password", user.UserPassword)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Delete user
func (user *User) DeleteUserByID(id string) error {
	result := connections.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get User Profile
// TODO: Add user type
func (user *User) GetUserProfile(userType []string) (User, string) {
	result := connections.DB.Where(&user).First(&user)
	if result.Error != nil {
		return User{}, result.Error.Error()
	}

	return *user, ""
}
