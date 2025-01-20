package models

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type Subject struct {
	SubjectID          int64  `gorm:"primaryKey"`
	SubjectName        string `json:"name" binding:"required"`
	DurationPerSession int    `json:"duration_session" binding:"required" validate:"gte=0"`
	DurationPerWeek    int    `json:"duration_week" binding:"required" validate:"gte=0"`
	lib.BaseModel
}

func (Subject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "subjects")
}

func (subject *Subject) GetAllSubjects() (subjects []Subject, err error) {
	result := connections.DB.Find(&subjects)
	if result.Error != nil {
		return []Subject{}, result.Error
	} else if result.RowsAffected == 0 {
		return []Subject{}, fmt.Errorf("no subjects found")
	}

	return subjects, nil
}

func (subject *Subject) GetSubjectByID(subjectID string) (Subject, error) {
	result := connections.DB.Where("subject_id = ?", subjectID).First(&subject)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return *subject, fmt.Errorf("subject not found")
		}
		return *subject, result.Error
	} else if result.RowsAffected == 0 {
		return *subject, fmt.Errorf("no subject found")
	}

	return *subject, nil
}
