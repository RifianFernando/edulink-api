package seed

import (
	"github.com/edulink-api/models"
)

// ClassSeeder seeds the Class data into the database.
func SubjectSeeder() (subjects []models.Subject) {
	subjects = []models.Subject{
		{
			GradeID:                1,
			SubjectName:            "Math",
			SubjectDurationMinutes: 60 * 4,
		},
		{
			GradeID:                1,
			SubjectName:            "Science",
			SubjectDurationMinutes: 60 * 4,
		},
		{
			GradeID:                1,
			SubjectName:            "Biology",
			SubjectDurationMinutes: 60 * 4,
		},
		{
			GradeID:                1,
			SubjectName:            "PKN",
			SubjectDurationMinutes: 60 * 2,
		},
	}

	return subjects
}
