package routes

import (
	"os"

	controllers "github.com/edulink-api/controllers"
	authController "github.com/edulink-api/controllers/auth"
	"github.com/edulink-api/middleware"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// Initialize Version
	apiV1 := router.Group(os.Getenv("API_V1"))
	{
		// Student CRUD
		student := apiV1.Group("/student")
		student.Use(middleware.AlreadyLoggedIn())
		{
			student.GET(
				"/",
				controllers.GetAllStudent(),
			)
			student.GET(
				"/:student_id",
				controllers.GetStudentById(),
			)
		}
		studentAdminStaff := student.Group("/")
		studentAdminStaff.Use(middleware.AdminStaffOnly())
		{
			studentAdminStaff.POST(
				"/create",
				controllers.CreateStudent(),
			)
			studentAdminStaff.POST(
				"/create-all",
				controllers.CreateAllStudent(),
			)
			studentAdminStaff.PUT(
				"/update-all-student-class-id",
				controllers.UpdateManyStudentClassID(),
			)
			studentAdminStaff.PUT(
				"/update/:student_id",
				controllers.UpdateStudentById(),
			)
			studentAdminStaff.DELETE(
				"/delete/:student_id",
				controllers.DeleteStudentById(),
			)
		}

		// Teacher CRUD
		teacher := apiV1.Group("/teacher")
		teacher.Use(middleware.AlreadyLoggedIn(), middleware.AdminStaffOnly())
		{
			teacher.GET(
				"/",
				controllers.GetAllTeacher(),
			)
			teacher.GET(
				"/:teacher_id",
				controllers.GetTeacherById(),
			)
			teacher.POST(
				"/create",
				controllers.CreateTeacher(),
			)
			// teacher.POST(
			// 	"/create-all",
			// 	controllers.CreateAllTeacher(),
			// )
			teacher.PUT(
				"/update/:teacher_id",
				controllers.UpdateTeacherById(),
			)
			teacher.DELETE(
				"/delete/:teacher_id",
				controllers.DeleteTeacherById(),
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
				controllers.GetClassNameById(),
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
				authController.Login(),
			)
			auth.POST(
				"/logout",
				middleware.AlreadyLoggedIn(),
				authController.Logout(),
			)
			// validate access token
			auth.GET(
				"/validate-token",
				authController.ValidateAccessToken(),
			)
			auth.GET(
				"/refresh-token",
				middleware.IsNotLoggedIn(),
				authController.RefreshToken(),
			)
			auth.POST(
				"/forget-password",
				middleware.IsNotLoggedIn(),
				authController.ForgetPassword(),
			)
			auth.POST(
				"/reset-password",
				middleware.IsNotLoggedIn(),
				authController.ResetPassword(),
			)
		}

		// attendance
		attendance := apiV1.Group("/attendance")
		attendance.Use(middleware.AlreadyLoggedIn(), middleware.IsTeacherHomeRoom())
		attendance.GET(
			"/summary/:class_id/:date",
			controllers.GetAllAttendanceMonthSummaryByClassID(),
		)
		attendance.GET(
			"/all-student/:class_id/:date",
			controllers.GetAllStudentAttendanceDateByClassID(),
		)
	}
}
