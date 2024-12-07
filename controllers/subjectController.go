package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/res"
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

func GetAllSubjectClassName(c *gin.Context) {
	// get user role and id
	userRole, exist := c.Get("user_type")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User role found"})
		return
	}
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	// Get all subjects
	type DTOAllSubjectsClassName struct {
		SubjectID      int64  `json:"subject_id"`
		GradeClassName string `json:"grade_class_name"`
		SubjectName    string `json:"subject_name"`
	}
	var subjectClassNameDTO []DTOAllSubjectsClassName

	if userRole == "teacher" {
		// get teacher id
		var teacher models.Teacher
		teacher.UserID = userID.(int64)
		err := teacher.GetTeacherByModel()
		if err != nil || teacher.TeacherID == 0 {
			res.AbortUnauthorized(c)
		}

		var teacherSubject models.TeacherSubjectGrade
		teacherSubject.TeacherID = teacher.TeacherID
		result, err := teacherSubject.GetTeachingSubjectByID()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		for _, subject := range result {
			for _, subjectDetail := range subject.TeachingClassSubject {
				var className models.ClassName
				classNameID := strconv.FormatInt(subjectDetail.ClassNameID, 10)
				classResult, err := className.GetClassNameById(classNameID)
				if err != nil {
					res.AbortUnauthorized(c)
				}
				var gradeClassName = strconv.FormatInt(int64(subject.Subject.Grade.Grade), 10) + "-" + classResult.Name
				subjectClassNameDTO = append(subjectClassNameDTO, DTOAllSubjectsClassName{
					SubjectID:      subject.SubjectID,
					GradeClassName: gradeClassName,
					SubjectName:    subject.Subject.SubjectName,
				})
			}
		}

		c.JSON(http.StatusOK, gin.H{"subjects": subjectClassNameDTO})
		return
	}

	// Map the subjects to DTO
	result, err := models.GetAllSubjectsClassName()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, subject := range result {
		var gradeClassName = strconv.FormatInt(int64(subject.Grade), 10) + "-" + subject.Name
		subjectClassNameDTO = append(subjectClassNameDTO, DTOAllSubjectsClassName{
			SubjectID:      subject.SubjectID,
			GradeClassName: gradeClassName,
			SubjectName:    subject.SubjectName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"subjects": subjectClassNameDTO})
}
