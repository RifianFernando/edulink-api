package helper

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/models"
	req "github.com/edulink-api/request/schedule"
)

// GenerateTeachingClassSubject handles the generation of teaching class subjects
func GenerateTeachingClassSubject(req req.InsertScheduleRequest) error {
	// Get or create academic year
	academicYear, _ := GetOrCreateAcademicYear()

	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Lists for bulk processing
	var listTeacherIDIn []int64
	var listTeacherSubjectIDIn []int64
	var records []models.TeachingClassSubject
	// Update teaching hours
	query := "UPDATE academic.teachers SET teaching_hour = CASE"

	// Loop through schedule requests
	for _, teacher := range req.ScheduleRequest {
		query += fmt.Sprintf(" WHEN teacher_id = %d THEN %d", teacher.TeacherID, teacher.TeachingHour)
		listTeacherIDIn = append(listTeacherIDIn, teacher.TeacherID)

		for _, DataTeaching := range teacher.DataTeaching {
			// Get teacher subject
			TeacherSubject := models.TeacherSubject{}
			if err := tx.Where("teacher_id = ? AND subject_id = ?", teacher.TeacherID, DataTeaching.SubjectID).First(&TeacherSubject).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("teacherID with: %d with subjectID %d not found", teacher.TeacherID, DataTeaching.SubjectID)
			}

			// Collect teacher_subject_id
			listTeacherSubjectIDIn = append(listTeacherSubjectIDIn, TeacherSubject.TeacherSubjectID)

			// Prepare records for upsert
			for _, ClassTeaching := range DataTeaching.ClassNameID {
				records = append(records, models.TeachingClassSubject{
					TeacherSubjectID: TeacherSubject.TeacherSubjectID,
					ClassNameID:      ClassTeaching,
					AcademicYearID:   academicYear.AcademicYearID,
				})
			}
		}
	}
	query += " END WHERE teacher_id IN ?"
	if err := tx.Exec(query, listTeacherIDIn).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Upsert teaching class subjects
	if err := models.UpsertTeachingClassSubjects(tx, listTeacherSubjectIDIn, records); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
