package models

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
)

type Subject struct {
	SubjectID          int64  `gorm:"primaryKey"`
	GradeID            int64  `json:"id_grade" binding:"required"`
	SubjectName        string `json:"name" binding:"required"`
	DurationPerSession int    `json:"duration_session" binding:"required" validate:"gte=0"`
	DurationPerWeek    int    `json:"duration_week" binding:"required" validate:"gte=0"`
	lib.BaseModel
}

type SubjectModel struct {
	Subject
	Grade Grade `gorm:"foreignKey:GradeID;references:GradeID"`
}

func (Subject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "subjects")
}

func (subject *SubjectModel) GetAllSubjects() (subjects []SubjectModel, err error) {
	result := connections.DB.Preload("Grade").Find(&subjects)
	if result.Error != nil {
		return []SubjectModel{}, result.Error
	} else if result.RowsAffected == 0 {
		return []SubjectModel{}, fmt.Errorf("no subjects found")
	}

	return subjects, nil
}

type DTOAllClassSubjects struct {
	SubjectID   int64  `json:"subject_id"`
	ClassNameID int64  `json:"class_name_id"`
	Grade       int    `json:"grade"`
	Name        string `json:"name"`
	SubjectName string `json:"subject_name"`
}

func GetAllSubjectsClassName() (subjectsClasses []DTOAllClassSubjects, err error) {
	result := connections.DB.
		Table("academic.subjects as s").
		Select("s.subject_id, g.grade, cn.class_name_id, cn.name, s.subject_name").
		Joins("JOIN academic.class_names cn ON s.grade_id = cn.grade_id").
		Joins("JOIN academic.grades g ON s.grade_id = g.grade_id").
		Order("g.grade, cn.name, s.subject_name").
		Scan(&subjectsClasses)

	if result.Error != nil {
		return []DTOAllClassSubjects{}, result.Error
	} else if result.RowsAffected == 0 {
		return []DTOAllClassSubjects{}, fmt.Errorf("no subjects found")
	}

	return subjectsClasses, nil
}
