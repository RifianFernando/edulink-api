package controllers

import (
	"net/http"
	"github.com/edulink-api/lib"
	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/staff"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func CreateStaff(c *gin.Context) {
	var request request.InsertStaffRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateStaff(); len(err) > 0 {
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
		UserPhoneNum:     request.UserPhoneNum,
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


	var staff = models.Staff{
		UserID:       user.UserID,
	}
	err = staff.CreateStaff()
	if err != nil || staff.StaffId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while create staff: " + err.Error(),
		})
		return
	}



	// return it
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"staff": staff,
	})
}

// func CreateAllStaff() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var request request.InsertAllStaffRequest

// 		// Bind the request JSON to the CreateAllStaffRequest struct
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "should bind json" + err.Error(),
// 			})
// 			return
// 		}

// 		// Validate the request
// 		if err := request.ValidateAllStaff(); err != nil {
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
// 		var staffs []models.Staff
// 		for index, staff := range request.InsertStaffRequest {
// 			index = index + 1
// 			// Check if StaffName is already in the map
// 			if nameMap[staff.StaffName] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate StaffName: " + staff.StaffName + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}
// 			// Check if StaffID is already in the map
// 			if idMap[staff.StaffID] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate StaffID: " + staff.StaffID + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}
// 			// Check if StaffNumPhone is already in the map
// 			if numPhoneMap[staff.StaffNumPhone] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate StaffNumPhone: " + staff.StaffNumPhone + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}
// 			// Check if StaffEmail is already in the map
// 			if emailMap[staff.StaffEmail] {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Duplicate StaffEmail: " + staff.StaffEmail + " on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			// Parse dates (if needed) and construct staff models
// 			DateOfBirth, err := time.Parse("2006-01-02", staff.DateOfBirth)
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Invalid date format on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			// checking database name, id, number phone, and email if already exist
// 			var staffSearch = models.Staff{
// 				StaffName: student.StaffName,
// 			}
// 			result, err := staffSearch.GetStaff()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.StaffID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Staff named: " + staff.StaffName + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			staffSearch = models.Staff{
// 				StaffID: staff.StaffID,
// 			}
// 			result, err = staffSearch.GetStaff()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.StaffID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Staff ID: " + staff.StaffID + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			staffSearch = models.Staff{
// 				StaffNumPhone: staff.StaffNumPhone,
// 			}
// 			result, err = staffSearch.GetStaff()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.StaffID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Staff number_phone: " + staff.StaffNumPhone + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			staffSearch = models.Staff{
// 				StaffEmail: staff.StaffEmail,
// 			}
// 			result, err = staffSearch.GetStaff()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			} else if result.StaffID != 0 {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": "Staff email: " + staff.StaffEmail + " already exist on index: " + strconv.Itoa(index),
// 				})
// 				return
// 			}

// 			staffs = append(staffs, models.Staff{
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
// 			nameMap[staff.StaffName] = true
// 			idMap[staff.StaffID] = true
// 			numPhoneMap[staff.StaffNumPhone] = true
// 			emailMap[staff.StaffEmail] = true
// 		}

// 		// Create all staffs
// 		err := models.CreateAllStaffs(staffs)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err,
// 			})
// 			return
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{
// 				"staffs": staffs,
// 			})
// 		}
// 	}
// }

func GetAllStaff(c *gin.Context) {
	var staffs models.StaffModel
	result, err := staffs.GetAllUserStaffWithUser()
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"staffs": result,
	})
}

// func GetStaffById(c *gin.Context) {
// 	id := c.Param("staff_id")

// 	var staff models.Staff
// 	result, err := staff.GetStaffById(id)
// 	if err != nil || staff.StaffID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"staff": result,
// 	})
// }

// func UpdateStaffById(c *gin.Context) {
// 	var request request.UpdateStaffRequest
// 	// Bind the request JSON to the InsertStaffRequest struct
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "should bind json: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Validate the request
// 	if err := request.Validate(); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// Parse date strings to time.Time
// 	DateOfBirth, err := request.ParseDates()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid date format: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Get staff id from URL and parse it to int64
// 	id := c.Param("staff_id")
// 	var staff models.Staff
// 	_, err = staff.GetStaffById(id)
// 	if err != nil || staff.StaffID == 0 {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	teachingHour, err := strconv.ParseInt(request.TeachingHour, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "while parse teaching hour: " + err.Error(),
// 		})
// 		return
// 	}

// 	// update staff and user
// 	staff.User.UserName = request.UserName
// 	staff.User.UserGender = request.UserGender
// 	staff.User.UserPlaceOfBirth = request.UserPlaceOfBirth
// 	staff.User.UserReligion = request.UserReligion
// 	staff.User.UserDateOfBirth = DateOfBirth
// 	staff.User.UserAddress = request.UserAddress
// 	staff.User.UserPhoneNum = request.UserPhoneNum
// 	staff.User.UserEmail = request.UserEmail
// 	staff.TeachingHour = int32(teachingHour)
// 	staff.StaffSubject = nil
// 	for _, subjectID := range request.TeachingSubject {
// 		staff.StaffSubject = append(staff.StaffSubject, models.StaffSubject{
// 			StaffID: staff.StaffID,
// 			SubjectID: subjectID,
// 		})
// 	}

// 	err = staff.UpdateStaffById(&staff)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// Return success response
// 	c.JSON(http.StatusOK, gin.H{
// 		"success":     "Updated staff with ID: " + id,
// 		"DateOfBirth": DateOfBirth,
// 	})
// }

// func DeleteStaffById(c *gin.Context) {
// 	id := c.Param("staff_id")

// 	var staff models.Staff
// 	// if staff exist
// 	_, err := staff.GetStaffById(id)
// 	if err != nil || staff.StaffID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	err = staff.DeleteStaffById()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	err = staff.User.DeleteUserById(strconv.FormatInt(staff.UserID, 10))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"staff": "deleted staff with id " + id,
// 	})
// }
