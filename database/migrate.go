package main

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration"
)

func init() {
	connections.LoadEnvVariables()
	connections.ConnecToDB()
}

func main() {
	connections.DB.AutoMigrate(
		&migration.Subject{}, 
		&migration.User{}, 
		&migration.Class{}, 
		&migration.Student{}, 
		&migration.Syllabus{}, 
		&migration.ContentObjective{}, 
		&migration.DomainAchievement{}, 
		&migration.SyllabusDetail{}, 
		&migration.Grade{}, 
		&migration.LearningSchedule{}, 
		&migration.Teacher{}, 
		&migration.TeacherSubject{}, 
		&migration.Schedule{}, 
		&migration.EventSchedule{}, 
		&migration.Attendance{}, 
		&migration.AttendanceSummary{}, 
		&migration.Staff{},
	)
	// publicDB := connections.ConnecToDB("public")
	// academicDB := connections.ConnecToDB("academic")
	// administrationDB := connections.ConnecToDB("administration")

	// publicDB.AutoMigrate(
	// 	&migration.User{},
	// 	&migration.Class{},
	// 	&migration.Subject{},
	// )

	// academicDB.AutoMigrate(
	// 	&migration.Student{},
	// 	&migration.Teacher{},
	// 	&migration.Grade{},
	// 	&migration.Assignment{},
	// 	&migration.TeacherSubject{},
	// 	&migration.LearningSchedule{},
	// 	&migration.Syllabus{},
	// 	&migration.ContentObjective{},
	// 	&migration.DomainAchievement{},
	// 	&migration.SyllabusDetail{},
	// 	&migration.DomainAchievement{},
	// )

	// administrationDB.AutoMigrate(
	// 	&migration.Staff{},
	// 	&migration.Attendance{},
	// 	&migration.AttendanceSummary{},
	// 	&migration.EventSchedule{},	
	// 	&migration.Schedule{},
	// )
}
