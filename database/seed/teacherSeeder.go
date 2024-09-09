package seed

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/models"
)

// init initializes the package by loading environment variables and connecting to the database.
func init() {
	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

// ClassSeeder seeds the Class data into the database.
func TeacherSeeder() {
	teachers := []models.Teacher{
		{
			UserID:       1,
			TeachingHour: 20,
		},
	}

	for _, teacher := range teachers {
		connections.DB.Create(&teacher)
	}
}
