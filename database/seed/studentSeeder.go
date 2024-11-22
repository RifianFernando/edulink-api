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
			StudentAddress:           "Jl. Jakarta",
			StudentPhoneNumber:       "+62123456789",
			StudentEmail:             "test@gmail.com",
			StudentAcceptedDate:      time.Now(),
			StudentSchoolOfOrigin:    "SMP 1",
			StudentFatherName:        "Ayah 1",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+628123456789",
			StudentMotherName:        "Ibu 1",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456789",
		},
		{
			ClassNameID:              1,
			StudentName:              "Siswa 2",
			StudentNISN:              "0987654312",
			StudentGender:            "Male",
			StudentPlaceOfBirth:      "Jakarta",
			StudentDateOfBirth:       time.Now(),
			StudentReligion:          "Islam",
			StudentAddress:           "Jl. Jakarta",
			StudentPhoneNumber:       "+62123456798",
			StudentEmail:             "test2@gmail.com",
			StudentAcceptedDate:      time.Now(),
			StudentSchoolOfOrigin:    "SMP 1",
			StudentFatherName:        "Ayah 1",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+62123456789",
			StudentMotherName:        "Ibu 1",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456789",
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
			StudentFatherName:        "Ayah 1",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+62123456789",
			StudentMotherName:        "Ibu 1",
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
			StudentFatherName:        "Ayah 1",
			StudentFatherJob:         "PNS",
			StudentFatherPhoneNumber: "+62123456999",
			StudentMotherName:        "Ibu 1",
			StudentMotherJob:         "PNS",
			StudentMotherPhoneNumber: "+62123456999",
		},
	}

	return students
}
