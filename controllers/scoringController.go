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

var (
	userIDNotFound = "User ID not found"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": userIDNotFound})
		return
	}

	teacher, err := helper.IsTeachingClassSubjectExist(userID, subjectID, classNameID)
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
	listScoring, allErrors, err := helper.GetListScoring(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
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
}

func UpdateScoringBySubjectClassName(c *gin.Context) {
	listScoring, allErrors, err := helper.GetListScoring(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	err = models.UpdateScoringBySubjectClassName(listScoring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send a success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Scoring data updated successfully",
		"scoring": listScoring,
	})
}

func GetAllClassTeachingSubjectTeacher(c *gin.Context) {
	// Get parameters from the request
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": userIDNotFound})
		return
	}

	userIDParsed := strconv.FormatInt(userID.(int64), 10)
	classList, err := models.GetTeacherTeachingClassList(userIDParsed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type dtoClassList struct {
		SubjectID      int    `json:"subject_id"`
		ClassNameID    int    `json:"class_name_id"`
		GradeClassName string `json:"grade_class_name"`
		SubjectName    string `json:"subject_name"`
	}

	var classListDTO []dtoClassList

	for _, class := range classList {
		classListDTO = append(classListDTO, dtoClassList{
			SubjectID:      class.SubjectID,
			ClassNameID:    class.ClassNameID,
			GradeClassName: class.Grade + class.Name,
			SubjectName:    class.SubjectName,
		})
	}

	// Send the result as a response
	c.JSON(http.StatusOK, gin.H{"class_list": classListDTO})
}

func GetStudentScoresByStudentSubjectClassID(c *gin.Context) {
	// Get parameters from the request
	studentID := c.Param("student_id")
	subjectID := c.Param("subject_id")
	classNameID := c.Param("class_name_id")

	if subjectID == "" || classNameID == "" || studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject_id, class_name_id, and student_id are required"})
		return
	}

	// Get the scoring data from the model
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": userIDNotFound})
		return
	}

	teacher, err := helper.IsTeachingClassSubjectExist(userID, subjectID, classNameID)
	if err != nil || teacher.TeacherID == 0 {
		res.AbortUnauthorized(c)
		return
	}

	// Get the scoring data from the model
	var studentScores models.ScoreModel
	studentScores.StudentID, _ = strconv.ParseInt(studentID, 10, 64)
	studentScores.SubjectID, _ = strconv.ParseInt(subjectID, 10, 64)
	studentScores.ClassNameID, _ = strconv.ParseInt(classNameID, 10, 64)
	studentScores.TeacherID = teacher.TeacherID
	result, err := studentScores.GetStudentScoresAndTypeByStudentSubjectClassID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// dto
	type dtoStudentScores struct {
		StudentName string `json:"student_name"`
		SubjectName string `json:"subject_name"`
		Scores      []struct {
			AssignmentID   int64  `json:"assignment_id"`
			AssignmentType string `json:"assignment_type"`
			Score          int    `json:"score"`
		} `json:"scores"`
	}

	// group the result by student_name and subject_name
	resultDTO := dtoStudentScores{
		StudentName: result[0].Student.StudentName,
		SubjectName: result[0].Subject.SubjectName,
	}

	for _, score := range result {
		resultDTO.Scores = append(resultDTO.Scores, struct {
			AssignmentID   int64  `json:"assignment_id"`
			AssignmentType string `json:"assignment_type"`
			Score          int    `json:"score"`
		}{
			AssignmentID:   score.Assignment.AssignmentID,
			AssignmentType: score.Assignment.TypeAssignment,
			Score:          score.Score.Score,
		})
	}

	// Send the result as a response
	c.JSON(http.StatusOK, gin.H{"score": resultDTO})
}

func UpdateStudentScoresByStudentSubjectClassID(c *gin.Context) {
	var request request.UpdateStudentScoreRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.Validate(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	// Get the scoring data from the model
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": userIDNotFound})
		return
	}

	// Get parameters from the request
	subjectID := c.Param("subject_id")
	classNameID := c.Param("class_name_id")

	teacher, err := helper.IsTeachingClassSubjectExist(userID, subjectID, classNameID)
	if err != nil || teacher.TeacherID == 0 {
		res.AbortUnauthorized(c)
		return
	}

	// TODO: update the student scores by ID

	c.JSON(http.StatusOK, gin.H{"message": "success update student scores"})
}
