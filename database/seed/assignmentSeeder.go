package seed

import "github.com/edulink-api/database/models"

func AssignmentSeeder() (grades []models.Assignment) {
	grades = []models.Assignment{
		{
			TypeAssignment: "PAS",
		},
		{
			TypeAssignment: "PTS",
		},
		{
			TypeAssignment: "Quizz 1",
		},
		{
			TypeAssignment: "Quizz 2",
		},
	}

	return grades
}
