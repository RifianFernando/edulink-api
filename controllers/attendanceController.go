package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/edulink-api/models"
	req "github.com/edulink-api/request/attendance"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func getHomeRoomTeacherByTeacherID(c *gin.Context) (string, time.Time, error) {
	// get user role
	userRole, _ := c.Get("user_type")
	userID, _ := c.Get("user_id")

	Date, err := time.Parse("2006-01-02", c.Param("date"))
	// var student models.StudentModel
	ClassID := c.Param("class_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return "", time.Time{}, err
	}
	if userRole != "admin" && userRole != "staff" {
		// check homeroom teacher class that he is assigned to
		var teacher models.Teacher
		teacher.UserID = userID.(int64)
		err := teacher.GetTeacherByModel()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			return "", time.Time{}, err
		}

		// get class id
		var className models.ClassName
		className.TeacherID = teacher.TeacherID
		err = className.GetHomeRoomTeacherByTeacherID()
		if err != nil || className.ClassNameID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return "", time.Time{}, err
		}

		// convert to string
		ClassID = strconv.FormatInt(className.ClassNameID, 10)
	}

	return ClassID, Date, nil
}

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

	// get user role
	userRole, _ := c.Get("user_type")
	userID, _ := c.Get("user_id")

	// get date
	Date, err := time.Parse("2006-01-02", c.Param("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get class id
	ClassID := c.Param("class_id")
	if userRole != "admin" && userRole != "staff" {
		// check homeroom teacher class that he is assigned to
		var teacher models.Teacher
		teacher.UserID = userID.(int64)
		err := teacher.GetTeacherByModel()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			return
		}

		// get class id
		var className models.ClassName
		className.TeacherID = teacher.TeacherID
		err = className.GetHomeRoomTeacherByTeacherID()
		if err != nil || className.ClassNameID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return
		}

		// convert to string
		ClassID = strconv.FormatInt(className.ClassNameID, 10)
	}

	// get attendance
	var attendance models.Attendance
	err = c.ShouldBindJSON(&attendance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create attendance
	// err = models.CreateAttendance(ClassID, Date, attendance)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

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
		c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
		return
	}

	// Prepare students
	var attendances []models.UpdateClassDateAttendanceStudent
	for _, attendance := range request.AttendanceRequest {
		attendances = append(attendances, models.UpdateClassDateAttendanceStudent{
			StudentID:   attendance.StudentID,
			Reason: 	attendance.Reason,
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
