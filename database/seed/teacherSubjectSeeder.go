package seed

import (
	"github.com/edulink-api/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherSubjectSeeder() (teachers []models.TeacherSubject) {
	teachers = []models.TeacherSubject{
		{
			TeacherID:              1,
			SubjectID:              1,
		},
		{
			TeacherID:              1,
			SubjectID:              2,
		},
		{
			TeacherID:              1,
			SubjectID:              3,
		},
		{
			TeacherID:              1,
			SubjectID:              4,
		},
		{
			TeacherID:              2,
			SubjectID:              2,
		},
		{
			TeacherID:              3,
			SubjectID:              3,
		},
	}

	return teachers
}
