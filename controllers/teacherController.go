package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/lib"
	"github.com/edulink-api/models"
	request "github.com/edulink-api/request/teacher"
	"github.com/gin-gonic/gin"
)

func CreateTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.InsertTeacherRequest
		var allErrors []map[string]string

		// Bind the request JSON to the CreateStudentRequest struct
		if err := lib.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
			allErrors = append(allErrors, err...)
		}

		// Validate the request
		if err := request.ValidateTeacher(); len(err) > 0 {
			allErrors = append(allErrors, err...)
		}

		if len(allErrors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": allErrors,
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
			UserReligion:     request.UserReligion,
			UserAddress:      request.UserAddress,
			UserNumPhone:     request.UserNumPhone,
			UserEmail:        request.UserEmail,
			UserPassword: lib.HashPassword(
				request.UserEmail + request.DateOfBirth,
			),
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
			TeachingHour: 0,
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
		var teachers models.TeacherModel
		result, err := teachers.GetAllUserTeachersWithUser()
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

		var teacher models.TeacherModel
		result, err := teacher.GetTeacherById(id)
		if err != nil || teacher.TeacherID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"teacher": result,
		})
	}
}

func UpdateTeacherById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.UpdateTeacherRequest
		// Bind the request JSON to the InsertTeacherRequest struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "should bind json: " + err.Error(),
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
		DateOfBirth, err := request.ParseDates()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format: " + err.Error(),
			})
			return
		}

		// Get teacher id from URL and parse it to int64
		id := c.Param("teacher_id")

		var teacher models.TeacherModel
		_, err = teacher.GetTeacherById(id)
		if err != nil || teacher.TeacherID == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// update teacher and user
		teacher.User.UserName = request.UserName
		teacher.User.UserGender = request.UserGender
		teacher.User.UserPlaceOfBirth = request.UserPlaceOfBirth
		teacher.User.UserReligion = request.UserReligion
		teacher.User.UserDateOfBirth = DateOfBirth
		teacher.User.UserAddress = request.UserAddress
		teacher.User.UserNumPhone = request.UserNumPhone
		teacher.User.UserEmail = request.UserEmail
		teacher.TeachingHour = request.TeachingHour

		err = teacher.UpdateTeacherById(&teacher)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{
			"success":     "Updated teacher with ID: " + id,
			"DateOfBirth": DateOfBirth,
		})
	}
}

func DeleteTeacherById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("teacher_id")

		var teacher models.TeacherModel
		// if teacher exist
		_, err := teacher.GetTeacherById(id)
		if err != nil || teacher.TeacherID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = teacher.DeleteTeacherById(strconv.FormatInt(teacher.UserID, 10))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = teacher.User.DeleteUserById(strconv.FormatInt(teacher.UserID, 10))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"teacher": "deleted teacher with id " + id,
		})
	}
}
