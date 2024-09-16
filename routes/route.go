package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/controllers"
	"github.com/skripsi-be/middleware"
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
				middleware.IsAlreadyLogged, // try to implement this middleware
				controllers.DeleteStudentById,
			)
		}

		// Class CRUD
		class := apiV1.Group("/class")
		{
			class.GET(
				"/", 
				controllers.GetAllClass,
			)
			class.GET(
				"/:id", 
				controllers.GetClassById,
			)
			class.POST(
				"/create", 
				controllers.CreateClass,
			)
			class.PUT(
				"/update/:id", 
				controllers.UpdateClassById,
			)
			class.DELETE(
				"/delete/:id",
				controllers.DeleteClassById,
			)
		}

		// authentication
		auth := apiV1.Group("/auth")
		{
			auth.POST(
				"/login",
				controllers.Login,
			)
		}
	}
}
