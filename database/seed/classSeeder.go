package seed

import (
	"github.com/edulink-api/models"
)

func ClassSeeder() (className []models.ClassName) {
	className = []models.ClassName{
		{
			TeacherID: 1,
			GradeID:   1,
			Name:      "A",
		},
		{
			TeacherID: 1,
			GradeID:   1,
			Name:      "B",
		},
		{
			TeacherID: 1,
			GradeID:   1,
			Name:      "C",
		},
		{
			TeacherID: 2,
			GradeID:   1,
			Name:      "D",
		},
	}

	return className
}
