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

const parseStaffIDError = "while parse staff id: "

func handleValidationErrors(c *gin.Context, errors []map[string]string) bool {
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return true
	}
	return false
}

func bindAndValidateRequest(c *gin.Context, req interface{}, validateFunc func() []map[string]string) ([]map[string]string, bool) {
	var allErrors []map[string]string
	if err := res.ResponseMessage(c.ShouldBindJSON(req)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}
	if err := validateFunc(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}
	if handleValidationErrors(c, allErrors) {
		return nil, true
	}
	return allErrors, false
}

func CreateStaff(c *gin.Context) {
	var req request.InsertStaffRequest
	if _, invalid := bindAndValidateRequest(c, &req, req.ValidateStaff); invalid {
		return
	}

	DateOfBirth, err := req.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	user := models.User{
		UserName:         req.UserName,
		UserGender:       req.UserGender,
		UserPlaceOfBirth: req.UserPlaceOfBirth,
		UserDateOfBirth:  DateOfBirth,
		UserReligion:     req.UserReligion,
		UserAddress:      req.UserAddress,
		UserPhoneNum:     req.UserPhoneNum,
		UserEmail:        req.UserEmail,
		UserPassword: lib.HashPassword(
			req.UserEmail + req.DateOfBirth,
		),
	}

	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "while create user: " + err.Error()})
		return
	}

	staff := models.Staff{
		UserID:   user.UserID,
		Position: req.Position,
	}
	if err := staff.CreateStaff(); err != nil || staff.StaffID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "while create staff: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"staff": staff,
	})
}

func CreateAllStaff(c *gin.Context) {
	var req request.InsertAllStaffRequest
	if _, invalid := bindAndValidateRequest(c, &req, req.ValidateAllStaff); invalid {
		return
	}

	staffs, err := helper.PrepareStaffs(req.InsertStaffRequest, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateAllStaffs(staffs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staffs":     "Create all staffs",
		"staff-data": staffs,
	})
}

func GetAllStaff(c *gin.Context) {
	staff := models.StaffModel{}
	staffs, err := staff.GetAllStaffs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"staff": staffs})
}

func GetStaffByID(c *gin.Context) {
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

	staff := models.StaffModel{Staff: models.Staff{StaffID: staffIDParsed}}
	result, err := staff.GetStaffByModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"staff": result})
}

func UpdateStaffByID(c *gin.Context) {
	var req request.UpdateStaffRequest
	// Bind the request JSON to the InsertStaffRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "should bind json: " + err.Error(),
		})
		return
	}
	// Validate the request
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	DateOfBirth, err := req.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format: " + err.Error()})
		return
	}

	id := c.Param("staff_id")
	staffModel := models.StaffModel{}
	staffModel.StaffID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parseStaffIDError + err.Error()})
		return
	}

	if _, err = staffModel.GetStaffByModel(); err != nil || staffModel.StaffID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	staffModel.User = models.User{
		UserName:         req.UserName,
		UserGender:       req.UserGender,
		UserPlaceOfBirth: req.UserPlaceOfBirth,
		UserReligion:     req.UserReligion,
		UserDateOfBirth:  DateOfBirth,
		UserAddress:      req.UserAddress,
		UserPhoneNum:     req.UserPhoneNum,
		UserEmail:        req.UserEmail,
	}
	staffModel.Position = req.Position

	if err := staffModel.UpdateStaffByID(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     "Updated staff with ID: " + id,
		"DateOfBirth": DateOfBirth,
	})
}

func DeleteStaffByID(c *gin.Context) {
	id := c.Param("staff_id")
	staff := models.StaffModel{}

	staffID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parseStaffIDError + err.Error()})
		return
	}
	staff.StaffID = staffID

	if _, err := staff.GetStaffByModel(); err != nil || staff.StaffID == 0 || staff.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": parseStaffIDError + err.Error()})
		return
	}

	if err := staff.DeleteStaffByID(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := staff.User.DeleteUserByID(strconv.FormatInt(staff.UserID, 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"staff": "deleted staff with id " + id})
}
