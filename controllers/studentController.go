package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func CreateStudent(c *gin.Context) {
	var request request.InsertStudentRequest

	// Bind the request JSON to the CreateStudentRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error should bind json": err.Error(),
		})
		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error when validate": err.Error(),
		})
		return
	}

	// Parse date strings to time.Time
	dateOfBirth, acceptedDate, err := request.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error result": result.Error.Error(),
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func GetAllStudent(c *gin.Context) {
	var students []models.Student
	result := connections.DB.Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	result := connections.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func UpdateStudentById(c *gin.Context) {
	var request request.InsertStudentRequest

	// Bind the request JSON to the UpdateStudentRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error should bind json": err.Error(),
		})
		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error when validate": err.Error(),
		})
		return
	}

	// Parse date strings to time.Time
	dateOfBirth, acceptedDate, err := request.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format",
		})
		return
	}

	// Get student by id
	var student models.Student
	// result := connections.DB.First(&student, request.ID)
	// if result.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": result.Error.Error(),
	// 	})
	// 	return
	// }

	// Update student
	student.Name = request.Name
	student.Gender = request.Gender
	student.PlaceOfBirth = request.PlaceOfBirth
	student.DateOfBirth = dateOfBirth
	student.Religion = request.Religion
	student.Address = request.Address
	student.NumberPhone = request.NumberPhone
	student.Email = request.Email
	student.AcceptedDate = acceptedDate
	student.SchoolOrigin = request.SchoolOrigin
	student.IDClass = request.IDClass
	student.FatherName = request.FatherName
	student.FatherJob = request.FatherJob
	student.FatherNumberPhone = request.FatherNumberPhone
	student.MotherName = request.MotherName
	student.MotherJob = request.MotherJob
	student.MotherNumberPhone = request.MotherNumberPhone
	connections.DB.Save(&student)
}

func DeleteStudentById(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	// if student exist
	if connections.DB.First(&student, id).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Student not found",
		})
		return
	}
	result := connections.DB.Delete(&student{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}
