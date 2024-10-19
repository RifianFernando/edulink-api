package models

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration/lib"
	"gorm.io/gorm"
)

type ClassName struct {
	ClassNameID    int64  `gorm:"primaryKey" json:"id"`
	TeacherID      int64  `json:"id_teacher" binding:"required"`
	ClassName      string `json:"name" binding:"required"`
	ClassNameGrade string `json:"grade" binding:"required"`
	lib.BaseModel
}

func (ClassName) TableName() string {
	return lib.GenerateTableName(lib.Public, "class_names")
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
	classNames []ClassName,
	msg string,
) {
	result := connections.DB.Find(&classNames)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No class found"
	}

	return classNames, ""
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
