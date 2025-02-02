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
		student.GET("/class/:class_id", middleware.IsTeacherHomeRoom(), controllers.GetAllStudentByClassNameID)
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
		teacher.POST("/create-all", controllers.CreateAllTeacher)
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
		auth.GET("/user-type", middleware.AlreadyLoggedIn(), authController.GetUserType)
	}

	// Attendance
	attendance := apiV1.Group("/attendance", middleware.AlreadyLoggedIn(), middleware.IsTeacherHomeRoom())
	{
		attendance.GET("/summary/:class_id/:date", controllers.GetAllAttendanceMonthSummaryByClassID)
		attendance.GET("/summaries/:class_id/:year", controllers.GetAllAttendanceYearSummaryByClassID)
		attendance.GET("/all-student/:class_id/:date", controllers.GetAllStudentAttendanceDateByClassID)
		attendance.POST("/:class_id", controllers.CreateStudentAttendance)
		attendance.PUT("/:class_id/:date", controllers.UpdateStudentAttendance)
	}

	// Subject
	subject := apiV1.Group("/subject", middleware.AlreadyLoggedIn())
	{
		subject.GET("", middleware.AdminOnly(), controllers.GetAllSubject)
		// subject.GET("/class", controllers.GetAllSubjectClassName)
		subject.GET("/:subject_id/:class_id", controllers.GetSubjectClassNameStudentsByID)
	}

	// Scoring
	scoring := apiV1.Group("/scoring", middleware.AlreadyLoggedIn())
	{
		scoring.GET(
			"/summaries/:class_id",
			controllers.GetSummariesScoringStudentBySubjectClassName,
		)
		scoring.GET(
			"/get-all-class-teaching-subject-teacher",
			controllers.GetAllClassTeachingSubjectTeacher,
		)
	}
	scoringOnlyTeacher := scoring.Group("/", middleware.OnlyTeacher())
	{
		const CRUDScoring = "/:subject_id/:class_name_id"
		scoringOnlyTeacher.POST(CRUDScoring, controllers.CreateStudentsScoringBySubjectClassName)
		scoringOnlyTeacher.GET(CRUDScoring, controllers.GetAllScoringBySubjectClassName)
		scoring.PUT(CRUDScoring, controllers.UpdateScoringBySubjectClassName)
	}

	// assignment
	assignment := apiV1.Group("/assignment", middleware.AlreadyLoggedIn(), middleware.OnlyTeacher())
	{
		assignment.POST("", controllers.CreateAssignmentType)
		// no need to get assignment type because user will be create it, and if asg is exist, it will return the existing one from the db
		// assignment.GET("", controllers.GetAllAssignmentType)
	}

	// generator schedule
	generatorSchedule := apiV1.Group("/generator-schedule", middleware.AlreadyLoggedIn(), middleware.AdminStaffOnly())
	{
		generatorSchedule.POST("", controllers.GenerateAndCreateScheduleTeachingClassSubject)
	}

	// staff CRUD
	staff := apiV1.Group("/staff", middleware.AlreadyLoggedIn(), middleware.AdminStaffOnly())
	{
		staff.GET("", controllers.GetAllStaff)
		staff.GET("/:staff_id", controllers.GetStaffByID)
		staff.POST(create, controllers.CreateStaff)
		staff.PUT("/update/:staff_id", controllers.UpdateStaffByID)
		staff.DELETE("/delete/:staff_id", controllers.DeleteStaffByID)
		staff.POST("/create-all", controllers.CreateAllStaff)
	}

	// profile
	profile := apiV1.Group("/profile", middleware.AlreadyLoggedIn())
	{
		profile.GET("", authController.GetUserProfile)
	}

	// get academic year list
	academicYear := apiV1.Group("/academic-year", middleware.AlreadyLoggedIn())
	{
		academicYear.GET("", controllers.GetAcademicYearList)
	}

	// get archive academic calendar
	archiveData := apiV1.Group("/archive", middleware.AlreadyLoggedIn(), middleware.AdminStaffOnly())
	{
		/*
			We use param academic_year_start and academic_year_end for readability instead of academic_year_id
			* because it's easier to understand the range of academic year
			* example: 2019/2020
			* */

		/*
		* archiveData student personal-data
		* example: /student-personal-data/2019/2020
		* */
		archiveData.GET("/student-personal-data/:academic_year_start/:academic_year_end", controllers.GetAllStudentPersonalDataArchive)

		/*
		* archiveData student attendance
		* example: /student-attendance/2019/2020
		* example: /student-attendance/2019/2020/1 -> with class_id
		* */
		archiveData.GET("/student-attendance/:academic_year_start/:academic_year_end", controllers.GetAllStudentAttendanceArchive)

		/*
		* archiveData student score
		* example: /student-score/2019/2020
		* example: /student-score/2019/2020/1 -> with class_id
		* */
		archiveData.GET("/student-score/:academic_year_start/:academic_year_end", controllers.GetAllStudentScoreArchive)
		archiveData.GET("/student-score/:academic_year_start/:academic_year_end/:class_id", controllers.GetAllStudentScoreArchive)

		/*
		* archiveData class and student list
		* example: /class/2019/2020
		* */
		archiveData.GET("/class/:academic_year_start/:academic_year_end/:grade_id", controllers.GetAllClassArchiveByGradeID)
	}

	// create event 
	event := apiV1.Group("/event", middleware.AlreadyLoggedIn(), middleware.AdminStaffOnly())
	{
		event.POST("", controllers.CreateEvent)
		event.PUT("/:event_id", controllers.UpdateEvent)
		event.DELETE("/:event_id", controllers.DeleteEvent)
	}
}
