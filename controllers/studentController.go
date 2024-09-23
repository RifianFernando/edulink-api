package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		DateOfBirth, AcceptedDate, err := request.ParseDates()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format",
			})
			return
		}

		// create student
		student := models.Student{
			ClassID:               request.ClassID,
			StudentName:           request.StudentName,
			StudentNISN:           request.StudentNISN,
			StudentGender:         request.StudentGender,
			StudentPlaceOfBirth:   request.StudentPlaceOfBirth,
			StudentDateOfBirth:    DateOfBirth,
			StudentReligion:       request.StudentReligion,
			StudentAddress:        request.StudentAddress,
			StudentNumPhone:       request.StudentNumPhone,
			StudentEmail:          request.StudentEmail,
			StudentAcceptedDate:   AcceptedDate,
			StudentSchoolOrigin:   request.StudentSchoolOrigin,
			StudentFatherName:     request.StudentFatherName,
			StudentFatherJob:      request.StudentFatherJob,
			StudentFatherNumPhone: request.StudentFatherNumPhone,
			StudentMotherName:     request.StudentMotherName,
			StudentMotherJob:      request.StudentMotherJob,
			StudentMotherNumPhone: request.StudentMotherNumPhone,
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
}

func GetAllStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []models.Student
		result := connections.DB.Find(&students)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			})
			return
		} else if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No student found",
			})
			return
		}

		// return it

		c.JSON(http.StatusOK, gin.H{
			"students": students,
		})
	}
}

func GetStudentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("student_id")

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
}

func UpdateStudentById() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		result := connections.DB.First(&student, c.Param("student_id"))
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		// update student if exist
		connections.DB.Model(&student).Updates(models.Student{
			ClassID:               request.ClassID,
			StudentName:           request.StudentName,
			StudentGender:         request.StudentGender,
			StudentPlaceOfBirth:   request.StudentPlaceOfBirth,
			StudentDateOfBirth:    dateOfBirth,
			StudentReligion:       request.StudentReligion,
			StudentAddress:        request.StudentAddress,
			StudentNumPhone:       request.StudentNumPhone,
			StudentEmail:          request.StudentEmail,
			StudentAcceptedDate:   acceptedDate,
			StudentSchoolOrigin:   request.StudentSchoolOrigin,
			StudentFatherName:     request.StudentFatherName,
			StudentFatherJob:      request.StudentFatherJob,
			StudentFatherNumPhone: request.StudentFatherNumPhone,
			StudentMotherName:     request.StudentMotherName,
			StudentMotherJob:      request.StudentMotherJob,
			StudentMotherNumPhone: request.StudentMotherNumPhone,
		})

		c.JSON(http.StatusOK, gin.H{
			"student": student,
		})
	}
}

func DeleteStudentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("student_id")

		var student models.Student
		// if student exist
		if connections.DB.First(&student, id).Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Student not found",
			})
			return
		}

		result := connections.DB.Delete(&student, id)
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
}
