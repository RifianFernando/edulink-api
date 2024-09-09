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
func ClassSeeder() {
	classes := []models.Class{
		{
			TeacherID:  1,
			ClassName:  "XII IPA 1",
			ClassGrade: "XII",
		},
	}

	for _, class := range classes {
		connections.DB.Create(&class)
	}
}
