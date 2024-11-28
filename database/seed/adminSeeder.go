package seed

import (
	"github.com/edulink-api/models"
)

// ClassSeeder seeds the Class data into the database.
func AdminSeeder() (admins []models.Admin) {
	admins = []models.Admin{
		{
			UserID:   2,
			Position: "admin",
		},
		{
			UserID:   4,
			Position: "admin",
		},
	}

	return admins
}
