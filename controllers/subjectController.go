package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/database/models"
	"github.com/gin-gonic/gin"
)

func GetAllSubject(c *gin.Context) {
	// Get all subjects
	type DTOAllSubjects struct {
		SubjectID          int64  `json:"subject_id"`
		SubjectName        string `json:"subject_name"`
		DurationPerSession int    `json:"subject_duration_session"`
		DurationPerWeek    int    `json:"subject_duration_per_week"`
	}
	var subject models.Subject
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
			SubjectName:        subject.SubjectName,
			DurationPerSession: subject.DurationPerSession,
			DurationPerWeek:    subject.DurationPerWeek,
		})
	}

	c.JSON(http.StatusOK, gin.H{"subjects": subjectsDTO})
}

func GetSubjectClassNameStudentsByID(c *gin.Context) {
	// Get the subject ID
	subjectID := c.Param("subject_id")
	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subject ID is required"})
		return
	}

	// Get the subject
	var subject models.Subject
	subjectResult, err := subject.GetSubjectByID(subjectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the class name
	classNameID := c.Param("class_id")
	if classNameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class name ID is required"})
		return
	}
	var className models.ClassNameModel
	err = className.GetClassNameModelByID(classNameID)
	if err != nil || className.ClassNameID == 0 || className.Grade.Grade == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the students by class name
	var student models.Student
	resultStudents, err := student.GetAllStudentsByClassID(classNameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error students": err.Error()})
		return
	}

	// Map the subject and class name to DTO
	type DTOSubjectClassName struct {
		GradeClassName string           `json:"grade_class_name"`
		SubjectName    string           `json:"subject_name"`
		Students       []models.Student `json:"students"`
	}

	subjectClassNameDTO := DTOSubjectClassName{
		GradeClassName: strconv.FormatInt(int64(className.Grade.Grade), 10) + "-" + className.Name,
		SubjectName:    subjectResult.SubjectName,
		Students:       resultStudents,
	}

	c.JSON(http.StatusOK, gin.H{"subject": subjectClassNameDTO})
}
