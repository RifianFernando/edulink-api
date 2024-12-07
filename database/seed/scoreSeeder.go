package seed

import "github.com/edulink-api/database/models"

func ScoreSeeder() (grades []models.Score) {
	grades = []models.Score{
		// Siswa 1 7A PAS Math - teacher naem: 1 -> guru1 (math)
		{
			StudentID:    1,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        90,
		},
		// Siswa 2 7A PAS Math - teacher name: 1 -> guru1 (math)
		{
			StudentID:    2,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        80,
		},
		{
			StudentID:    3,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        85,
		},
		{
			StudentID:    4,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        50,
		},
		{
			StudentID:    5,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    5,
			Score:        100,
		},
		{
			StudentID:    6,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    5,
			Score:        90,
		},
		{
			StudentID:    7,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        60,
		},
		{
			StudentID:    8,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    5,
			Score:        78,
		},
		{
			StudentID:    9,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        88,
		},
		{
			StudentID:    10,
			AssignmentID: 1,
			SubjectID:    1,
			TeacherID:    1,
			Score:        98,
		},
	}

	return grades
}
