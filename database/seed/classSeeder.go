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
func ClassSeeder() {
	classes := []migration.Class{
		{
			ClassID: 1,
			ClassName:  "XII IPA 1",
			ClassGrade: "XII",
			TeacherID:  1,
		},
		{
			ClassID: 2,
			ClassName:  "XII IPA 2",
			ClassGrade: "XII",
			TeacherID:  2,
		},
	}

	for _, class := range classes {
		connections.DB.Create(&class)
	}
}
