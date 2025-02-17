package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/lib"
	request "github.com/edulink-api/request/personal-data/teacher"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	var request request.InsertTeacherRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateTeacher(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	// Parse date strings to time.Time
	DateOfBirth, err := request.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format",
		})
		return
	}

	var user = models.User{
		UserName:         request.UserName,
		UserGender:       request.UserGender,
		UserPlaceOfBirth: request.UserPlaceOfBirth,
		UserDateOfBirth:  DateOfBirth,
		UserReligion:     request.UserReligion,
		UserAddress:      request.UserAddress,
		UserPhoneNum:     request.UserPhoneNum,
		UserEmail:        request.UserEmail,
		UserPassword: lib.HashPassword(
			request.UserEmail + request.DateOfBirth,
		),
	}

	err = user.CreateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while create user: " + err.Error(),
		})
		return
	}

	// create teacher
	teachingHour, err := strconv.ParseInt(request.TeachingHour, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while parse teaching hour: " + err.Error(),
		})
		return
	}
	var teacher = models.Teacher{
		UserID:       user.UserID,
		TeachingHour: int32(teachingHour),
	}
	err = teacher.CreateTeacher()
	if err != nil || teacher.TeacherID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while create teacher: " + err.Error(),
		})
		return
	}

	// create teacher subject
	var teacherSubjects []models.TeacherSubject
	for _, subjectID := range request.TeachingSubject {
		teacherSubjects = append(teacherSubjects, models.TeacherSubject{
			TeacherID: teacher.TeacherID,
			SubjectID: subjectID,
		})
	}
	err = models.CreateTeacherSubject(teacherSubjects)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while creating teacher subjects: " + err.Error(),
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"teacher": teacher,
	})
}

func CreateAllTeacher(c *gin.Context) {
	// var req request.InsertAllTeacherRequest
	// if _, invalid := bindAndValidateRequest(c, &req, req.ValidateAllTeacher); invalid {
	// 	return
	// }

	// teacher, err := helper.PrepareTeachers(req.InsertTeacherRequest, c)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if err := models.CreateAllTeachers(teacher); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"teacher": "all teacher created",
	})
}

func GetAllTeacher(c *gin.Context) {
	var teachers models.TeacherModel
	result, err := teachers.GetAllUserTeachersWithUser()
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teachers": result,
	})
}

func GetTeacherById(c *gin.Context) {
	id := c.Param("teacher_id")

	var teacher models.TeacherModel
	result, err := teacher.GetTeacherById(id)
	if err != nil || teacher.TeacherID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacher": result,
	})
}

func UpdateTeacherById(c *gin.Context) {
	var request request.UpdateTeacherRequest
	// Bind the request JSON to the InsertTeacherRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "should bind json: " + err.Error(),
		})
		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse date strings to time.Time
	DateOfBirth, err := request.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format: " + err.Error(),
		})
		return
	}

	// Get teacher id from URL and parse it to int64
	id := c.Param("teacher_id")
	var teacher models.TeacherModel
	_, err = teacher.GetTeacherById(id)
	if err != nil || teacher.TeacherID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	teachingHour, err := strconv.ParseInt(request.TeachingHour, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while parse teaching hour: " + err.Error(),
		})
		return
	}

	// update teacher and user
	teacher.User.UserName = request.UserName
	teacher.User.UserGender = request.UserGender
	teacher.User.UserPlaceOfBirth = request.UserPlaceOfBirth
	teacher.User.UserReligion = request.UserReligion
	teacher.User.UserDateOfBirth = DateOfBirth
	teacher.User.UserAddress = request.UserAddress
	teacher.User.UserPhoneNum = request.UserPhoneNum
	teacher.User.UserEmail = request.UserEmail
	teacher.TeachingHour = int32(teachingHour)
	teacher.TeacherSubject = nil
	for _, subjectID := range request.TeachingSubject {
		teacher.TeacherSubject = append(teacher.TeacherSubject, models.TeacherSubject{
			TeacherID: teacher.TeacherID,
			SubjectID: subjectID,
		})
	}

	err = teacher.UpdateTeacherById(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"success":     "Updated teacher with ID: " + id,
		"DateOfBirth": DateOfBirth,
	})
}

func DeleteTeacherById(c *gin.Context) {
	id := c.Param("teacher_id")

	var teacher models.TeacherModel
	// if teacher exist
	_, err := teacher.GetTeacherById(id)
	if err != nil || teacher.TeacherID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = teacher.DeleteTeacherById()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = teacher.User.DeleteUserByID(strconv.FormatInt(teacher.UserID, 10))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacher": "deleted teacher with id " + id,
	})
}
