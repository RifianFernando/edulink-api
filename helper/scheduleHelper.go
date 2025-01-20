package helper

import (
	"fmt"
	"strconv"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/models"
	req "github.com/edulink-api/request/schedule"
)

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

	// create teacher subject and update teacher subject with upsert method sql query like teacher model
	queryUpsertTeacherSubject := "INSERT INTO academic.teaching_class_subjects (teacher_subject_id, class_name_id) VALUES "
	for _, teacher := range req.ScheduleRequest {
		// first update the teaching hour using bulk update
		teacherIDParsed := strconv.FormatInt(teacher.TeacherID, 10)
		queryUpdateTeacherTeachingHour += fmt.Sprintf(" WHEN teacher_id = %s THEN %d", teacherIDParsed, teacher.TeachingHour)
		listTeacherIDIn = append(listTeacherIDIn, teacher.TeacherID)

		// upsert the teaching class subject too like the teacher subject
		for _, DataTeaching := range teacher.DataTeaching {
			//TODO: get the teacher subject ID with get teaching class subject, but I think we need to get optimize this database querry for get teacher subject ID and then we are going through join the teaching class subject, for assign new value method with bulk upsert

			// TODO: e.g. using query select for searching the teacher subject ID
			var teachingClassSubject models.TeachingClassSubject
			teachingClassSubject.TeacherSubjectID = int64(DataTeaching.SubjectID)
			for _, ClassTeaching := range DataTeaching.ClassNameID {
				teachingClassSubject.ClassNameID = ClassTeaching

				// TODO: create the query for upsert value for teaching class subject using the existing data with bulk upsert for best practice
				queryUpsertTeacherSubject += ""
			}
		}
	}
	// execute update teacher teaching hour and upsert teacher subject
	queryUpdateTeacherTeachingHour += " END WHERE teacher_id IN ?"
	// TODO: after refactoring table teacher_subjects, this query should be used
	// queryUpsertTeacherSubject += " ON CONFLICT (teacher_id, subject_id) DO NOTHING"

	if err := tx.Exec(queryUpdateTeacherTeachingHour, listTeacherIDIn).Error; err != nil {
		tx.Rollback()
		return err
	}

	// if (len(teacherSubjects) == 0) {
	// return fmt.Errorf("invalid data")
	// }

	return tx.Commit().Error
}
