package controllers

import (
	"net/http"

	req "github.com/edulink-api/request/schedule"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func GenerateAndCreateScheduleTeachingClassSubject(c *gin.Context) {
	var request req.InsertScheduleRequest
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

	c.JSON(http.StatusOK, gin.H{"schedule": request})
}
