package controllers

import (
	"net/http"
	"time"

	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllAttendanceByClassID() gin.HandlerFunc {
	return func(c *gin.Context) {

		// var student models.StudentModel
		ClassID := c.Param("class_id")

		Date, err := time.Parse("2006-01-02", c.Param("date"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := models.GetAllStudentsAttendanceByClassIDAndDate(ClassID, Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})

		// if err != "" {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": msg})
		// 	return
		// }

		// c.JSON(http.StatusOK, gin.H{"data": attendances})
	}
}
