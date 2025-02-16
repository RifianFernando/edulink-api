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

	// 1. Get all teaching class subjects
	tx.Model(&models.TeachingClassSubject{}).Where("academic_year_id = ?", academicYear.AcademicYearID).Find(&models.TeachingClassSubject{})
	
	// 2. generate schedule teaching class subjects
	// 3. insert schedule teaching class subjects
	return nil
}
