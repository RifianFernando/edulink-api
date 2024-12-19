package helper

import (
	"fmt"
	"strconv"

	"github.com/edulink-api/connections"
	req "github.com/edulink-api/request/schedule"
)

// TODO: after refactoring table teacher_subjects, this function should be updated
func GenerateAndCreateScheduleTeachingClassSubject(req req.InsertScheduleRequest) error {

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
	queryUpdateTeacherTeachingHour := "UPDATE academic.teachers SET teaching_hour = CASE"

	// create teacher subject with upsert method sql query
	queryUpsertTeacherSubject := "INSERT INTO academic.teacher_subjects (teacher_id, subject_id) VALUES "
	for _, teacher := range req.ScheduleRequest {
		teacherIDParsed := strconv.FormatInt(teacher.TeacherID, 10)
		queryUpdateTeacherTeachingHour += fmt.Sprintf(" WHEN teacher_id = %s THEN %d", teacherIDParsed, teacher.TeachingHour)
		listTeacherIDIn = append(listTeacherIDIn, teacher.TeacherID)

		for idx, SubjectID := range teacher.SubjectID {
			queryUpsertTeacherSubject += fmt.Sprintf("(%d, %d)", teacher.TeacherID, SubjectID)
			if idx != len(teacher.SubjectID)-1 {
				queryUpsertTeacherSubject += ", "
			}
		}
	}
	// execute update teacher teaching hour and upsert teacher subject
	queryUpdateTeacherTeachingHour += " END WHERE teacher_id IN ?"
	queryUpsertTeacherSubject += " ON CONFLICT (teacher_id, subject_id) DO NOTHING"

	if err := tx.Exec(queryUpdateTeacherTeachingHour, listTeacherIDIn).Error; err != nil {
		tx.Rollback()
		return err
	}

	// if (len(teacherSubjects) == 0) {
	// return fmt.Errorf("invalid data")
	// }

	return tx.Commit().Error
}
