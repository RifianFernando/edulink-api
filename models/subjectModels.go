package models

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
)

type Subject struct {
	SubjectID              int64  `gorm:"primaryKey"`
	GradeID                int64  `json:"id_grade" binding:"required"`
	SubjectName            string `json:"name" binding:"required"`
	SubjectDurationMinutes int    `json:"duration" binding:"required" validate:"gte=0"`
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
