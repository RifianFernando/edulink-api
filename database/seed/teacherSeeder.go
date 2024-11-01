package seed

import (
	"github.com/edulink-api/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherSeeder() (teachers []models.Teacher) {
	teachers = []models.Teacher{
		{
			UserID:       1,
			TeachingHour: 20,
		},
		{
			UserID:       2,
			TeachingHour: 20,
		},
	}

	return teachers
}
