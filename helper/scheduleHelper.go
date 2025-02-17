package helper

import (
	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/models"
)
// start hour schedule each day is 7 am
// const startHour = 7
// end hour schedule each day is 5 pm
// const endHour = 17

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
	var teachingClassSubjects []models.TeachingClassSubjectModel
	tx.Model(&models.TeachingClassSubjectModel{}).
		Preload("ClassName").
		Preload("TeacherSubject").
		Where("academic_year_id = ?", academicYear.AcademicYearID).Find(&teachingClassSubjects)

	// TODO: jadi gw bikinnya kgk ngecek database yang lama karena pengen cepat jadi gw bikinnya langsung insert aja kalo mau best practice coba di check databasenya sesuai sama academic year sekarang terus juga cek yang deleted_at nya null dan juga check learning schedulenya masih aktif atau tidak

	// 2. generate schedule teaching class subjects
	// TODO: ohh iya bisa pakai hash map buat nentuin jadwal gk boleh sama saat melakukan create INGATTTTTTTTTTTTT!!!!!!!!!!!!!!!
	// [teacher_subject_id] => [schedule_id]
	// [class_name_id] => [schedule_id]
	uniqueTeacherSchedule := make(map[int64]int)
	uniqueClassChedule := make(map[int64]int)

	// teaching class subject is need duration session for each schedule
	for _, teachingClassSubject := range teachingClassSubjects {
	// create learning schedule by schedule id
		// i for id schedule
		for i := 1; i <= 70; i++ {
			teacherID := teachingClassSubject.TeacherSubject.TeacherID
			teachingClassNameID := teachingClassSubject.ClassNameID
			if _, teacherExists := uniqueTeacherSchedule[teacherID]; !teacherExists {
				if _, classExists := uniqueClassChedule[teachingClassNameID]; !classExists {
					uniqueTeacherSchedule[teacherID] = i
					uniqueClassChedule[teachingClassNameID] = i
					var learningSchedule models.LearningSchedule
					learningSchedule.TeachingClassSubjectID = teachingClassSubject.TeachingClassSubjectID
				}
			}
		}
	}

	// 3. insert schedule teaching class subjects
	return nil
}
