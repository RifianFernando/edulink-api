package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllScoringBySubjectClassName(c *gin.Context) {
	// Get all subjects
	type DTOAllScoring struct {
		ScoringID int64  `json:"scoring_id"`
		SubjectID int64  `json:"subject_id"`
		StudentID int64  `json:"student_id"`
		Score     int    `json:"score"`
		Grade     int    `json:"grade"`
		Subject   string `json:"subject"`
		Student   string `json:"student"`
	}
	// var scoring []models.Score
	// scoring, err := models.GetAllScoringBySubjectClassID()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Map the subjects to DTO
	var scoringDTO []DTOAllScoring
	// for _, scoring := range scoring {
	// 	scoringDTO = append(scoringDTO, DTOAllScoring{
	// 		ScoringID: scoring.ScoringID,
	// 		SubjectID: scoring.Subject.SubjectID,
	// 		StudentID: scoring.Student.StudentID,
	// 		Score:     scoring.Score,
	// 		Grade:     scoring.Subject.Grade.Grade,
	// 		Subject:   scoring.Subject.SubjectName,
	// 		Student:   scoring.Student.StudentName,
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{"scoring": scoringDTO})
}
