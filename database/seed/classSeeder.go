package seed

import (
	"github.com/edulink-api/database/models"
)

func ClassSeeder() (className []models.ClassName) {
	className = []models.ClassName{
		{ // 1
			TeacherID: 1,
			GradeID:   1,
			Name:      "A",
		},
		{ // 2
			TeacherID: 1,
			GradeID:   1,
			Name:      "B",
		},
		{ // 3
			TeacherID: 1,
			GradeID:   1,
			Name:      "C",
		},
		{ // 4
			TeacherID: 2,
			GradeID:   1,
			Name:      "D",
		},
		{ // 5
			TeacherID: 2,
			GradeID:   2,
			Name:      "A",
		},
		{ // 6
			TeacherID: 2,
			GradeID:   2,
			Name:      "B",
		},
		{ // 7
			TeacherID: 2,
			GradeID:   2,
			Name:      "C",
		},
	}

	return className
}
