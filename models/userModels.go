package models

import (
	"time"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration/lib"
)

type User struct {
	UserID           int64     `gorm:"primaryKey" json:"id"`
	UserName         string    `json:"name" binding:"required"`
	UserGender       string    `json:"gender" binding:"required"`
	UserPlaceOfBirth string    `json:"place_of_birth" binding:"required"`
	UserDateOfBirth  time.Time `json:"date_of_birth"`
	UserAddress      string    `json:"address" binding:"required"`
	UserNumPhone     string    `json:"num_phone" binding:"required"`
	UserEmail        string    `json:"email" binding:"required"`
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
