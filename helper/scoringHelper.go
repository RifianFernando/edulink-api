package helper

import (
	"fmt"
	"strconv"

	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/score"
)

// Define the score struct
type Score struct {
	AssignmentID   int64  `json:"AssignmentID"`
	AssignmentType string `json:"AssignmentType"`
	SubjectName    string `json:"SubjectName"`
	Score          int    `json:"Score"`
}

// Define the DTO struct to group scores by student
type DTOAllScoringBySubjectClassName struct {
	StudentID   int64   `json:"StudentID"`
	StudentName string  `json:"StudentName"`
	Scores      []Score `json:"Scores"`
}

func RemapScoringStudentBySubjectClassName(
	result []models.ScoringBySubjectClassName,
) []DTOAllScoringBySubjectClassName {

	// Map to group scores by StudentID
	groupedResult := make(map[int64]DTOAllScoringBySubjectClassName)

	// Iterate through the result and group the data
	for _, item := range result {
		studentID := item.StudentID
		studentName := item.StudentName

		// Create a new score object for each item
		score := Score{
			AssignmentID:   item.AssignmentID,
			AssignmentType: item.TypeAssignment,
			SubjectName:    item.SubjectName,
			Score:          item.Score,
		}

		// If the student doesn't exist in the map, initialize it
		if _, exists := groupedResult[studentID]; !exists {
			groupedResult[studentID] = DTOAllScoringBySubjectClassName{
				StudentID:   studentID,
				StudentName: studentName,
				Scores:      []Score{},
			}
		}

		// Append the score to the student's scores list
		groupedResult[studentID] = DTOAllScoringBySubjectClassName{
			StudentID:   studentID,
			StudentName: studentName,
			Scores:      append(groupedResult[studentID].Scores, score),
		}
	}

	// Convert the map to a slice
	var resultDTO []DTOAllScoringBySubjectClassName
	for _, student := range groupedResult {
		resultDTO = append(resultDTO, student)
	}

	return resultDTO
}

func GetListScoringCreateAndUpdate(
	subjectID string,
	classNameID string,
	userID any,
	request request.InsertAllStudentScoreRequest,
) ([]models.Score, error) {
	if subjectID == "" || classNameID == "" {
		return nil, fmt.Errorf("subject_id and class_name_id are required")
	}

	teacher, err := IsTeachingClassSubjectExist(userID, subjectID, classNameID)
	if err != nil || teacher.TeacherID == 0 {
		return nil, fmt.Errorf("teacher not found")
	}

	// get academic year
	academicYear, err := GetOrCreateAcademicYear()
	if err != nil || academicYear.AcademicYearID == 0 {
		return nil, err
	}

	parsedSubjectID, err := strconv.ParseInt(subjectID, 10, 64)
	if err != nil {
		return nil, err
	}

	// create scoring
	var listScoring []models.Score
	for _, item := range request.InsertStudentRequest {
		scoring := models.Score{
			StudentID:      item.StudentID,
			AssignmentID:   request.AssignmentID,
			TeacherID:      teacher.TeacherID,
			SubjectID:      parsedSubjectID,
			AcademicYearID: academicYear.AcademicYearID,
			Score:          item.Score,
		}
		listScoring = append(listScoring, scoring)
	}

	return listScoring, nil
}
