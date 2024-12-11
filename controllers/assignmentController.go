package controllers

import (
	"net/http"

	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/assignment"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func CreateAssignmentType(c *gin.Context) {
	var request request.InsertAssignmentRequest
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

	var assignment models.Assignment
	assignment.TypeAssignment = request.TypeAssignment

	// check if the assignment type already exists
	result, err := assignment.GetAssignmentByType()
	if err == nil && result.TypeAssignment == request.TypeAssignment {
		c.JSON(http.StatusOK, gin.H{"assignment": result})
		return
	}

	// Create the assignment type
	result, err = assignment.CreateAssignmentType()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"assignment": result})
}

func GetAllAssignmentType(c *gin.Context) {
	var assignment models.Assignment
	result, err := assignment.GetAllAssignmentType()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"assignment": result})
}
