package models

import (
	"fmt"
	"time"

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

type GetTeacherByIDWithoutPassword struct {
	TeacherID        int64     `json:"teacher_id"`
	UserID           int64     `json:"user_id"`
	UserName         string    `json:"name" binding:"required"`
	UserGender       string    `json:"gender" binding:"required,oneof=Male Female"`
	UserPlaceOfBirth string    `json:"place_of_birth" binding:"required"`
	UserDateOfBirth  time.Time `json:"date_of_birth"`
	UserReligion     string    `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	UserAddress      string    `json:"address" binding:"required" validate:"required,min=10,max=200"`
	UserNumPhone     string    `json:"num_phone" binding:"required,e164"`
	UserEmail        string    `json:"email" binding:"required,email"`
	TeachingHour     int32     `json:"teaching_hour"`
}

// Get teacher by id
func (teacher *TeacherModel) GetTeacherById(id string) (GetTeacherByIDWithoutPassword, error) {

	result := connections.DB.Preload("User").First(&teacher, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return GetTeacherByIDWithoutPassword{}, fmt.Errorf("no teacher found")
		}
		return GetTeacherByIDWithoutPassword{}, result.Error
	}

	var teacherDTO GetTeacherByIDWithoutPassword

	teacherDTO.TeacherID = teacher.TeacherID
	teacherDTO.UserID = teacher.UserID
	teacherDTO.UserName = teacher.User.UserName
	teacherDTO.UserGender = teacher.User.UserGender
	teacherDTO.UserPlaceOfBirth = teacher.User.UserPlaceOfBirth
	teacherDTO.UserDateOfBirth = teacher.User.UserDateOfBirth
	teacherDTO.UserReligion = teacher.User.UserReligion
	teacherDTO.UserAddress = teacher.User.UserAddress
	teacherDTO.UserNumPhone = teacher.User.UserNumPhone
	teacherDTO.UserEmail = teacher.User.UserEmail
	teacherDTO.TeachingHour = teacher.TeachingHour

	return teacherDTO, nil
}

// Get teacher by user id
func (teacher *Teacher) GetTeacherByModel() error {
	result := connections.DB.Where(&teacher).First(&teacher)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("teacher not found")
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
	result := connections.DB.Delete(&teacher, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
