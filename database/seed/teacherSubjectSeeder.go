package seed

import (
	"github.com/edulink-api/database/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherSubjectSeeder() (teachers []models.TeacherSubject) {
	teachers = []models.TeacherSubject{
		//1. Guru 1 mengajar mata pelajaran kelas 7 Math a, b
		{
			TeacherID:              1,
			SubjectID:              1,
		},
		//2. Guru 1 mengajar mata pelajaran kelas 7 science
		{
			TeacherID:              1,
			SubjectID:              2,
		},
		//3. Guru 1 mengajar mata pelajaran kelas 7 Biology
		{
			TeacherID:              1,
			SubjectID:              3,
		},
		//4. Guru 1 mengajar mata pelajaran kelas 7 PKN
		{
			TeacherID:              1,
			SubjectID:              4,
		},
		//5. Guru Admin mengajar mata pelajaran kelas 7 Math c, dan d
		{
			TeacherID:              2,
			SubjectID:              1,
		},
		//6. Guru Admin mengajar mata pelajaran kelas 7 Science
		{
			TeacherID:              2,
			SubjectID:              2,
		},
		//7. Teacher Non Homeroom mengajar mata pelajaran kelas 7 Biology
		{
			TeacherID:              3,
			SubjectID:              3,
		},
	}

	return teachers
}
