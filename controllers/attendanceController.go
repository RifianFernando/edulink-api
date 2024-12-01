package controllers

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	req "github.com/edulink-api/request/attendance"
	"github.com/gin-gonic/gin"
)

func GetAllAttendanceMonthSummaryByClassID(c *gin.Context) {
	ClassID, Date, err := helper.GetHomeRoomTeacherByTeacherID(c)
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
	ClassID, Date, err := helper.GetHomeRoomTeacherByTeacherID(c)
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
	ClassID, Date, attendances := helper.HandleCreateUpdateStudentAttendance(c, req.AllAttendanceRequest{})

	result, err := models.GetAllStudentAttendanceDateByClassID(ClassID, Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(result) > 0 {
		if err := models.UpdateStudentAttendanceByClassIDAndDate(ClassID, Date, attendances); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Attendance updated successfully",
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
		"request": req.AllAttendanceRequest{},
		"Date":    Date,
		"ClassID": ClassID,
	})
}

func UpdateStudentAttendance(c *gin.Context) {
	ClassID, Date, attendances := helper.HandleCreateUpdateStudentAttendance(c, req.AllAttendanceRequest{})

	// update attendance by class id and date and student id
	if err := models.UpdateStudentAttendanceByClassIDAndDate(ClassID, Date, attendances); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Attendance updated successfully",
		"request": req.AllAttendanceRequest{},
		"Date":    Date,
	})
}
