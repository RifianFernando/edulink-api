package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/helper"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.InsertStudentRequest

		// Bind the request JSON to the CreateStudentRequest struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "should bind json" + err.Error(),
			})
			return
		}

		// Validate the request
		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "when validate" + err.Error(),
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
		var student = models.Student{
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

		err = student.CreateStudent()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "while create student: " + err.Error(),
			})
			return
		}

		// return it
		c.JSON(http.StatusOK, gin.H{
			"student": student,
		})
	}
}

func CreateAllStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.InsertAllStudentRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "should bind json: " + err.Error()})
			return
		}

		if err := request.ValidateAllStudent(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "validation error: " + err.Error()})
			return
		}

		students, err := helper.PrepareStudents(request.InsertStudentRequest, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := models.CreateAllStudents(students); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"students": students})
	}
}

func GetAllStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var students models.Student
		result, err := students.GetAllStudents()
		if err != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"students": result,
		})
	}
}

func GetStudentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("student_id")

		var student models.Student
		result, err := student.GetStudentById(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if result.StudentID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Student not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"student": result,
		})
	}
}

func UpdateStudentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.InsertStudentRequest

		// Bind the request JSON to the UpdateStudentRequest struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "should bind json" + err.Error(),
			})
			return
		}

		// Validate the request
		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
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
		id := c.Param("student_id")

		var student models.Student
		result, err := student.GetStudentById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		} else if result == (models.Student{}) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Student not found",
			})

			return
		}

		// update student if exist
		student.StudentID = result.StudentID
		student.ClassID = request.ClassID
		student.StudentName = request.StudentName
		student.StudentNISN = request.StudentNISN
		student.StudentGender = request.StudentGender
		student.StudentPlaceOfBirth = request.StudentPlaceOfBirth
		student.StudentDateOfBirth = dateOfBirth
		student.StudentReligion = request.StudentReligion
		student.StudentAddress = request.StudentAddress
		student.StudentNumPhone = request.StudentNumPhone
		student.StudentEmail = request.StudentEmail
		student.StudentAcceptedDate = acceptedDate
		student.StudentSchoolOrigin = request.StudentSchoolOrigin
		student.StudentFatherName = request.StudentFatherName
		student.StudentFatherJob = request.StudentFatherJob
		student.StudentFatherNumPhone = request.StudentFatherNumPhone
		student.StudentMotherName = request.StudentMotherName
		student.StudentMotherJob = request.StudentMotherJob
		student.StudentMotherNumPhone = request.StudentMotherNumPhone

		err = student.UpdateStudentById(&student)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "while updating" + err.Error(),
			})
			return
		}

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
		err := student.DeleteStudentById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"student": "deleted student with id " + id,
		})
	}
}
