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
	// Load environment variables and establish DB connection
	connections.LoadEnvVariables()

	err := connections.ConnecToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run the migrations")
	migrateFreshFlag := flag.Bool("migrate:fresh", false, "Drop all tables and run the migrations")
	seedFlag := flag.Bool("seed", false, "Run the seeders")
	keyGenerate := flag.Bool("key:generate", false, "Generate and set a new application key")
	helpFlag := flag.Bool("help", false, "Show help")
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
	case *helpFlag:
		flag.PrintDefaults()
	default:
		log.Println("No valid command provided. Use -help for help")
		os.Exit(1)
	}
}

func generateAppKey() {
	appKey := lib.GenerateMultipleRandomStrings(1, 32)[0]
	if _, err := lib.SetEnvValue("APP_KEY", appKey); err != nil {
		log.Printf("Application key successfully generated and set: %s", appKey)

		return
	}
	log.Printf("Application key successfully generated and set: %s", appKey)
}

// List of tables to migrate
var tables = []interface{}{
	&migration.AcademicYear{},
	&migration.Subject{},
	&migration.Grade{},
	&migration.User{},
	&migration.Session{},
	&migration.ClassName{},
	&migration.Student{},
	&migration.Assignment{},
	&migration.Score{},
	&migration.Teacher{},
	&migration.TeacherSubject{},
	&migration.TeachingClassSubject{},
	&migration.DaySchedule{},
	&migration.HourSchedule{},
	&migration.Schedule{},
	&migration.LearningSchedule{},
	&migration.EventSchedule{},
	&migration.Attendance{},
	&migration.Staff{},
	&migration.Admin{},
}

func performMigrations() {
	if err := connections.DB.AutoMigrate(tables...); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed successfully.")
}

func migrateFresh() {
	log.Println("Dropping all tables...")
	if err := connections.DB.Migrator().DropTable(tables...); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Tables dropped successfully. Running migrations...")
	performMigrations()
}

func runSeeders() {
	if err := seed.Seed(); err != nil {
		log.Fatalf("Failed to run seeders: %v", err)
	}
	log.Println("Seeders executed successfully.")
}
