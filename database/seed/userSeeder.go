package seed

import (
	"time"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/models"
)

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
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "guru2",
			UserGender:       "female",
			UserPlaceOfBirth: "Jakarta",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Raya Bogor",
			UserNumPhone:     "08123456798",
			UserEmail:        "testguru2@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "Admin1",
			UserGender:       "male",
			UserPlaceOfBirth: "Jakarta",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta",
			UserNumPhone:     "08123456781",
			UserEmail:        "testadmin@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
	}

	for _, user := range users {
		connections.DB.Create(&user)
	}
}
