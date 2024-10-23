package models

import (
	"github.com/skripsi-be/connections"
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

// Create teacher

func (teacher *Teacher) CreateTeacher() error {
	result := connections.DB.Create(&teacher)
	if result.Error != nil {
			return result.Error
	}
	return nil
}

func (teacher *Teacher) GetAllUserTeachers() (
	teachers []Teacher,
	msg string,
) {
	result := connections.DB.Preload("User").Find(&teachers)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No user teacher found"
	}

	return teachers, ""
}
