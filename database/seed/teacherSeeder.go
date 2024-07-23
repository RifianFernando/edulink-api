package seed

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration"
)

// init initializes the package by loading environment variables and connecting to the database.
func init() {
	connections.LoadEnvVariables()
	connections.ConnecToDB()
}

// ClassSeeder seeds the Class data into the database.
func TeacherSeeder() {
	classes := []migration.Teacher{
		{
			UserID: 1,
			TeachingHour: 20,
		},
	}

	for _, class := range classes {
		connections.DB.Create(&class)
	}
}
