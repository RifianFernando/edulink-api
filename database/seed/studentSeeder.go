package seed

import (
	"time"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
)

// ClassSeeder seeds the Class data into the database.
func StudentSeeder() {
	students := []models.Student{
		{
			ClassID:               1,
			StudentName:           "Siswa 1",
			StudentNISN:           "0987654321",
			StudentGender:         "male",
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
	}

	for _, student := range students {
		connections.DB.Create(&student)
	}
}
