package controllers

import (
	"net/http"
	"strconv"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/helper"
	"github.com/edulink-api/lib"
	request "github.com/edulink-api/request/staff"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func CreateStaff(c *gin.Context) {
	var request request.InsertStaffRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStaffRequest struct
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
		"user":  user,
		"staff": staff,
	})
}

func CreateAllStaff(c *gin.Context) {
	var request request.InsertAllStaffRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStaffRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateAllStaff(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	staffs, err := helper.PrepareStaffs(request.InsertStaffRequest, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create all staffs
	err = models.CreateAllStaffs(staffs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staffs":     "Create all staffs",
		"staff-data": staffs,
	})
}

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

func GetStaffByID(c *gin.Context) {
	id := c.Param("staff_id")

	var staff models.StaffModel
	var err error
	staff.StaffID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while parse staff id: " + err.Error(),
		})
		return
	}
	err = staff.GetStaffByModel()
	if err != nil || staff.Staff.StaffID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staff": staff,
	})
}

func UpdateStaffByID(c *gin.Context) {
	var request request.UpdateStaffRequest
	// Bind the request JSON to the InsertStaffRequest struct
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

	// Get staff id from URL and parse it to int64
	id := c.Param("staff_id")
	var staffModel models.StaffModel
	staffModel.StaffID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while parse staff id: " + err.Error(),
		})
		return
	}

	err = staffModel.GetStaffByModel()
	if err != nil || staffModel.StaffID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update staff and user
	staffModel.User.UserName = request.UserName
	staffModel.User.UserGender = request.UserGender
	staffModel.User.UserPlaceOfBirth = request.UserPlaceOfBirth
	staffModel.User.UserReligion = request.UserReligion
	staffModel.User.UserDateOfBirth = DateOfBirth
	staffModel.User.UserAddress = request.UserAddress
	staffModel.User.UserPhoneNum = request.UserPhoneNum
	staffModel.User.UserEmail = request.UserEmail
	staffModel.Position = request.Position

	err = staffModel.UpdateStaffByID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"success":     "Updated staff with ID: " + id,
		"DateOfBirth": DateOfBirth,
	})
}

func DeleteStaffByID(c *gin.Context) {
	id := c.Param("staff_id")

	var staff models.StaffModel
	var err error
	staff.StaffID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while parse staff id: " + err.Error(),
		})
		return
	}

	err = staff.GetStaffByModel()
	if err != nil || staff.StaffID == 0 || staff.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "while parse staff id: " + err.Error(),
		})
		return
	}

	err = staff.DeleteStaffByID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = staff.User.DeleteUserByID(strconv.FormatInt(staff.UserID, 10))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staff": "deleted staff with id " + id,
	})
}
