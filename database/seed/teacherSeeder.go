package seed

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherSeeder() {
	teachers := []models.Teacher{
		{
			UserID:       1,
			TeachingHour: 20,
		},
		{
			UserID:       2,
			TeachingHour: 20,
		},
	}

	for _, teacher := range teachers {
		connections.DB.Create(&teacher)
	}
}
