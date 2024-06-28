package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func CreateStudent(c *gin.Context) {
	var request request.CreateStudentRequest

	// Bind the request JSON to the CreateStudentRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error should bind json": err.Error()})
		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validate": err.Error()})
		return
	}

	// Parse date strings to time.Time
	dateOfBirth, acceptedDate, err := request.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	// create student
	student := models.Student{
		Name:              request.Name,
		Gender:            request.Gender,
		PlaceOfBirth:      request.PlaceOfBirth,
		DateOfBirth:       dateOfBirth,
		Religion:          request.Religion,
		Address:           request.Address,
		NumberPhone:       request.NumberPhone,
		Email:             request.Email,
		AcceptedDate:      acceptedDate,
		SchoolOrigin:      request.SchoolOrigin,
		IDClass:           request.IDClass,
		FatherName:        request.FatherName,
		FatherJob:         request.FatherJob,
		FatherNumberPhone: request.FatherNumberPhone,
		MotherName:        request.MotherName,
		MotherJob:         request.MotherJob,
		MotherNumberPhone: request.MotherNumberPhone,
	}

	result := connections.DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error result": result.Error.Error()})
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}
