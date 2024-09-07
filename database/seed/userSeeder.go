package seed

import (
	"time"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration"
	"github.com/skripsi-be/lib"
)

// init initializes the package by loading environment variables and connecting to the database.
func init() {
	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

// ClassSeeder seeds the Class data into the database.
func UserSeeder() {
	classes := []migration.User{
		{
			UserName: "guru1",
			UserGender: "male",
			UserPlaceOfBirth: "Jakarta",
			UserDateOfBirth: time.Now(),
			UserAddress: "Jl. Jakarta",
			UserNumPhone: "08123456789",
			UserEmail: "test@gmail.com",
			UserPassword: "123456",
		},
	}

	for _, class := range classes {
		connections.DB.Create(&class)
	}
}
