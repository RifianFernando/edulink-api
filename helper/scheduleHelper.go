package helper

import (
	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/models"
)

func GenerateNewScheduleTeachingClassSubject(academicYear models.AcademicYear) error {
	// Get or create academic year
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get all teaching class subjects
	return nil
}
