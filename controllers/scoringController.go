package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/helper"
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

	// TODO: use this function in helper/teacherHelper to check is the teacher teaching the class and subject
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

	resultDTO := helper.RemapScoringStudentBySubjectClassName(result)

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

	teacher, err := helper.IsTeachingClassSubjectExist(userID, subjectID, classNameID)
	if err != nil || teacher.TeacherID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedSubjectID, err := strconv.ParseInt(subjectID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get academic year
	academicYear, err := helper.GetOrCreateAcademicYear()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

func GetSummariesScoringStudentBySubjectClassName(c *gin.Context) {
	// Get parameters from the request
	classID := c.Param("class_id")

	if classID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "class_id and year are required"})
		return
	}

	// get or create academic year
	year, err := helper.GetOrCreateAcademicYear()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the scoring data from the model
	parsedAcademicYearID := strconv.FormatInt(year.AcademicYearID, 10)
	result, err := models.GetSummariesScoringStudentBySubjectClassName(classID, parsedAcademicYearID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultMap := helper.RemapScoringStudentBySubjectClassName(result)

	// group result map to get the average score of subject_name
	var resultDTO []helper.DTOAllScoringBySubjectClassName
	for _, student := range resultMap {
		// group the score by subject_name
		groupedResult := make(map[string]int)
		countGroupedResult := make(map[string]int)
		var totalAssignment = 0
		for _, score := range student.Scores {
			totalAssignment++
			// if the subject_name doesn't exist in the map, initialize it
			if _, exists := groupedResult[score.SubjectName]; !exists {
				groupedResult[score.SubjectName] = 0
				countGroupedResult[score.SubjectName] = 0
			}
			groupedResult[score.SubjectName] += score.Score
			countGroupedResult[score.SubjectName]++
		}

		// calculate the average score
		var scores []helper.Score
		for subjectName, score := range groupedResult {
			// var totalAssignment = len(student.Scores)
			scores = append(scores, helper.Score{
				SubjectName: subjectName,
				Score:       score / countGroupedResult[subjectName],
			})
		}

		// append the result to the DTO
		resultDTO = append(resultDTO, helper.DTOAllScoringBySubjectClassName{
			StudentID:   student.StudentID,
			StudentName: student.StudentName,
			Scores:      scores,
		})
	}
	// Send the result as a response
	c.JSON(http.StatusOK, gin.H{"score": resultDTO})
	// c.JSON(http.StatusOK, gin.H{"score": resultMap})
}
