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
			res.AbortUnauthorized(c)
			return
		}

		// get class id
		var className models.ClassName
		className.TeacherID = teacher.TeacherID
		classes, err := className.GetHomeRoomTeacherByTeacherID()
		if err != nil {
			res.AbortUnauthorized(c)
			return
		}

		if len(classes) == 0 {
			res.AbortUnauthorized(c)
			return
		}

		// check if there's any class that the teacher is assigned to
		var isAssigned bool
		for _, class := range classes {
			if strconv.Itoa(int(class.ClassNameID)) == ClassID {
				isAssigned = true
				break
			}
		}

		if !isAssigned {
			res.AbortUnauthorized(c)
			return
		}
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
		res.AbortUnauthorized(c)
		return
	}

	// Prepare students
	var attendances []models.UpdateClassDateAttendanceStudent
	for _, attendance := range request.AttendanceRequest {
		attendances = append(attendances, models.UpdateClassDateAttendanceStudent{
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
