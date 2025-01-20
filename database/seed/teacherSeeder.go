package seed

import (
	"github.com/edulink-api/database/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherSeeder() (teachers []models.Teacher) {
	teachers = []models.Teacher{
		{ // ngajar 7c dan 8a
			UserID:       1,
			TeachingHour: 20,
		},
		{ // ngajar 7c dan 8a
			UserID:       2,
			TeachingHour: 20,
		},
		{
			UserID:       3,
			TeachingHour: 10,
		},
		{
			UserID:       4,
			TeachingHour: 10,
		},
	}

	return teachers
}
