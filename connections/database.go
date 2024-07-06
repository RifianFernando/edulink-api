package connections

import (
	"fmt"
	"log"

	"github.com/skripsi-be/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecToDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			lib.GetEnvValue("DB_HOST"),
			lib.GetEnvValue("DB_USERNAME"),
			lib.GetEnvValue("DB_PASSWORD"),
			lib.GetEnvValue("DB_NAME"),
			lib.GetEnvValue("DB_PORT"),
		)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database")
	}
}
