package models

import (
	"fmt"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type ClassName struct {
	ClassNameID int64          `gorm:"primaryKey" json:"id"`
	GradeID     int64          `json:"id_grade" binding:"required"`
	TeacherID   int64          `json:"id_teacher" binding:"required"`
	Name        string         `json:"name" binding:"required" validate:"len=1"`
	// Manually managed by gorm because I can't find a way to automatically manage it
	CreatedAt   time.Time      // Automatically managed by GORM for creation timecreating time
	UpdatedAt   time.Time      // Automatically managed by GORM for update time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ClassNameModel struct {
	ClassName
	Teacher Teacher `gorm:"foreignKey:TeacherID;references:TeacherID"`
	Grade   Grade   `gorm:"foreignKey:GradeID;references:GradeID"`
}

type ClassNameGrade struct {
	ClassName
	Grade Grade `gorm:"foreignKey:GradeID;references:GradeID"`
}

func (ClassName) TableName() string {
	return lib.GenerateTableName(lib.Academic, "class_names")
}

// Create class
func (className *ClassName) CreateClassName() error {
	if err := connections.DB.Create(&className).Error; err != nil {
		return err
	}

	return nil
}

// Get all ClassName
func (class *ClassNameModel) GetAllClassName() (
	className []ClassNameModel,
	msg string,
) {
	result := connections.DB.Preload("Teacher").Preload("Grade").Find(&className)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No class found"
	}

	if class.Teacher.UserID != 0 {
		var classList []ClassNameModel
		for _, v := range className {
			if v.Teacher.UserID == class.Teacher.UserID {
				classList = append(classList, v)
			}
		}
		// return classess by teacher id
		return classList, ""
	}

	// return all classess
	return className, ""
}

// get class by id
func (className *ClassName) GetClassNameById(id string) (ClassName, error) {
	var classNames ClassName
	result := connections.DB.First(&classNames, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return classNames, nil
		}
		return classNames, result.Error
	}

	return classNames, nil
}

// get class Name by models
func (className *ClassNameModel) GetClassNameModelByID(classNameID string) error {
	result := connections.DB.Where("class_name_id = ?", classNameID).Preload("Teacher").Preload("Grade").First(&className)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// update class by id
func (class *ClassName) UpdateClassNameByObject() error {
	result := connections.DB.Model(&class).Updates(&class)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// delete class by id
func (class *ClassName) DeleteClassNameById(id string) error {
	result := connections.DB.Delete(&class, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (className *ClassName) GetHomeRoomTeacherByTeacherID() (classes []ClassName, err error) {
	result := connections.DB.Where("teacher_id = ?", className.TeacherID).Find(&classes)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound || len(classes) == 0 {
		if result.Error == gorm.ErrRecordNotFound {
			return []ClassName{}, fmt.Errorf("no class found")
		}
		return []ClassName{}, result.Error
	}

	return classes, nil
}
