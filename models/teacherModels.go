package models

import (
	"github.com/skripsi-be/database/migration/lib"
)

type Teacher struct {
	TeacherID    int64       `gorm:"primaryKey"`
	UserID       int64       `json:"id_user" binding:"required"`
	TeachingHour int32       `json:"teaching_hour" binding:"required"`
	ClassNames   []ClassName `gorm:"foreignKey:TeacherID;references:TeacherID"`
	User         User        `gorm:"foreignKey:UserID;references:UserID"` // Belongs-to with User
	// Scores       []Score     `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	lib.BaseModel
}

func (Teacher) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teachers")
}
