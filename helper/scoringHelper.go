package helper

import "github.com/edulink-api/database/models"

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
