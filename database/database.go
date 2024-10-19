package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration"
	"github.com/skripsi-be/database/seed"
	"github.com/skripsi-be/lib"
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
	generateSession := flag.Bool("session:generate", false, "Run generate session key")
	flag.Parse()

	switch {
	case *migrateFlag:
		performMigrations()
	case *migrateFreshFlag:
		migrateFresh()
	case *seedFlag:
		runSeeders()
	case *generateSession:
		generateSessionKey()
	default:
		fmt.Println("No valid command provided. Use -migrate, -migrate-fresh, or -seed.")
		os.Exit(1)
	}
}

func generateSessionKey() {
	sessionKey := lib.GenerateMultipleRandomStrings(1, 32)[0]
	value, err := lib.SetEnvValue("SESSION_KEY", sessionKey)
	if err != nil {
		log.Fatalf("Failed to set session key: %v", err)
	}
	fmt.Println("SESSION_KEY:", value)
}

// save the migration to the one variable
var table = []interface{}{
	&migration.Subject{},
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
