package connections

import (
	"fmt"
	"log"

	"github.com/skripsi-be/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnecToDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		lib.GetEnvValue("DB_HOST"),
		lib.GetEnvValue("DB_USERNAME"),
		lib.GetEnvValue("DB_PASSWORD"),
		lib.GetEnvValue("DB_NAME"),
		lib.GetEnvValue("DB_PORT"),
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

// func ConnecToDB(tablePrefix string) (*gorm.DB) {
// 	db, err := ConnectDB(tablePrefix)

// 	if err != nil {
// 		panic("failed to connect database" + err.Error())
// 	}

// 	return db
// }
