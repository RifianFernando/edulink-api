package seed

import (
	"time"

	"github.com/edulink-api/models"
)

// ClassSeeder seeds the Class data into the database.
func StudentSeeder() (students []models.Student) {
	students = []models.Student{
		{
			ClassNameID:              1,
			StudentName:              "Siswa 1",
			StudentNISN:              "0987654321",
			StudentGender:            "Male",
			StudentPlaceOfBirth:      "Jakarta",
			StudentDateOfBirth:       time.Now(),
			StudentReligion:          "Islam",
			StudentAddress:           "Jl. Jakarta barat no 1",
			StudentPhoneNumber:       "+62123456987",
			StudentEmail:             "test@gmail.com",
			StudentAcceptedDate:      time.Now(),
			StudentSchoolOfOrigin:    "SMP 1",
			StudentFatherName:        "Ayah 1",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+628123456789",
			StudentMotherName:        "Ibu 1",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456887",
		},
		{
			ClassNameID:              1,
			StudentName:              "Siswa 2",
			StudentNISN:              "0987654312",
			StudentGender:            "Male",
			StudentPlaceOfBirth:      "Jakarta",
			StudentDateOfBirth:       time.Now(),
			StudentReligion:          "Islam",
			StudentAddress:           "Jl. Jakarta no 2",
			StudentPhoneNumber:       "+62123456798",
			StudentEmail:             "test2@gmail.com",
			StudentAcceptedDate:      time.Now(),
			StudentSchoolOfOrigin:    "SMP 1",
			StudentFatherName:        "Ayah 2",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+62123456897",
			StudentMotherName:        "Ibu 2",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456987",
		},
		{
			ClassNameID:              1,
			StudentName:              "Siswa 3",
			StudentNISN:              "0987654313",
			StudentGender: 			  "Male",
			StudentPlaceOfBirth:      "Jakarta",
			StudentDateOfBirth:       time.Now(),
			StudentReligion:          "Islam",
			StudentAddress:           "Jl. Jakarta",
			StudentPhoneNumber:       "+62123456543",
			StudentEmail:             "doejohn@gmail.com",
			StudentAcceptedDate:      time.Now(),
			StudentSchoolOfOrigin:    "SMP 1",
			StudentFatherName:        "Ayah 3",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+62123458698",
			StudentMotherName:        "Ibu 4",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456789",
		},
		{
			ClassNameID:              1,
			StudentName:              "Arif",
			StudentNISN:              "0987654222",
			StudentGender: 			  "Male",
			StudentPlaceOfBirth:      "Jakarta",
			StudentDateOfBirth:       time.Now(),
			StudentReligion:          "Konghucu",
			StudentAddress:           "Jl. Jakarta",
			StudentPhoneNumber:       "+62123456098",
			StudentEmail:             "arif@gmail.com",
			StudentAcceptedDate:      time.Now(),
			StudentSchoolOfOrigin:    "SMP 1",
			StudentFatherName:        "Ayah 5",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+62123456999",
			StudentMotherName:        "Ibu 5",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456991",
		},
	}

	return students
}
