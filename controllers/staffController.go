package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/lib"
	request "github.com/edulink-api/request/personal-data/staff"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func GetAllStaff(c *gin.Context) {
	// get from models
	staff := models.StaffModel{}
	staffs, err := staff.GetAllStaff()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"staff": staffs,
	})
}

func GetStaffByID(c *gin.Context) {
	// Get the scoring data from the model
	staffID := c.Param("staff_id")
	if staffID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check your parameter"})
		return
	}

	staffIDParsed, err := strconv.ParseInt(staffID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff_id must be a number"})
		return
	}

	// get from models
	staff := models.StaffModel{
		Staff: models.Staff{
			StaffID: staffIDParsed,
		},
	}
	result, err := staff.GetStaffByModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"staff": result,
	})
}

func CreateStaff(c *gin.Context) {
	var request request.InsertStaffRequest
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

	// create staff
	var staff = models.Staff{
		UserID:   user.UserID,
		Position: request.Position,
	}
	err = staff.CreateStaff()
	if err != nil || staff.StaffID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while create staff: " + err.Error(),
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"staff": staff,
	})
}

func UpdateStaffByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Staff By ID",
	})
}

func DeleteStaffByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Staff By ID",
	})
}
