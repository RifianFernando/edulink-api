package connections

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecToDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		// configure the database on here
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "public", // schema name
		// 	SingularTable: false,
		// },
	})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = DB.Exec("CREATE SCHEMA IF NOT EXISTS public").Error
	if err != nil {
		return err
	}

	err = DB.Exec("CREATE SCHEMA IF NOT EXISTS academic").Error
	if err != nil {
		return err
	}

	err = DB.Exec("CREATE SCHEMA IF NOT EXISTS administration").Error
	if err != nil {
		return err
	}

	return nil
}
