package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		var className = models.ClassName{
			TeacherID:  request.TeacherID,
			ClassName:  request.ClassName,
			ClassGrade: request.ClassGrade,
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
}

func GetAllClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ClassName models.Class
		result, err := ClassName.GetAllClass()
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
}

func GetClassById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("class_id")

		var class models.Class
		result, err := class.GetClassById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Class not found",
			})
			return
		} else if result.ClassID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "class doesn't exist",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"class": result,
		})
	}
}

func UpdateClassById() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		var class models.Class
		class, err := class.GetClassById(c.Param("class_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Class not found",
			})
			return
		}

		class.ClassName = request.ClassName
		class.TeacherID = request.TeacherID
		class.ClassGrade = request.ClassGrade
		err = class.UpdateClassByObject()
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
}

func DeleteClassById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var class models.Class
		var id = c.Param("class_id")

		err := class.DeleteClassById(id)
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
}
