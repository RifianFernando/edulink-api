package models

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration/lib"
	"gorm.io/gorm"
)

type ClassName struct {
	ClassNameID int64   `gorm:"primaryKey" json:"id"`
	GradeID     int64   `json:"id_grade" binding:"required"`
	TeacherID   int64   `json:"id_teacher" binding:"required"`
	Name        string  `json:"name" binding:"required" validate:"len=1"`
	Teacher     Teacher `gorm:"foreignKey:TeacherID;references:TeacherID"`
	Grade       Grade   `gorm:"foreignKey:GradeID;references:GradeID"`
	lib.BaseModel
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
func (class *ClassName) GetAllClassName() (
	className []ClassName,
	msg string,
) {
	result := connections.DB.Preload("Teacher.User").Preload("Grade").Find(&className)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No class found"
	}

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
