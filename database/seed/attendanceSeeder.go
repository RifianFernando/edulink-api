package seed

import (
	"time"

	"github.com/edulink-api/database/models"
)

func AttendanceSeeder() (attendance []models.Attendance) {
	attendance = []models.Attendance{
		{
			StudentID:        1,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        2,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        3,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        1,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 8, 2, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        1,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        1,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 7, 2, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Absent",
		},
		{
			StudentID:        4,
			ClassNameID:      1,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Absent",
		},
		{
			StudentID:        5,
			ClassNameID:      4,
			AttendanceDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},

		// seeder for testing Class Archive
		{
			StudentID:        41,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        42,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        43,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        44,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 2, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        45,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 6, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        46,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 7, 2, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Absent",
		},
		{
			StudentID:        47,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Absent",
		},
		{
			StudentID:        48,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        49,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
		{
			StudentID:        50,
			ClassNameID:      8,
			AttendanceDate:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local),
			AttendanceStatus: "Present",
		},
	}

	return attendance
}
