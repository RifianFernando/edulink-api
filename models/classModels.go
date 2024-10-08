package models

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration/lib"
	"gorm.io/gorm"
)

type Class struct {
	ClassID    int64  `gorm:"primaryKey" json:"id"`
	TeacherID  int64  `json:"id_teacher" binding:"required"`
	ClassName  string `json:"name" binding:"required"`
	ClassGrade string `json:"grade" binding:"required"`
	lib.BaseModel
}

func (Class) TableName() string {
	return lib.GenerateTableName(lib.Public, "classes")
}

// Create class
func (class *Class) CreateClass() error {
	if err := connections.DB.Create(&class).Error; err != nil {
		return err
	}

	return nil
}

// Get all classes
func (class *Class) GetAllClass() (
	classes [] Class,
	msg string,
) {
	result := connections.DB.Find(&classes)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No class found"
	}

	return classes, ""
}

// get class by id
func (class *Class) GetClassById(id string) (Class, error) {
	var classes Class
	result := connections.DB.First(&classes, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return classes, nil
		}
		return classes, result.Error
	}

	return classes, nil
}

// update class by id
func (class *Class) UpdateClassByObject() error {
	result := connections.DB.Model(&class).Updates(&class)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// delete class by id
func (class *Class) DeleteClassById(id string) error {
	result := connections.DB.Delete(&class, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
