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
		ClassNameID    int64  `json:"class_name_id"`
		GradeClassName string `json:"grade_class_name"`
		SubjectName    string `json:"subject_name"`
	}
	var subjectClassNameDTO []DTOAllSubjectsClassName

	if userRole == "teacher" || userRole == "homeroom_teacher" {
		// get teacher id
		var teacher models.Teacher
		teacher.UserID = userID.(int64)
		err := teacher.GetTeacherByModel()
		if err != nil || teacher.TeacherID == 0 {
			res.AbortUnauthorized(c)
			return
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
					return
				}
				var gradeClassName = strconv.FormatInt(int64(subject.Subject.Grade.Grade), 10) + "-" + classResult.Name
				subjectClassNameDTO = append(subjectClassNameDTO, DTOAllSubjectsClassName{
					SubjectID:      subject.SubjectID,
					ClassNameID:    subjectDetail.ClassNameID,
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
		var gradeClassName = strconv.FormatInt(int64(subject.Grade), 10) + subject.Name
		subjectClassNameDTO = append(subjectClassNameDTO, DTOAllSubjectsClassName{
			SubjectID:      subject.SubjectID,
			ClassNameID:    subject.ClassNameID,
			GradeClassName: gradeClassName,
			SubjectName:    subject.SubjectName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"subjects": subjectClassNameDTO})
}

func GetSubjectClassNameStudentsByID(c *gin.Context) {
	// Get the subject ID
	subjectID := c.Param("subject_id")
	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subject ID is required"})
		return
	}

	// Get the subject
	var subject models.SubjectModel
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
	var className models.ClassName
	classNameResult, err := className.GetClassNameById(classNameID)
	if err != nil {
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
		GradeClassName: strconv.FormatInt(int64(subjectResult.Grade.Grade), 10) + "-" + classNameResult.Name,
		SubjectName:    subjectResult.SubjectName,
		Students:       resultStudents,
	}

	c.JSON(http.StatusOK, gin.H{"subject": subjectClassNameDTO})
}
