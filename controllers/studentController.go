package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

		// Bind the request JSON to the CreateAllStudentRequest struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "should bind json" + err.Error(),
			})
			return
		}

		// Validate the request
		if err := request.ValidateAllStudent(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "when validate" + err.Error(),
			})
			return
		}

		// return
		nameMap := make(map[string]bool)
		nisnMap := make(map[string]bool)
		numPhoneMap := make(map[string]bool)
		emailMap := make(map[string]bool)
		var students []models.Student
		for index, student := range request.InsertStudentRequest {
			index = index + 1
			// Check if StudentName is already in the map
			if nameMap[student.StudentName] {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Duplicate StudentName: " + student.StudentName + " on index: " + strconv.Itoa(index),
				})
				return
			}
			// Check if StudentNISN is already in the map
			if nisnMap[student.StudentNISN] {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Duplicate StudentNISN: " + student.StudentNISN + " on index: " + strconv.Itoa(index),
				})
				return
			}
			// Check if StudentNumPhone is already in the map
			if numPhoneMap[student.StudentNumPhone] {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Duplicate StudentNumPhone: " + student.StudentNumPhone + " on index: " + strconv.Itoa(index),
				})
				return
			}
			// Check if StudentEmail is already in the map
			if emailMap[student.StudentEmail] {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Duplicate StudentEmail: " + student.StudentEmail + " on index: " + strconv.Itoa(index),
				})
				return
			}

			// Parse dates (if needed) and construct student models
			DateOfBirth, err := time.Parse("2006-01-02", student.DateOfBirth)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid date format on index: " + strconv.Itoa(index),
				})
				return
			}
			AcceptedDate, err := time.Parse("2006-01-02", student.AcceptedDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid date format on index: " + strconv.Itoa(index),
				})
				return
			}

			// checking database name, nisn, number phone, and email if already exist
			var studentSearch = models.Student{
				StudentName: student.StudentName,
			}
			result, err := studentSearch.GetStudent()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			} else if result.StudentID != 0 || result.ClassID != 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Student named: " + student.StudentName + " already exist on index: " + strconv.Itoa(index),
				})
				return
			}

			studentSearch = models.Student{
				StudentNISN: student.StudentNISN,
			}
			result, err = studentSearch.GetStudent()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			} else if result.StudentID != 0 || result.ClassID != 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Student NISN: " + student.StudentNISN + " already exist on index: " + strconv.Itoa(index),
				})
				return
			}

			studentSearch = models.Student{
				StudentNumPhone: student.StudentNumPhone,
			}
			result, err = studentSearch.GetStudent()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			} else if result.StudentID != 0 || result.ClassID != 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Student number_phone: " + student.StudentNumPhone + " already exist on index: " + strconv.Itoa(index),
				})
				return
			}

			studentSearch = models.Student{
				StudentEmail: student.StudentEmail,
			}
			result, err = studentSearch.GetStudent()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			} else if result.StudentID != 0 || result.ClassID != 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Student email: " + student.StudentEmail + " already exist on index: " + strconv.Itoa(index),
				})
				return
			}

			// checking classId if exist on database
			var class models.Class
			classIDStr := strconv.FormatInt(student.ClassID, 10)
			resultClass, err := class.GetClassById(classIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			} else if resultClass.ClassID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Class with id: " + classIDStr + " doesn't exist on index: " + strconv.Itoa(index),
				})
				return
			}

			students = append(students, models.Student{
				ClassID:               student.ClassID,
				StudentName:           student.StudentName,
				StudentNISN:           student.StudentNISN,
				StudentGender:         student.StudentGender,
				StudentPlaceOfBirth:   student.StudentPlaceOfBirth,
				StudentDateOfBirth:    DateOfBirth,
				StudentReligion:       student.StudentReligion,
				StudentAddress:        student.StudentAddress,
				StudentNumPhone:       student.StudentNumPhone,
				StudentEmail:          student.StudentEmail,
				StudentAcceptedDate:   AcceptedDate,
				StudentSchoolOrigin:   student.StudentSchoolOrigin,
				StudentFatherName:     student.StudentFatherName,
				StudentFatherJob:      student.StudentFatherJob,
				StudentFatherNumPhone: student.StudentFatherNumPhone,
				StudentMotherName:     student.StudentMotherName,
				StudentMotherJob:      student.StudentMotherJob,
				StudentMotherNumPhone: student.StudentMotherNumPhone,
			})

			// Mark each field as seen in the map
			nameMap[student.StudentName] = true
			nisnMap[student.StudentNISN] = true
			numPhoneMap[student.StudentNumPhone] = true
			emailMap[student.StudentEmail] = true
		}

		// Create all students
		err := models.CreateAllStudents(students)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"students": students,
			})
		}
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
		} else if result == (models.Student{}) && err == nil {
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
