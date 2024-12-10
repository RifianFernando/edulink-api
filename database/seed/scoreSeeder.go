package seed

import "github.com/edulink-api/database/models"

func ScoreSeeder() (grades []models.Score) {
	// grades = []models.Score{
	// 	// Siswa 1 7A PAS Math - teacher naem: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      1,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          90,
	// 	},
	// 	// Siswa 1 7A PTS Math - teacher naem: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      1,
	// 		AssignmentID:   2,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          100,
	// 	},
	// 	// Siswa 2 7A PAS Math - teacher name: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      2,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          80,
	// 	},
	// 	// siswa 3 7A PAS Math - teacher name: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      3,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          85,
	// 	},
	// 	// Arif 7A PAS Math - teacher name: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      4,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          50,
	// 	},
	// 	// Arif 7A PTS Math - teacher name: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      4,
	// 		AssignmentID:   2,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          60,
	// 	},
	// 	// siswa 5 7D PAS Math - teacherID: 2 -> guru2 (math)
	// 	{
	// 		StudentID:      5, // Rifian
	// 		AssignmentID:   1, // PAS
	// 		SubjectID:      1, // math
	// 		TeacherID:      2, // guru5
	// 		AcademicYearID: 1,
	// 		Score:          100,
	// 	},
	// 	// siswa 6 7D PAS Math - teacherID: 2 -> guru1 (math)
	// 	{
	// 		StudentID:      6,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      2,
	// 		AcademicYearID: 1,
	// 		Score:          90,
	// 	},
	// 	// siswa 7 7B PAS Math - teacherID: 2 -> guru2 (math)
	// 	{
	// 		StudentID:      7,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          60,
	// 	},
	// 	// siswa 8 7C PAS Math - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      8,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      2,
	// 		AcademicYearID: 1,
	// 		Score:          78,
	// 	},
	// 	// siswa 9 7B PAS Math - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      9,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          88,
	// 	},
	// 	// siswa 10 7B PAS Math - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      10,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          98,
	// 	},
	// 	// siswa 1 7A PAS Science - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      1,
	// 		AssignmentID:   1,
	// 		SubjectID:      2,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          100,
	// 	},
	// 	// siswa 2 7A PAS Science - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      2,
	// 		AssignmentID:   1,
	// 		SubjectID:      2,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          90,
	// 	},
	// 	// siswa 3 7A PAS Science - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      3,
	// 		AssignmentID:   1,
	// 		SubjectID:      2,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          60,
	// 	},
	// 	// Arif 7A PTS Science - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      4,
	// 		AssignmentID:   2,
	// 		SubjectID:      2,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          60,
	// 	},
	// 	// Arif 7A PTS PKN - teacherID: 1 -> guru1 (math)
	// 	{
	// 		StudentID:      4,
	// 		AssignmentID:   2,
	// 		SubjectID:      4,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          100,
	// 	},
	// }

	// grades = []models.Score{
	// 	{
	// 		StudentID:      1,
	// 		AssignmentID:   1,
	// 		SubjectID:      1,
	// 		TeacherID:      1,
	// 		AcademicYearID: 1,
	// 		Score:          90,
	// 	},
	// }

	return grades
}
