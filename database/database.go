package main

import (
	"flag"
	"log"
	"os"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration"
	"github.com/edulink-api/database/seed"
	"github.com/edulink-api/lib"
)

func init() {
	connections.LoadEnvVariables()

	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run the migrations")
	migrateFreshFlag := flag.Bool("migrate:fresh", false, "Drop all tables and run the migrations")
	seedFlag := flag.Bool("seed", false, "Run the seeders")
	keyGenerate := flag.Bool("key:generate", false, "Run generate app key")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	switch {
	case *migrateFlag:
		performMigrations()
	case *migrateFreshFlag:
		migrateFresh()
	case *seedFlag:
		runSeeders()
	case *keyGenerate:
		generateAppKey()
	case *help:
		flag.PrintDefaults()
	default:
		log.Fatalf("No valid command provided. Use -help for help")
		os.Exit(1)
	}
}

func generateAppKey() {
	appKey := lib.GenerateMultipleRandomStrings(1, 32)[0]
	value, err := lib.SetEnvValue("APP_KEY", appKey)
	if err != nil {
		log.Fatalf("Failed to set app key: %v", err)
	}
	log.Printf("app key set to: %s", value)
}

// save the migration to the one variable
var table = []interface{}{
	&migration.AcademicYear{},
	&migration.Subject{},
	&migration.Grade{},
	&migration.ClassName{},
	&migration.User{},
	&migration.Session{},
	&migration.ClassName{},
	&migration.Assignment{},
	&migration.Student{},
	&migration.Syllabus{},
	&migration.ContentObjective{},
	&migration.DomainAchievement{},
	&migration.SyllabusDetail{},
	&migration.Score{},
	&migration.Report{},
	&migration.Teacher{},
	&migration.TeacherSubject{},
	&migration.Schedule{},
	&migration.LearningSchedule{},
	&migration.EventSchedule{},
	&migration.Attendance{},
	&migration.AttendanceSummary{},
	&migration.Staff{},
	&migration.Admin{},
}

func performMigrations() {
	err := connections.DB.AutoMigrate(table...)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}

func migrateFresh() {
	err := connections.DB.Migrator().DropTable(table...)

	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	performMigrations()
}

func runSeeders() {
	seed.UserSeeder()
	seed.TeacherSeeder()
	seed.ClassSeeder()
	seed.StudentSeeder()
	seed.AdminSeeder()
}
