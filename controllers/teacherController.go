package controllers

import (
	"net/http"
	// "strconv"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/edulink-api/models"
	"github.com/edulink-api/request"
)

func CreateTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.InsertTeacherRequest

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
		DateOfBirth, err := request.ParseDates()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format",
			})
			return
		}

		var user = models.User{
			UserName:         request.UserName,
			UserGender:       request.UserGender,
			UserPlaceOfBirth: request.UserPlaceOfBirth,
			UserDateOfBirth:  DateOfBirth,
			UserAddress:      request.UserAddress,
			UserNumPhone:     request.UserNumPhone,
			UserEmail:        request.UserEmail,
			UserPassword:     request.UserPassword,
		}

		err = user.CreateUser()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "while create user: " + err.Error(),
			})
			return
		}

		// create teacher
		var teacher = models.Teacher{
			UserID:       user.UserID,
			TeachingHour: request.TeachingHour,
		}

		err = teacher.CreateTeacher()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "while create teacher: " + err.Error(),
			})
			return
		}

		// return it
		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"teacher": teacher,
		})
	}
}

// func CreateAllTeacher() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var request request.InsertAllTeacherRequest

// 		// Bind the request JSON to the CreateAllTeacherRequest struct
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "should bind json" + err.Error(),
// 			})
// 			return
// 		}

// 		// Validate the request
// 		if err := request.ValidateAllTeacher(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "when validate" + err.Error(),
// 			})
// 			return
// 		}

// 		// return
// 		nameMap := make(map[string]bool)
// 		nisnMap := make(map[string]bool)
// 		numPhoneMap := make(map[string]bool)
// 		emailMap := make(map[string]bool)
// 		var teachers []models.Teacher
// 		for index, teacher := range request.InsertTeacherRequest {
// 			index = index + 1
// 			// Check if TeacherName is already in the map
// 			if nameMap[teacher.TeacherName] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate TeacherName: " + teacher.TeacherName + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}
// 			// Check if TeacherID is already in the map
// 			if idMap[teacher.TeacherID] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate TeacherID: " + teacher.TeacherID + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}
// 			// Check if TeacherNumPhone is already in the map
// 			if numPhoneMap[teacher.TeacherNumPhone] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate TeacherNumPhone: " + teacher.TeacherNumPhone + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}
// 			// Check if TeacherEmail is already in the map
// 			if emailMap[teacher.TeacherEmail] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate TeacherEmail: " + teacher.TeacherEmail + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			// Parse dates (if needed) and construct teacher models
// 			DateOfBirth, err := time.Parse("2006-01-02", teacher.DateOfBirth)
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Invalid date format on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			// checking database name, id, number phone, and email if already exist
// 			var teacherSearch = models.Teacher{
// 				TeacherName: student.TeacherName,
// 			}
// 			result, err := teacherSearch.GetTeacher()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.TeacherID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Teacher named: " + teacher.TeacherName + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			teacherSearch = models.Teacher{
// 				TeacherID: teacher.TeacherID,
// 			}
// 			result, err = teacherSearch.GetTeacher()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.TeacherID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Teacher ID: " + teacher.TeacherID + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			teacherSearch = models.Teacher{
// 				TeacherNumPhone: teacher.TeacherNumPhone,
// 			}
// 			result, err = teacherSearch.GetTeacher()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.TeacherID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Teacher number_phone: " + teacher.TeacherNumPhone + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			teacherSearch = models.Teacher{
// 				TeacherEmail: teacher.TeacherEmail,
// 			}
// 			result, err = teacherSearch.GetTeacher()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.TeacherID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Teacher email: " + teacher.TeacherEmail + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			teachers = append(teachers, models.Teacher{
// 				// ClassID:               student.ClassID,
// 				// StudentName:           student.StudentName,
// 				// StudentNISN:           student.StudentNISN,
// 				// StudentGender:         student.StudentGender,
// 				// StudentPlaceOfBirth:   student.StudentPlaceOfBirth,
// 				// StudentDateOfBirth:    DateOfBirth,
// 				// StudentReligion:       student.StudentReligion,
// 				// StudentAddress:        student.StudentAddress,
// 				// StudentNumPhone:       student.StudentNumPhone,
// 				// StudentEmail:          student.StudentEmail,
// 				// StudentAcceptedDate:   AcceptedDate,
// 				// StudentSchoolOrigin:   student.StudentSchoolOrigin,
// 				// StudentFatherName:     student.StudentFatherName,
// 				// StudentFatherJob:      student.StudentFatherJob,
// 				// StudentFatherNumPhone: student.StudentFatherNumPhone,
// 				// StudentMotherName:     student.StudentMotherName,
// 				// StudentMotherJob:      student.StudentMotherJob,
// 				// StudentMotherNumPhone: student.StudentMotherNumPhone,
// 			})

// 			// Mark each field as seen in the map
// 			nameMap[teacher.TeacherName] = true
// 			idMap[teacher.TeacherID] = true
// 			numPhoneMap[teacher.TeacherNumPhone] = true
// 			emailMap[teacher.TeacherEmail] = true
// 		}

// 		// Create all teachers
// 		err := models.CreateAllTeachers(teachers)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err,
// 			})
// 			return
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{
// 				"teachers": teachers,
// 			})
// 		}
// 	}
// }

func GetAllTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var teachers models.Teacher
		result, err := teachers.GetAllUserTeachers()
		if err != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"teachers": result,
		})
	}
}

func GetTeacherById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("teacher_id")

		var teacher models.Teacher
		result, err := teacher.GetTeacherById(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if result.TeacherID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Teacher not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"teacher": result,
		})
	}
}

// func UpdateTeacherById() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var request request.InsertTeacherRequest

// 		// Bind the request JSON to the UpdateTeacherRequest struct
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "should bind json" + err.Error(),
// 			})
// 			return
// 		}

// 		// Validate the request
// 		if err := request.Validate(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		// Parse date strings to time.Time
// 		dateOfBirth, err := request.ParseDates()
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid date format",
// 			})

// 			return
// 		}

// 		// Get teacher by id
// 		id := c.Param("teacher_id")

// 		var teacher models.Teacher
// 		result, err := teacher.GetTeacherById(id)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		} else if result == (models.Teacher{}) && err == nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Teacher not found",
// 			})

// 			return
// 		}

// 		// update teacher if exist
// 		// student.StudentID = result.StudentID
// 		// student.ClassID = request.ClassID
// 		// student.StudentName = request.StudentName
// 		// student.StudentNISN = request.StudentNISN
// 		// student.StudentGender = request.StudentGender
// 		// student.StudentPlaceOfBirth = request.StudentPlaceOfBirth
// 		// student.StudentDateOfBirth = dateOfBirth
// 		// student.StudentReligion = request.StudentReligion
// 		// student.StudentAddress = request.StudentAddress
// 		// student.StudentNumPhone = request.StudentNumPhone
// 		// student.StudentEmail = request.StudentEmail
// 		// student.StudentAcceptedDate = acceptedDate
// 		// student.StudentSchoolOrigin = request.StudentSchoolOrigin
// 		// student.StudentFatherName = request.StudentFatherName
// 		// student.StudentFatherJob = request.StudentFatherJob
// 		// student.StudentFatherNumPhone = request.StudentFatherNumPhone
// 		// student.StudentMotherName = request.StudentMotherName
// 		// student.StudentMotherJob = request.StudentMotherJob
// 		// student.StudentMotherNumPhone = request.StudentMotherNumPhone

// 		err = teacher.UpdateTeacherById(&teacher)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "while updating" + err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"teacher": teacher,
// 		})
// 	}
// }

func DeleteTeacherById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("teacher_id")

		var teacher models.Teacher
		// if teacher exist
		err := teacher.DeleteTeacherById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"teacher": "deleted teacher with id " + id,
		})
	}
}
