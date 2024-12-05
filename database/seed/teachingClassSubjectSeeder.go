package seed

import (
	"github.com/edulink-api/database/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherClassSubjectSeeder() (
	teachingClassSubject []models.TeachingClassSubject,
) {
	teachingClassSubject = []models.TeachingClassSubject{
		{
			TeacherSubjectID: 1,
			ClassNameID:      1,
		},
	}

	return teachingClassSubject
}
