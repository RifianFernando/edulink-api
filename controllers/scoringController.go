package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/score"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func GetAllScoringBySubjectClassName(c *gin.Context) {
	// Get parameters from the request
	subjectID := c.Param("subject_id")
	classNameID := c.Param("class_name_id")

	if subjectID == "" || classNameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject_id and class_name_id are required"})
		return
	}

	// Get the scoring data from the model
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	// get teacher id
	var teacher models.Teacher
	teacher.UserID = userID.(int64)
	err := teacher.GetTeacherByModel()
	if err != nil || teacher.TeacherID == 0 {
		res.AbortUnauthorized(c)
		return
	}
	teacherID := strconv.FormatInt(teacher.TeacherID, 10)
	result, err := models.GetAllScoringBySubjectClassID(subjectID, classNameID, teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	// Send the grouped result as a response
	c.JSON(http.StatusOK, gin.H{"score": resultDTO})
}

func CreateStudentsScoringBySubjectClassName(c *gin.Context) {
	var request request.InsertAllStudentScoreRequest
	var allErrors []map[string]string

	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateAllStudentScore(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	// Get parameters from the request
	subjectID := c.Param("subject_id")
	classNameID := c.Param("class_name_id")

	if subjectID == "" || classNameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject_id and class_name_id are required"})
		return
	}

	// Get the scoring data from the model
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	// get teacher id
	var teacher models.Teacher
	teacher.UserID = userID.(int64)
	err := teacher.GetTeacherByModel()
	if err != nil || teacher.TeacherID == 0 {
		res.AbortUnauthorized(c)
		return
	}
	teacherID := strconv.FormatInt(teacher.TeacherID, 10)
	classNameIDParsed, err := strconv.ParseInt(classNameID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := models.GetTeachingSubjectBySubjectID(
		subjectID,
		teacherID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var isExist = false
	for _, teacher := range result {
		if len(teacher.TeachingClassSubject) > 0 {
			for _, classSubject := range teacher.TeachingClassSubject {
				if classSubject.ClassNameID == classNameIDParsed {
					isExist = true
					break
				}
			}
		}
	}
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Teacher is not teaching this class"})
		return
	}

	parsedSubjectID, err := strconv.ParseInt(subjectID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get year now
	var academicSemesterYear string
	yearNow := time.Now().Year()

	// check if the semester 1 or 2 from the month
	if time.Now().Month() >= 7 {
		// concatenate the academic year
		academicSemesterYear = strconv.Itoa(yearNow) + "/" + strconv.Itoa(yearNow+1)
	} else {
		academicSemesterYear = strconv.Itoa(yearNow-1) + "/" + strconv.Itoa(yearNow)
	}

	// search for the assignment academic year if not exist create it
	var academicYear models.AcademicYear
	academicYear.AcademicYear = academicSemesterYear

	err = academicYear.GetAcademicYearByModel()
	if err != nil || academicYear.AcademicYearID == 0 {
		err = academicYear.CreateAcademicYear()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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

	err = models.CreateStudentsScoringBySubjectClassName(listScoring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send a success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Scoring data created successfully",
		"scoring": listScoring,
	})
}
