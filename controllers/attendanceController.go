package controllers

import (
	"net/http"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/helper"
	req "github.com/edulink-api/request/attendance"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func GetAllAttendanceMonthSummaryByClassID(c *gin.Context) {
	ClassID, Date, err := helper.GetHomeRoomTeacherByTeacherID(c)
	if err != nil {
		res.AbortUnauthorized(c)
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

func GetAllAttendanceYearSummaryByClassID(c *gin.Context) {
	ClassID, Date, err := helper.GetHomeRoomTeacherByTeacherID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Year := Date.Year()

	result, err := models.GetAllAttendanceYearSummaryByClassID(ClassID, Year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else if len(result) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No attendance found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attendance": result,
	})
}
