package models

import (
	"github.com/skripsi-be/database/migration/lib"
)

type Staff struct {
	StaffId  int64  `gorm:"primaryKey"`
	UserID   int64  `json:"id_user" binding:"required"`
	Position string `json:"user_position" binding:"required"`
	lib.BaseModel
}

func (Staff) TableName() string {
	return lib.GenerateTableName(lib.Administration, "staffs")
}
