package controllers

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/database/models"
	"github.com/edulink-api/request"
	"github.com/gin-gonic/gin"
)

func CreateClass(c *gin.Context) {
	var request request.InsertClassRequest

	// Bind the request JSON to the CreateClassRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	// Validate the request (assuming you have a custom validation method)
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"error":   err.Error(),
		})
		return
	}

	// Create class
	var className = models.ClassName{
		GradeID:   request.GradeID,
		TeacherID: request.TeacherID,
		Name:      request.Name,
	}

	err := className.CreateClassName()

	// Insert the class into the database
	if err != nil {
		// Handle different types of database errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create class",
			"error":   err,
		})
		return
	}

	// Return the newly created class as the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Class created successfully",
		"class":   className,
	})
}

func GetAllClass(c *gin.Context) {
	accessToken, error := helper.GetCookieValue(c, "access_token")
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error()})
		c.Abort()
		return
	}

	claims, msg := helper.ValidateToken(accessToken, "access_token")
	if msg != "" || claims == nil || error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		c.Abort()
		return
	}
	var userType = claims.User_type
	var result []models.ClassNameModel
	var err string
	var ClassName models.ClassNameModel
	if userType == "teacher" {
		uid := claims.UserID
		ClassName.Teacher.UserID = uid
	}
	result, err = ClassName.GetAllClassName()
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ClassName": result,
	})
}

func GetClassNameById(c *gin.Context) {
	id := c.Param("class_id")

	var class models.ClassName
	result, err := class.GetClassNameById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Class not found",
		})
		return
	} else if result.ClassNameID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "class doesn't exist",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class": result,
	})
}

func UpdateClassById(c *gin.Context) {
	var request request.InsertClassRequest

	// Bind the request JSON to the UpdateClassRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	var class models.ClassName
	class, err := class.GetClassNameById(c.Param("class_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Class not found",
		})
		return
	}

	class.Name = request.Name
	class.TeacherID = request.TeacherID
	class.GradeID = request.GradeID
	err = class.UpdateClassNameByObject()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class":   class,
		"message": "Class updated successfully",
	})
}

func DeleteClassById(c *gin.Context) {
	var class models.ClassName
	var id = c.Param("class_id")

	err := class.DeleteClassNameById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Class not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Class deleted successfully",
	})
}
