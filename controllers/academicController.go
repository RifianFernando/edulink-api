package controllers

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

func GetAcademicYearList(c *gin.Context) {
	// Get all academic year
	academicYear, err := helper.GetAcademicYearList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send the grouped result as a response
	c.JSON(http.StatusOK, gin.H{"academic-year": academicYear})
}
