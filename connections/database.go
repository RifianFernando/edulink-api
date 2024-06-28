package connections

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnvValue(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	} else if value == "" {
		log.Fatalf("Environment variable %s is empty", key)
	}
	return value
}

var DB *gorm.DB

func ConnecToDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		getEnvValue("DB_HOST"),
		getEnvValue("DB_USERNAME"),
		getEnvValue("DB_PASSWORD"),
		getEnvValue("DB_NAME"),
		getEnvValue("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database")
	}
}
