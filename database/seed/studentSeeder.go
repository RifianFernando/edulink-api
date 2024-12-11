package seed

import (
	"time"

	"github.com/edulink-api/database/models"
)

// StudentData defines the input data for creating a student
type StudentData struct {
	ClassNameID       int64
	StudentName       string
	StudentNISN       string
	StudentGender     string
	StudentPlaceOfBirth string
	StudentReligion   string
	StudentAddress    string
	StudentPhoneNumber string
	StudentEmail      string
	StudentSchoolOfOrigin string
	StudentFatherName string
	StudentFatherJob  string
	StudentFatherPhoneNumber string
	StudentMotherName string
	StudentMotherJob  string
	StudentMotherPhoneNumber string
}

// createStudent maps StudentData to a Student model
func createStudent(data StudentData) models.Student {
	return models.Student{
		ClassNameID:              data.ClassNameID,
		StudentName:              data.StudentName,
		StudentNISN:              data.StudentNISN,
		StudentGender:            data.StudentGender,
		StudentPlaceOfBirth:      data.StudentPlaceOfBirth,
		StudentDateOfBirth:       time.Now(),
		StudentReligion:          data.StudentReligion,
		StudentAddress:           data.StudentAddress,
		StudentPhoneNumber:       data.StudentPhoneNumber,
		StudentEmail:             data.StudentEmail,
		StudentAcceptedDate:      time.Now(),
		StudentSchoolOfOrigin:    data.StudentSchoolOfOrigin,
		StudentFatherName:        data.StudentFatherName,
		StudentFatherJob:         data.StudentFatherJob,
		StudentFatherPhoneNumber: data.StudentFatherPhoneNumber,
		StudentMotherName:        data.StudentMotherName,
		StudentMotherJob:         data.StudentMotherJob,
		StudentMotherPhoneNumber: data.StudentMotherPhoneNumber,
	}
}

var PhoneNumberMother = "+62123456999"

// StudentSeeder seeds the Student data into the database.
func StudentSeeder() (students []models.Student) {
	studentDataList := []StudentData{
		{1, "Siswa 1", "0987654321", "Male", "Jakarta", "Islam", "Jl. Jakarta barat no 1", "+62123456987", "test@gmail.com", "SMP 1", "Ayah 1", "PNS", "+628123456789", "Ibu 1", "PNS", "+62123456887"},
		{1, "Siswa 2", "0987654312", "Male", "Jakarta", "Islam", "Jl. Jakarta no 2", "+62123456798", "test2@gmail.com", "SMP 1", "Ayah 2", "PNS", "+62123456897", "Ibu 2", "PNS", "+62123456987"},
		{1, "Siswa 3", "0987654313", "Male", "Jakarta", "Islam", "Jl. Jakarta", "+62123456543", "doejohn@gmail.com", "SMP 1", "Ayah 3", "PNS", "+62123458698", "Ibu 4", "PNS", "+62123456789"},
		{1, "Arif", "0987654222", "Male", "Jakarta", "Konghucu", "Jl. Jakarta", "+62123456098", "arif@gmail.com", "SMP 1", "Ayah 5", "PNS", PhoneNumberMother, "Ibu 5", "PNS", "+62123456991"},
		{4, "Rifian", "0987652222", "Male", "Jakarta", "Konghucu", "Jl. Jakarta Selatan", "+628123456789", "rifian@gmail.com", "SMP 1", "Ayah 6", "PNS", PhoneNumberMother, "Ibu 6", "PNS", "+628223344556"},
		{3, "Raven", "0998877665", "Female", "Jakarta", "Islam", "Jl. Jakarta barat no 1", "+628765432101", "raven@gmail.com", "SMP 2", "Ayah 7", "PNS", PhoneNumberMother, "Ibu 7", "PNS", "+628334455667"},
		{2, "Charles", "9977776658", "Female", "Bogor", "Konghucu", "Jl. Jakarta Selatan No 2", "+628998877665", "ejun@gmail.com", "SMP 3", "Ayah 8", "PNS", PhoneNumberMother, "Ibu 8", "PNS", "+628445566778"},
		{3, "Michael", "9977776885", "Male", "Cengkareng", "Kristen Katolik", "Jl. Jakarta Barat No 3", "+628112233445", "diva@gmail.com", "SMP 4", "Ayah 9", "PNS", PhoneNumberMother, "Ibu 9", "PNS", "+628556677889"},
		{2, "Gracella", "9976777885", "Female", "Kelapa Gading", "Buddha", "Jl. Jakarta Utara No 4", "+628567890123", "grace@gmail.com", "SMP 5", "Ayah 10", "PNS", PhoneNumberMother, "Ibu 10", "PNS", "+628667788990"},
		{2, "Amabel", "9976777557", "Female", "Alsut", "Kristen Protestan", "Jl. Jakarta Utara No 5", "+628910111213", "abel@gmail.com", "SMP 6", "Ayah 11", "PNS", PhoneNumberMother, "Ibu 11", "PNS", "+62123456199"},
	}

	for _, data := range studentDataList {
		students = append(students, createStudent(data))
	}

	return students
}
