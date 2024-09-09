package seed

import (
	"time"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/models"
)

// init initializes the package by loading environment variables and connecting to the database.
func init() {
	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

// ClassSeeder seeds the Class data into the database.
func UserSeeder() {
	users := []models.User{
		{
			UserName:         "guru1",
			UserGender:       "male",
			UserPlaceOfBirth: "Jakarta",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta",
			UserNumPhone:     "08123456789",
			UserEmail:        "test@gmail.com",
			UserPassword:     "123456",
		},
	}

	for _, user := range users {
		connections.DB.Create(&user)
	}
}
