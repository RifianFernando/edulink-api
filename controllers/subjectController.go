package controllers

import (
	"net/http"

	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllSubject(c *gin.Context) {
	// Get all subjects
	type DTOAllSubjects struct {
		SubjectID              int64  `json:"subject_id"`
		GradeID                int64  `json:"grade_id"`
		SubjectName            string `json:"subject_name"`
		SubjectDurationMinutes int    `json:"subject_duration_minutes"`
	}
	var subject models.Subject
	subjects, err := subject.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Map the subjects to DTO
	var subjectsDTO []DTOAllSubjects
	for _, subject := range subjects {
		subjectsDTO = append(subjectsDTO, DTOAllSubjects{
			SubjectID:              subject.SubjectID,
			GradeID:                subject.GradeID,
			SubjectName:            subject.SubjectName,
			SubjectDurationMinutes: subject.SubjectDurationMinutes,
		})
	}

	c.JSON(http.StatusOK, gin.H{"subjects": subjectsDTO})
}
