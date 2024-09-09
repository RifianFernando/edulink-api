package models

import (
	"github.com/skripsi-be/database/migration/lib"
)

type Teacher struct {
	TeacherID    int64 `gorm:"primaryKey"`
	UserID       int64 `json:"id_user" binding:"required"`
	TeachingHour int32 `json:"teaching_hour" binding:"required"`
	lib.BaseModel
}

func (Teacher) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teachers")
}
