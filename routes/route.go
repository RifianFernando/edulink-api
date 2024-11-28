package routes

import (
	"os"

	controllers "github.com/edulink-api/controllers"
	authController "github.com/edulink-api/controllers/auth"
	"github.com/edulink-api/middleware"
	"github.com/gin-gonic/gin"
)

var (
	create = "/create"
)

func Route(router *gin.Engine) {
	// Initialize Version
	apiV1 := router.Group(os.Getenv("API_V1"))
	if apiV1 == nil {
		panic("API_V1 environment variable is not set. Ensure it is defined in the .env file.")
	}

	// Student CRUD
	student := apiV1.Group("/student", middleware.AlreadyLoggedIn())
	{
		student.GET("", controllers.GetAllStudent)
		// student.GET("/", controllers.GetAllStudent)
		student.GET("/:student_id", controllers.GetStudentById)
	}
	studentAdminStaff := student.Group("/", middleware.AdminStaffOnly())
	{
		studentAdminStaff.POST(create, controllers.CreateStudent)
		studentAdminStaff.POST("/create-all", controllers.CreateAllStudent)
		studentAdminStaff.PUT("/update-all-student-class-id", controllers.UpdateManyStudentClassID)
		studentAdminStaff.PUT("/update/:student_id", controllers.UpdateStudentById)
		studentAdminStaff.DELETE("/delete/:student_id", controllers.DeleteStudentById)
	}

	// Teacher CRUD
	teacher := apiV1.Group("/teacher", middleware.AlreadyLoggedIn(), middleware.AdminStaffOnly())
	{
		teacher.GET("", controllers.GetAllTeacher)
		teacher.GET("/:teacher_id", controllers.GetTeacherById)
		teacher.POST(create, controllers.CreateTeacher)
		teacher.PUT("/update/:teacher_id", controllers.UpdateTeacherById)
		teacher.DELETE("/delete/:teacher_id", controllers.DeleteTeacherById)
	}

	// Class CRUD
	class := apiV1.Group("/class")
	{
		class.GET("", controllers.GetAllClass)
		class.GET("/:class_id", controllers.GetClassNameById)
		class.POST(create, controllers.CreateClass)
		class.PUT("/update/:class_id", controllers.UpdateClassById)
		class.DELETE("/delete/:class_id", controllers.DeleteClassById)
	}

	// Authentication
	auth := apiV1.Group("/auth")
	{
		auth.POST("/login", middleware.IsNotLoggedIn(), authController.Login)
		auth.POST("/logout", middleware.AlreadyLoggedIn(), authController.Logout)
		auth.GET("/validate-token", authController.ValidateAccessToken)
		auth.GET("/refresh-token", middleware.IsNotLoggedIn(), authController.RefreshToken)
		auth.POST("/forget-password", middleware.IsNotLoggedIn(), authController.ForgetPassword)
		auth.POST("/reset-password", middleware.IsNotLoggedIn(), authController.ResetPassword)
	}

	// Attendance
	attendance := apiV1.Group("/attendance", middleware.AlreadyLoggedIn(), middleware.IsTeacherHomeRoom())
	{
		attendance.GET("/summary/:class_id/:date", controllers.GetAllAttendanceMonthSummaryByClassID)
		attendance.GET("/all-student/:class_id/:date", controllers.GetAllStudentAttendanceDateByClassID)
	}
}
