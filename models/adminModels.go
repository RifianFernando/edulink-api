package models

import (
	"github.com/skripsi-be/database/migration/lib"
)

type Admin struct {
	AdminID  uint   `gorm:"primaryKey" json:"id"`
	UserID   int64 `json:"id_user" binding:"required"`
	Position string `json:"teacher_position" binding:"required"`
	lib.BaseModel
}

func (Admin) TableName() string {
	return lib.GenerateTableName(lib.Public, "admins")
}
