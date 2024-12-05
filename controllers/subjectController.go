package controllers

import (
	"net/http"

	"github.com/edulink-api/database/models"
	"github.com/gin-gonic/gin"
)

func GetAllSubject(c *gin.Context) {
	// Get all subjects
	type DTOAllSubjects struct {
		SubjectID          int64  `json:"subject_id"`
		Grade              int    `json:"grade"`
		SubjectName        string `json:"subject_name"`
		DurationPerSession int    `json:"subject_duration_minutes"`
	}
	var subject models.SubjectModel
	subjects, err := subject.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map the subjects to DTO
	var subjectsDTO []DTOAllSubjects
	for _, subject := range subjects {
		subjectsDTO = append(subjectsDTO, DTOAllSubjects{
			SubjectID:          subject.SubjectID,
			Grade:              subject.Grade.Grade,
			SubjectName:        subject.SubjectName,
			DurationPerSession: subject.DurationPerSession,
		})
	}

	c.JSON(http.StatusOK, gin.H{"subjects": subjectsDTO})
}
