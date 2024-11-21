package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/edulink-api/models"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func GetAllAttendanceMonthSummaryByClassID() gin.HandlerFunc {
	return func(c *gin.Context) {

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
			return
		}
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
}

func GetAllStudentAttendanceDateByClassID() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			return
		}
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
}
