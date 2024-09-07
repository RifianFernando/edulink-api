package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func CreateClass(c *gin.Context) {
	var request request.InsertClassRequest

	// Bind the request JSON to the CreateClassRequest struct
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

	// create class
	class := models.Class{
		IDTeacher: request.IDTeacher,
		Name: request.Name,
		Grade: request.Grade,
	}

	result := connections.DB.Create(&class)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class": class,
	})
}

func GetAllClass(c *gin.Context) {
	var classes []models.Class
	result := connections.DB.Find(&classes)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
	})
}

func GetClassById(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	result := connections.DB.First(&class, id)	
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Class not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class": class,
	})
}

func UpdateClassById(c *gin.Context) {
	var request request.InsertClassRequest

	// Bind the request JSON to the UpdateClassRequest struct
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

	// Get class by id
	var class models.Class
	result := connections.DB.First(&class, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Class not found",
		})
		return
	}

	// update class
	class.Name = request.Name
	class.IDTeacher = request.IDTeacher
	class.Name = request.Name
	class.Grade = request.Grade

	result = connections.DB.Save(&class)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class": class,
	})
}

func DeleteClassById(c *gin.Context) {
	var class models.Class
	result := connections.DB.First(&class, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Class not found",
		})
		return
	}

	result = connections.DB.Delete(&class)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Class deleted successfully",
	})
}