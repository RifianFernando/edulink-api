package routes

import (
	"github.com/edulink-api/controllers"
	"github.com/edulink-api/middleware"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// Initialize Version
	apiV1 := router.Group("/api/v1")
	{
		// Student CRUD
		student := apiV1.Group("/student")
		student.Use(middleware.AlreadyLoggedIn(), middleware.AdminOnly())
		{
			student.GET(
				"/",
				controllers.GetAllStudent(),
			)
			student.GET(
				"/:student_id",
				controllers.GetStudentById(),
			)
			student.POST(
				"/create",
				controllers.CreateStudent(),
			)
			student.POST(
				"/create-all",
				controllers.CreateAllStudent(),
			)
			student.PUT(
				"/update/:student_id",
				// middleware.HaveStore(),
				controllers.UpdateStudentById(),
			)
			student.DELETE(
				"/delete/:student_id",
				controllers.DeleteStudentById(),
			)
		}

		// Class CRUD
		class := apiV1.Group("/class")
		{
			class.GET(
				"/",
				controllers.GetAllClass(),
			)
			class.GET(
				"/:class_id",
				controllers.GetClassById(),
			)
			class.POST(
				"/create",
				controllers.CreateClass(),
			)
			class.PUT(
				"/update/:class_id",
				controllers.UpdateClassById(),
			)
			class.DELETE(
				"/delete/:class_id",
				controllers.DeleteClassById(),
			)
		}

		// authentication
		auth := apiV1.Group("/auth")
		{
			auth.POST(
				"/login",
				middleware.IsNotLoggedIn(),
				controllers.Login(),
			)
			auth.POST(
				"/logout",
				middleware.AlreadyLoggedIn(),
				controllers.Logout(),
			)
		}
	}
}
