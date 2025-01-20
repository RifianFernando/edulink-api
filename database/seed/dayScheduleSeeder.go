package seed

import (
	"github.com/edulink-api/database/models"
)

func DayScheduleSeeder() (DaySchedule []models.DaySchedule) {
	DaySchedule = []models.DaySchedule{
		{ // 1
			DayName: "Monday",
		},
		{ // 2
			DayName: "Tuesday",
		},
		{ // 3
			DayName: "Wednesday",
		},
		{ // 4
			DayName: "Thursday",
		},
		{ // 5
			DayName: "Friday",
		},
		{ // 6
			DayName: "Saturday",
		},
		{ // 7
			DayName: "Sunday",
		},
	}

	return DaySchedule
}
