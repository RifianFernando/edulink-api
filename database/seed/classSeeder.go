package seed

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
)

// ClassSeeder seeds the Class data into the database.
func ClassSeeder() {
	// should add the grade first to reference the class
	grades := []models.Grade{
		{
			Grade: 7,
		},
		{
			Grade: 8,
		},
		{
			Grade: 9,
		},
	}

	for _, grade := range grades {
		connections.DB.Create(&grade)
	}

	ClassNames := []models.ClassName{
		{
			TeacherID: 1,
			GradeID:   1,
			ClassName: "A",
		},
		{
			TeacherID: 1,
			GradeID:   1,
			ClassName: "B",
		},
		{
			TeacherID: 1,
			GradeID:   1,
			ClassName: "C",
		},
		{
			TeacherID: 2,
			GradeID:   1,
			ClassName: "D",
		},
	}

	for _, class := range ClassNames {
		connections.DB.Create(&class)
	}
}
