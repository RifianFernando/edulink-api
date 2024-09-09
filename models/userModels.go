package models

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

type User struct {
	UserID           uint      `gorm:"primaryKey" json:"id"`
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
