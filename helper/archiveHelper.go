package helper

import (
	"github.com/edulink-api/database/models"
)

// func GetAllStudentPersonalDataArchive() (models.AcademicYear, error) {

// 	return academicYear, nil
// }

func GetAllStudentAttendanceArchive(
	GetAcademicYearStart string,
	GetAcademicYearEnd string,
) (
	AllStudentAttendanceArchive []models.AttendanceYearSummaryStudent,
	err error,
) {
	// Get all student attendance
	AllStudentAttendanceArchive, err = models.GetAllAttendanceArchive(
		GetAcademicYearStart,
		GetAcademicYearEnd,
	)
	if err != nil {
		return nil, err
	}

	// Get all student attendance

	return AllStudentAttendanceArchive, nil
}
