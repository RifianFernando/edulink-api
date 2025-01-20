package models

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type TeacherSubject struct {
	TeacherSubjectID     int64                  `gorm:"primaryKey" json:"id"`
	TeacherID            int64                  `json:"teacher_id" binding:"required"`
	SubjectID            int64                  `json:"subject_id" binding:"required"`
	Subject              Subject                `gorm:"foreignKey:SubjectID;references:SubjectID"`
	TeachingClassSubject []TeachingClassSubject `gorm:"foreignKey:TeacherSubjectID;references:TeacherSubjectID"`
	lib.BaseModel
}

func (TeacherSubject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teacher_subjects")
}

func CreateTeacherSubject(teacherSubject []TeacherSubject) error {
	result := connections.DB.Create(&teacherSubject)
	if result.Error != nil {
		if result.Error == gorm.ErrInvalidData {
			return fmt.Errorf("invalid data")
		} else if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("data not found")
		}
		return result.Error
	}

	return nil
}

type TeacherSubjectGrade struct {
	TeacherSubjectID     int64                  `gorm:"primaryKey" json:"id"`
	TeacherID            int64                  `json:"teacher_id" binding:"required"`
	SubjectID            int64                  `json:"subject_id" binding:"required"`
	TeachingClassSubject []TeachingClassSubject `gorm:"foreignKey:TeacherSubjectID;references:TeacherSubjectID"`
	lib.BaseModel
}

func (TeacherSubjectGrade) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teacher_subjects")
}

func (
	teacherSubject *TeacherSubjectGrade,
) GetTeachingSubjectByID() (teacherSubjects []TeacherSubjectGrade, err error) {
	result := connections.DB.Preload("Subject.Grade").Preload("TeachingClassSubject").Find(&teacherSubjects, "teacher_id = ?", teacherSubject.TeacherID)
	if result.Error != nil {
		return []TeacherSubjectGrade{}, result.Error
	} else if result.RowsAffected == 0 {
		return []TeacherSubjectGrade{}, fmt.Errorf("no teacher subject found")
	}

	return teacherSubjects, nil
}

func GetTeachingSubjectBySubjectID(
	subjectID string,
	teacherID string,
) (
	teacherSubjects []TeacherSubjectGrade,
	err error,
) {
	result := connections.DB.Preload("TeachingClassSubject").Find(&teacherSubjects, "teacher_id = ? AND subject_id = ?", teacherID, subjectID)
	if result.Error != nil {
		return []TeacherSubjectGrade{}, result.Error
	} else if result.RowsAffected == 0 {
		return []TeacherSubjectGrade{}, fmt.Errorf("no teacher subject found")
	}

	return teacherSubjects, nil
}
