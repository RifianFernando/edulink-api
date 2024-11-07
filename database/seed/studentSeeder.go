package seed

import (
	"time"

	"github.com/edulink-api/models"
)

// ClassSeeder seeds the Class data into the database.
func StudentSeeder() (students []models.Student) {
	students = []models.Student{
		{
			ClassNameID:           1,
			StudentName:           "Siswa 1",
			StudentNISN:           "0987654321",
			StudentGender:         "Male",
			StudentPlaceOfBirth:   "Jakarta",
			StudentDateOfBirth:    time.Now(),
			StudentReligion:       "Islam",
			StudentAddress:        "Jl. Jakarta",
			StudentNumPhone:       "08123456789",
			StudentEmail:          "test@gmail.com",
			StudentAcceptedDate:   time.Now(),
			StudentSchoolOrigin:   "SMP 1",
			StudentFatherName:     "Ayah 1",
			StudentFatherJob:      "PNS",
			StudentFatherNumPhone: "08123456789",
			StudentMotherName:     "Ibu 1",
			StudentMotherJob:      "PNS",
			StudentMotherNumPhone: "08123456789",
		},
		{
			ClassNameID:           1,
			StudentName:           "Siswa 2",
			StudentNISN:           "0987654312",
			StudentGender:         "Male",
			StudentPlaceOfBirth:   "Jakarta",
			StudentDateOfBirth:    time.Now(),
			StudentReligion:       "Islam",
			StudentAddress:        "Jl. Jakarta",
			StudentNumPhone:       "08123456798",
			StudentEmail:          "test2@gmail.com",
			StudentAcceptedDate:   time.Now(),
			StudentSchoolOrigin:   "SMP 1",
			StudentFatherName:     "Ayah 1",
			StudentFatherJob:      "PNS",
			StudentFatherNumPhone: "08123456789",
			StudentMotherName:     "Ibu 1",
			StudentMotherJob:      "PNS",
			StudentMotherNumPhone: "08123456789",
		},
	}

	return students
}
