package controllers

import (
	"net/http"

	"github.com/edulink-api/models"
	req "github.com/edulink-api/request/attendance"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func GetAllAttendanceMonthSummaryByClassID(c *gin.Context) {
	ClassID, Date, err := getHomeRoomTeacherByTeacherID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := models.GetAllAttendanceMonthSummaryByClassID(ClassID, Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attendance": result,
	})
}

func GetAllStudentAttendanceDateByClassID(c *gin.Context) {
	ClassID, Date, err := getHomeRoomTeacherByTeacherID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := models.GetAllStudentAttendanceDateByClassID(ClassID, Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attendance": result,
	})
}

func CreateStudentAttendance(c *gin.Context) {
	var request req.AllAttendanceRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateAllAttendance(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	ClassID, Date, err := getHomeRoomTeacherByTeacherID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get attendance
	var attendances []models.ClassDateAttendanceStudent
	for _, attendance := range request.AttendanceRequest {
		attendances = append(attendances, models.ClassDateAttendanceStudent{
			StudentID: attendance.StudentID,
			Reason:    attendance.Reason,
		})
	}

	result, err := models.GetAllStudentAttendanceDateByClassID(ClassID, Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if (len(result) > 0) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Attendance already exist",
		})
		return
	}


	// create attendance
	err = models.CreateStudentClassAttendance(ClassID, Date, attendances)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Attendance created successfully",
		"request": request,
		"Date":    Date,
		"ClassID": ClassID,
	})
}

func UpdateStudentAttendance(c *gin.Context) {
	var request req.AllAttendanceRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateAllAttendance(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	ClassID, Date, err := getHomeRoomTeacherByTeacherID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if ClassID != c.Param("class_id") {
		res.AbortUnauthorized(c)
		return
	}

	// Prepare students
	var attendances []models.ClassDateAttendanceStudent
	for _, attendance := range request.AttendanceRequest {
		attendances = append(attendances, models.ClassDateAttendanceStudent{
			StudentID: attendance.StudentID,
			Reason:    attendance.Reason,
		})
	}

	// update attendance by class id and date and student id
	if err := models.UpdateStudentAttendanceByClassIDAndDate(ClassID, Date, attendances); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Attendance updated successfully",
		"request": request,
		"Date":    Date,
	})
}
