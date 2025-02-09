package helper

import (
	"fmt"
	"strconv"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/models"
	req "github.com/edulink-api/request/schedule"
)

func GenerateAndCreateScheduleTeachingClassSubject(
	req req.InsertScheduleRequest,
) error {

	// get academic year
	academicYear, err := GetOrCreateAcademicYear()
	if err != nil || academicYear.AcademicYearID == 0 {
		return err
	}

	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// update teacher teaching hour
	var listTeacherIDIn []int64
	var listTeacherSubjectIDIn []int64
	queryUpdateTeacherTeachingHour := "UPDATE academic.teachers SET teaching_hour = CASE"

	// create teacher subject and update teacher subject with upsert method sql query like teacher model
	queryUpsertTeacherSubject := "INSERT INTO academic.teaching_class_subjects (teacher_subject_id, class_name_id, academic_year_id) VALUES "
	for _, teacher := range req.ScheduleRequest {
		// first update the teaching hour using bulk update
		teacherIDParsed := strconv.FormatInt(teacher.TeacherID, 10)
		queryUpdateTeacherTeachingHour += fmt.Sprintf(" WHEN teacher_id = %s THEN %d", teacherIDParsed, teacher.TeachingHour)
		listTeacherIDIn = append(listTeacherIDIn, teacher.TeacherID)

		// upsert the teaching class subject too like the teacher subject
		for _, DataTeaching := range teacher.DataTeaching {
			// get teacher subject by teacher_id and subject_id
			TeacherSubject := models.TeacherSubject{}
			if err := tx.Where("teacher_id = ? AND subject_id = ?", teacher.TeacherID, DataTeaching.SubjectID).First(&TeacherSubject).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("teacherID with: %d with subjectID %d not found", teacher.TeacherID, DataTeaching.SubjectID)
			}

			// append the teacher subject id
			listTeacherSubjectIDIn = append(listTeacherSubjectIDIn, TeacherSubject.TeacherSubjectID)

			// upsert the teaching class subject
			for index, ClassTeaching := range DataTeaching.ClassNameID {
				if index == len(DataTeaching.ClassNameID)-1 {
					queryUpsertTeacherSubject += fmt.Sprintf("(%d, %d, %d)", TeacherSubject.TeacherSubjectID, ClassTeaching, academicYear.AcademicYearID)
				} else {
					queryUpsertTeacherSubject += fmt.Sprintf("(%d, %d, %d),", TeacherSubject.TeacherSubjectID, ClassTeaching, academicYear.AcademicYearID)
				}
			}
		}
	}
	// execute update teacher teaching hour and upsert teacher subject
	queryUpdateTeacherTeachingHour += " END WHERE teacher_id IN ?"
	queryUpsertTeacherSubject += fmt.Sprintf(" ON CONFLICT (id_teacher_subject, id_class_name) DO UPDATE SET deleted_at = NULL, updated_at = NOW(), academic_year_id = %d", academicYear.AcademicYearID)

	// first will be delete all teaching class subject by teacher subject id that will be updated
	for _, teacherSubjectID := range listTeacherSubjectIDIn {
		result := tx.Where("teacher_subject_id = ?", teacherSubjectID).Delete(&models.TeachingClassSubject{})
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	if err := tx.Exec(queryUpdateTeacherTeachingHour, listTeacherIDIn).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
