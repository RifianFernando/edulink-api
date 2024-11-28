package seed

import (
	"time"

	"github.com/edulink-api/models"
)

// Helper function to create a student
func createStudent(
	classNameID int64,
	studentName, studentNISN, studentGender, studentPlaceOfBirth, studentReligion, studentAddress, studentPhoneNumber, studentEmail,
	studentSchoolOfOrigin, studentFatherName, studentFatherJob, studentFatherPhoneNumber,
	studentMotherName, studentMotherJob, studentMotherPhoneNumber string,
) models.Student {
	return models.Student{
		ClassNameID:              classNameID,
		StudentName:              studentName,
		StudentNISN:              studentNISN,
		StudentGender:            studentGender,
		StudentPlaceOfBirth:      studentPlaceOfBirth,
		StudentDateOfBirth:       time.Now(),
		StudentReligion:          studentReligion,
		StudentAddress:           studentAddress,
		StudentPhoneNumber:       studentPhoneNumber,
		StudentEmail:             studentEmail,
		StudentAcceptedDate:      time.Now(),
		StudentSchoolOfOrigin:    studentSchoolOfOrigin,
		StudentFatherName:        studentFatherName,
		StudentFatherJob:         studentFatherJob,
		StudentFatherPhoneNumber: studentFatherPhoneNumber,
		StudentMotherName:        studentMotherName,
		StudentMotherJob:         studentMotherJob,
		StudentMotherPhoneNumber: studentMotherPhoneNumber,
	}
}

// StudentSeeder seeds the Student data into the database.
func StudentSeeder() (students []models.Student) {
	students = []models.Student{
		createStudent(1, "Siswa 1", "0987654321", "Male", "Jakarta", "Islam", "Jl. Jakarta barat no 1", "+62123456987", "test@gmail.com",
			"SMP 1", "Ayah 1", "PNS", "+628123456789", "Ibu 1", "PNS", "+62123456887"),
		createStudent(1, "Siswa 2", "0987654312", "Male", "Jakarta", "Islam", "Jl. Jakarta no 2", "+62123456798", "test2@gmail.com",
			"SMP 1", "Ayah 2", "PNS", "+62123456897", "Ibu 2", "PNS", "+62123456987"),
		createStudent(1, "Siswa 3", "0987654313", "Male", "Jakarta", "Islam", "Jl. Jakarta", "+62123456543", "doejohn@gmail.com",
			"SMP 1", "Ayah 3", "PNS", "+62123458698", "Ibu 4", "PNS", "+62123456789"),
		createStudent(1, "Arif", "0987654222", "Male", "Jakarta", "Konghucu", "Jl. Jakarta", "+62123456098", "arif@gmail.com",
			"SMP 1", "Ayah 5", "PNS", "+62123456999", "Ibu 5", "PNS", "+62123456991"),
	}

	return students
}
