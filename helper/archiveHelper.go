package helper

import (
	"strconv"

	"github.com/edulink-api/database/models"
)

func GetAllStudentPersonalDataArchive(
	academicYearStart string,
	academicYearEnd string,
) (
	student []models.Student,
	err error,
) {
	student, err = models.GetAllStudentPersonalDataArchive(
		academicYearStart,
		academicYearEnd,
	)
	if err != nil {
		return student, err
	}

	return student, nil
}

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

func GetAllStudentScoreArchive(
	acadmicYear models.AcademicYear,
	classID string,
) (
	[]DTOAllScoringBySubjectClassName,
	error,
) {
	// Get the scoring data from the model
	parsedAcademicYearID := strconv.FormatInt(acadmicYear.AcademicYearID, 10)
	result, err := models.GetSummariesScoringStudentBySubjectClassName(classID, parsedAcademicYearID)
	if err != nil {
		return nil, err
	}

	resultMap := RemapScoringStudentBySubjectClassName(result)
	resultDTO := GetAverageScoreByStudentResult(resultMap)

	return resultDTO, nil
}
