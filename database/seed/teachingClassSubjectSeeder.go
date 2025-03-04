package seed

import (
	"github.com/edulink-api/database/models"
)

// ClassSeeder seeds the Class data into the database.
func TeacherClassSubjectSeeder() (
	teachingClassSubject []models.TeachingClassSubject,
) {
	teachingClassSubject = []models.TeachingClassSubject{
		// Guru 1 mengajar kelas 7A mata pelajaran Math
		{
			TeacherSubjectID: 1,
			ClassNameID:      1,
			AcademicYearID:   1,
		},
		// Guru 1 mengajar kelas 7B mata pelajaran Math
		{
			TeacherSubjectID: 1,
			ClassNameID:      2,
			AcademicYearID:   1,
		},
		// Guru 1 mengajar kelas 7A mata pelajaran Biology
		{
			TeacherSubjectID: 3,
			ClassNameID:      1,
			AcademicYearID:   1,
		},
		// Guru 2 mengajar kelas 7C mata pelajaran Math
		{
			TeacherSubjectID: 5,
			ClassNameID:      3,
			AcademicYearID:   1,
		},
		// Guru 2 mengajar kelas 7D mata pelajaran Math
		{
			TeacherSubjectID: 5,
			ClassNameID:      5,
			AcademicYearID:   1,
		},
		// Guru testadmin mengajar kelas 8 A mata pelajaran sosiology
		{
			TeacherSubjectID: 8,
			ClassNameID:      5,
			AcademicYearID:   1,
		},
		// Guru testadmin mengajar kelas 8 B mata pelajaran sosiology
		{
			TeacherSubjectID: 8,
			ClassNameID:      6,
			AcademicYearID:   1,
		},
		// Guru testadmin mengajar kelas 8 C mata pelajaran sosiology
		{
			TeacherSubjectID: 8,
			ClassNameID:      7,
			AcademicYearID:   1,
		},
	}

	return teachingClassSubject
}
