package helper

import (
	"strconv"

	"github.com/edulink-api/database/models"
)

func GetAllStudentPersonalDataArchive(
	academicYearStart string,
	academicYearEnd string,
) (
	student []models.StudentModel,
	err error,
) {
	academicYearStart, academicYearEnd = FormatSemesterDateAcademicYear(academicYearStart, academicYearEnd)
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
	classID string,
) (
	AllStudentAttendanceArchive []models.AttendanceYearSummaryStudent,
	err error,
) {
	// Get all student attendance
	GetAcademicYearStart, GetAcademicYearEnd = FormatSemesterDateAcademicYear(
		GetAcademicYearStart, GetAcademicYearEnd,
	)
	AllStudentAttendanceArchive, err = models.GetAllAttendanceArchive(
		GetAcademicYearStart,
		GetAcademicYearEnd,
		classID,
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

func GetAllClassArchiveByGradeID(
	academicYearStart string,
	academicYearEnd string,
	gradeID string,
) (
	studentsClassArchive []models.ClassListArchive,
	err error,
) {
	academicYearStart, academicYearEnd = FormatSemesterDateAcademicYear(academicYearStart, academicYearEnd)

	studentsClassArchive, err = models.GetAllClassArchiveByGradeID(
		academicYearStart,
		academicYearEnd,
		gradeID,
	)
	if err != nil {
		return nil, err
	}

	return studentsClassArchive, nil
}
