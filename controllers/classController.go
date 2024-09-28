package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func CreateClass() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		class := models.Class{
			TeacherID:  request.TeacherID,
			ClassName:  request.ClassName,
			ClassGrade: request.ClassGrade,
		}

		// Insert the class into the database
		if err := connections.DB.Create(&class).Error; err != nil {
			// Handle different types of database errors
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create class",
				"error":   err.Error(),
			})
			return
		}

		// Return the newly created class as the response
		c.JSON(http.StatusOK, gin.H{
			"message": "Class created successfully",
			"class":   class,
		})
	}
}


func GetAllClass() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func GetClassById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("class_id")

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
}


func UpdateClassById() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		result := connections.DB.First(&class, c.Param("class_id"))
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Class not found",
			})
			return
		}

		// update class if exist
		connections.DB.Model(&class).Updates(models.Class{
			ClassName:  request.ClassName,
			TeacherID:  request.TeacherID,
			ClassGrade: request.ClassGrade,
		})

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
}

func DeleteClassById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var class models.Class
		result := connections.DB.First(&class, c.Param("class_id"))
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
}
