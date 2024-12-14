package seed

import (
	"github.com/edulink-api/database/models"
)

func ScheduleSeeder() (Schedule []models.Schedule) {
	// total learning hours in a day is 10 from 7:00 to 17:00
	var DayScheduleID int64 = 0
	var HourScheduleID int64 = 1
	for i := 0; i < 70; i++ {
		if i % 10 == 0 {
			DayScheduleID++
			HourScheduleID = 1
		}
		Schedule = append(Schedule, models.Schedule{
			DayScheduleID: DayScheduleID,
			HourScheduleID: HourScheduleID,
			AcademicYearID: 1,
		})
		HourScheduleID++
	}

	return Schedule
}
