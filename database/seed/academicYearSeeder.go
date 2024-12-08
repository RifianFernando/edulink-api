package seed

import "github.com/edulink-api/database/models"

// ClassSeeder seeds the Class data into the database.
func AcademicYearSeeder() (admins []models.AcademicYear) {
	admins = []models.AcademicYear{
		{
			AcademicYear: "2024/2025",
		},
		// {
		// 	AcademicYear: "2021/2022",
		// },
		// {
		// 	AcademicYear: "2022/2023",
		// },
		// {
		// 	AcademicYear: "2023/2024",
		// }
	}

	return admins
}
