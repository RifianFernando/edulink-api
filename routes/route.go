package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/controllers"
	// "github.com/skripsi-be/middleware"
)

func Route(router *gin.Engine) {
	// Initialize Version
	apiV1 := router.Group("/api/v1")
	{

		// Student CRUD
		student := apiV1.Group("/student")
		{
			student.GET(
				"/", 
				controllers.GetAllStudent,
			)
			student.GET(
				"/:id", 
				controllers.GetStudentById,
			)
			student.POST(
				"/create", 
				controllers.CreateStudent,
			)
			student.PUT(
				"/update/:id", 
				// middleware.HaveStore(), 
				controllers.UpdateStudentById,
			)
			student.DELETE(
				"/delete/:id", 
				// middleware.HaveStore(), try to implement this middleware
				controllers.DeleteStudentById,
			)
		}
	}
}
