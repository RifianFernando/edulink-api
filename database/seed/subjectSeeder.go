package seed

import (
	"github.com/edulink-api/database/models"
)

// ClassSeeder seeds the Class data into the database.
func SubjectSeeder() (subjects []models.Subject) {
	subjects = []models.Subject{
		// 1
		{
			SubjectName:        "Math",
			DurationPerSession: 60 * 4,
			DurationPerWeek:    60 * 4 * 3,
		},
		// 2
		{
			SubjectName:        "Science",
			DurationPerSession: 60 * 4,
			DurationPerWeek:    60 * 4 * 2,
		},
		// 3
		{
			SubjectName:        "Biology",
			DurationPerSession: 60 * 4,
			DurationPerWeek:    60 * 4 * 4,
		},
		// 4
		{
			SubjectName:        "PKN",
			DurationPerSession: 60 * 2,
			DurationPerWeek:    60 * 4 * 5,
		},
		// 5
		{
			SubjectName:        "Sosiology",
			DurationPerSession: 60 * 4,
			DurationPerWeek:    60 * 4 * 2,
		},
	}

	return subjects
}
