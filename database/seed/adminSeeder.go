package seed

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
)

// ClassSeeder seeds the Class data into the database.
func AdminSeeder() {
	admins := []models.Admin{
		{
			UserID:   2,
			Position: "admin",
		},
		{
			UserID:   3,
			Position: "admin",
		},
	}

	for _, admin := range admins {
		connections.DB.Create(&admin)
	}
}
