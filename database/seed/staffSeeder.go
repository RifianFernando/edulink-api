package seed

import (
	"github.com/edulink-api/database/models"
)

func StaffSeeder() (staffs []models.Staff) {
	staffs = []models.Staff{
		{
			UserID:   5,
			Position: "TU",
		},
	}

	return staffs
}
