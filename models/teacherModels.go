package models

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type Teacher struct {
	TeacherID    int64 `gorm:"primaryKey"`
	UserID       int64 `json:"id_user" binding:"required"`
	TeachingHour int32 `json:"teaching_hour" binding:"required"`
	lib.BaseModel
}

type TeacherModel struct {
	Teacher
	ClassNames []ClassName `gorm:"foreignKey:TeacherID;references:TeacherID"`
	User       User        `gorm:"foreignKey:UserID;references:UserID"` // Belongs-to with User
	// Scores       []Score     `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
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

// Get all teacher
func (teacher *TeacherModel) GetAllUserTeachersWithUser() (
	teachers []TeacherModel,
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

// Get teacher by id
func (teacher *TeacherModel) GetTeacherById(id string) (TeacherModel, error) {
	var teachers TeacherModel
	result := connections.DB.Preload("User").First(&teachers, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return teachers, nil
		}
		return teachers, result.Error
	}

	return teachers, nil
}

// Get teacher by user id
func (teacher *Teacher) GetTeacherByModel() error {
	var teachers Teacher
	result := connections.DB.First(&teachers)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}

	return nil
}

// Update teacher and user
func (teacher *TeacherModel) UpdateTeacherById(teacherData *TeacherModel) error {
	// 1. Update teacher fields (excluding User and BaseModel)
	result := connections.DB.Model(&teacher.Teacher).Updates(
		&Teacher{
			TeachingHour: teacherData.TeachingHour,
		},
	)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("teacher not found")
	}

	// 2. Update associated user
	userData := teacherData.User
	result = connections.DB.Model(&teacher.User).Updates(
		&User{
			UserName:         userData.UserName,
			UserGender:       userData.UserGender,
			UserPlaceOfBirth: userData.UserPlaceOfBirth,
			UserDateOfBirth:  userData.UserDateOfBirth,
			UserAddress:      userData.UserAddress,
			UserNumPhone:     userData.UserNumPhone,
			UserReligion:     userData.UserReligion,
			UserEmail:        userData.UserEmail,
		},
	)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Delete teacher by id
func (teacher *Teacher) DeleteTeacherById(id string) error {
	result := connections.DB.Unscoped().Delete(&teacher, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
