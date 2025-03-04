package seed

import (
	"time"

	"github.com/edulink-api/lib"
	"github.com/edulink-api/database/models"
)

// ClassSeeder seeds the Class data into the database.
func UserSeeder() (users []models.User) {
	users = []models.User{
		{
			UserName:         "guru1 math",
			UserGender:       "Male",
			UserPlaceOfBirth: "Jakarta",
			UserDateOfBirth:  time.Now(),
			UserReligion:     "Islam",
			UserAddress:      "Jl. Jakarta no 1",
			UserPhoneNum:     "+628123456789",
			UserEmail:        "guru1@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "guru aja",
			UserGender:       "Female",
			UserPlaceOfBirth: "Jakarta",
			UserReligion:     "Kristen Katolik",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Raya Bogor",
			UserPhoneNum:     "+628123456798",
			UserEmail:        "guruadmin@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "Teacher Non Home Room",
			UserGender:       "Male",
			UserPlaceOfBirth: "Jakarta",
			UserReligion:     "Konghucu",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta no 2",
			UserPhoneNum:     "+628123456178",
			UserEmail:        "nothomeroom@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "guru2",
			UserGender:       "Male",
			UserPlaceOfBirth: "Jakarta",
			UserReligion:     "Kristen Protestan",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta no 3",
			UserPhoneNum:     "+628123457771",
			UserEmail:        "guru2@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "Staff1",
			UserGender:       "Male",
			UserPlaceOfBirth: "Jakarta",
			UserReligion:     "Kristen Katolik",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta no 4",
			UserPhoneNum:     "+628123456782",
			UserEmail:        "staff1@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "Staff2",
			UserGender:       "Female",
			UserPlaceOfBirth: "Jakarta Barat",
			UserReligion:     "Kristen Protestan",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta no 5",
			UserPhoneNum:     "+628123451772",
			UserEmail:        "staff2@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "Test Admin",
			UserGender:       "Female",
			UserPlaceOfBirth: "Jakarta Barat",
			UserReligion:     "Kristen Protestan",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta no 5",
			UserPhoneNum:     "+628123451712",
			UserEmail:        "testadmin@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},
		{
			UserName:         "Admin2",
			UserGender:       "Female",
			UserPlaceOfBirth: "Jakarta Barat",
			UserReligion:     "Kristen Protestan",
			UserDateOfBirth:  time.Now(),
			UserAddress:      "Jl. Jakarta no 5",
			UserPhoneNum:     "+628123451722",
			UserEmail:        "admin2@gmail.com",
			UserPassword:     lib.HashPassword("123456"),
		},

	}

	return users
}
