package seed

import (
	"github.com/edulink-api/database/models"
)

func HourScheduleSeeder() (HourSchedule []models.HourSchedule) {
	// total learning hours in a day is 10 from 7:00 to 17:00
	for i := 7; i < 17; i++ {
		HourSchedule = append(HourSchedule, models.HourSchedule{
			StartHour: i,
			EndHour:   i + 1,
		})
	}

	return HourSchedule
}
