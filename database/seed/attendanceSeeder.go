package seed

import (
	"time"

	"github.com/edulink-api/models"
)

func AttendanceSeeder() (attendance []models.Attendance) {
	attendance = []models.Attendance{
		{
			StudentID:        1,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        2,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        3,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        1,
			AttendanceDate:   time.Date(2024, 8, 2, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        1,
			AttendanceDate:   time.Date(2024, 7, 2, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Absent",
		},
	}

	return attendance
}
