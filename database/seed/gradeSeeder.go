package seed

import "github.com/edulink-api/database/models"

func GradeSeeder() (grades []models.Grade) {
	grades = []models.Grade{
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

	return grades
}
