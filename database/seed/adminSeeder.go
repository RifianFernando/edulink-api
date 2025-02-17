package seed

import (
	"github.com/edulink-api/database/models"
	"github.com/edulink-api/database/user"
)

// ClassSeeder seeds the Class data into the database.
func AdminSeeder() (admins []models.Admin) {
	admins = []models.Admin{
		{
			UserID:   7,
			Position: user.Admin,
		},
		{
			UserID:   8,
			Position: user.Admin,
		},
	}

	return admins
}
