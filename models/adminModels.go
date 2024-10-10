package models

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration/lib"
)

type Admin struct {
	AdminID  int64  `gorm:"primaryKey" json:"id"`
	UserID   int64  `json:"id_user" binding:"required"`
	Position string `json:"teacher_position" binding:"required"`
	lib.BaseModel
}

func (Admin) TableName() string {
	return lib.GenerateTableName(lib.Public, "admins")
}

// get admin by id
func (admin *Admin) GetAdminByUserID() error {
	// result := connections.DB.Where("user_id = ?", id).First(&admin)
	result := connections.DB.Where(&admin).First(&admin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
