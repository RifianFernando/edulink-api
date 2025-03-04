package seed

import (
	"time"

	"github.com/edulink-api/database/models"
	"gorm.io/gorm"
)

// StudentData defines the input data for creating a student
type StudentData struct {
	ClassNameID              int64
	StudentName              string
	StudentNISN              string
	StudentGender            string
	StudentPlaceOfBirth      string
	StudentReligion          string
	StudentAddress           string
	StudentPhoneNumber       string
	StudentEmail             string
	StudentSchoolOfOrigin    string
	StudentFatherName        string
	StudentFatherJob         string
	StudentFatherPhoneNumber string
	StudentMotherName        string
	StudentMotherJob         string
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
		{5, "Amabel Smith", "9976777551", "Female", "Jl. Jakarta Selatan No 3", "Kristen Protestan", "Jl. Jakarta Utara No 5", "+628910111214", "amabelsmith@gmail.com", "SMP 6", "Ayah 1", "PNS", "+62123456192", "Ibu 1", "PNS", "+62123456202"},
		{5, "Nathaniel Jones", "9976777553", "Male", "Jl. Jakarta Timur No 4", "Kristen Protestan", "Jl. Jakarta Utara No 6", "+628910111215", "nathanjones@gmail.com", "SMP 6", "Ayah 2", "PNS", "+62123456193", "Ibu 2", "PNS", "+62123456203"},
		{5, "Mia Anderson", "9976777558", "Female", "Jl. Jakarta Timur No 9", "Kristen Protestan", "Jl. Jakarta Utara No 11", "+628910111220", "miaanderson@gmail.com", "SMP 6", "Ayah 7", "PNS", "+62123456198", "Ibu 7", "PNS", "+62123456207"},
		{5, "Ethan Garcia", "9976777554", "Male", "Jl. Jakarta Pusat No 6", "Kristen Protestan", "Jl. Jakarta Utara No 8", "+628910111216", "ethangarcia@gmail.com", "SMP 6", "Ayah 4", "PNS", "+62123456194", "Ibu 4", "PNS", "+62123456204"},
		{5, "Isabella Martinez", "9976777555", "Female", "Jl. Jakarta Utara No 7", "Kristen Protestan", "Jl. Jakarta Utara No 9", "+628910111217", "isabellamartinez@gmail.com", "SMP 6", "Ayah 5", "PNS", "+62123456195", "Ibu 5", "PNS", "+62123456205"},
		{5, "Liam Wilson", "9976777556", "Male", "Jl. Jakarta Selatan No 8", "Kristen Protestan", "Jl. Jakarta Utara No 10", "+628910111218", "liamwilson@gmail.com", "SMP 6", "Ayah 6", "PNS", "+62123456196", "Ibu 6", "PNS", "+62123456206"},
		{5, "Sophia Brown", "9976777559", "Female", "Jl. Jakarta Barat No 5", "Kristen Protestan", "Jl. Jakarta Utara No 7", "+628910111221", "sophiabrown@gmail.com", "SMP 6", "Ayah 3", "PNS", "+62123456194", "Ibu 3", "PNS", "+62123456204"},
		{5, "Noah Lee", "9976777560", "Male", "Jl. Jakarta Barat No 10", "Kristen Protestan", "Jl. Jakarta Utara No 12", "+628910111222", "noahlee@gmail.com", "SMP 6", "Ayah 8", "PNS", "+62123456199", "Ibu 8", "PNS", "+62123456209"},
		{5, "Olivia Harris", "9976777561", "Female", "Jl. Jakarta Pusat No 11", "Kristen Protestan", "Jl. Jakarta Utara No 13", "+628910111223", "oliviaharris@gmail.com", "SMP 6", "Ayah 9", "PNS", "+62123456201", "Ibu 9", "PNS", "+62123456211"},
		{5, "James Clark", "9976777562", "Male", "Jl. Jakarta Selatan No 12", "Kristen Protestan", "Jl. Jakarta Utara No 14", "+628910111224", "jamesclark@gmail.com", "SMP 6", "Ayah 10", "PNS", "+62123456202", "Ibu 10", "PNS", "+62123456212"},

		{6, "Charlotte Scott", "9976777516", "Female", "Jl. Jakarta Selatan No 13", "Kristen Protestan", "Jl. Jakarta Utara No 15", "+628910121232", "charlottescott@gmail.com", "SMP 6", "Ayah 11", "PNS", "+62123456201", "Ibu 11", "PNS", "+62123456211"},
		{6, "Benjamin King", "9976777526", "Male", "Jl. Jakarta Timur No 14", "Kristen Protestan", "Jl. Jakarta Utara No 16", "+628910123224", "benjaminking@gmail.com", "SMP 6", "Ayah 12", "PNS", "+62123456202", "Ibu 12", "PNS", "+62123456212"},
		{6, "Emily Wright", "9976777563", "Female", "Jl. Jakarta Barat No 15", "Kristen Protestan", "Jl. Jakarta Utara No 17", "+628910111225", "emilywright@gmail.com", "SMP 6", "Ayah 13", "PNS", "+62123456203", "Ibu 13", "PNS", "+62123456213"},
		{6, "Lucas Baker", "9976777564", "Male", "Jl. Jakarta Pusat No 16", "Kristen Protestan", "Jl. Jakarta Utara No 18", "+628910111226", "lucasbaker@gmail.com", "SMP 6", "Ayah 14", "PNS", "+62123456204", "Ibu 14", "PNS", "+62123456214"},
		{6, "Amelia Adams", "9976777565", "Female", "Jl. Jakarta Utara No 17", "Kristen Protestan", "Jl. Jakarta Utara No 19", "+628910111227", "ameliaadams@gmail.com", "SMP 6", "Ayah 15", "PNS", "+62123456205", "Ibu 15", "PNS", "+62123456215"},
		{6, "Henry Moore", "9976777566", "Male", "Jl. Jakarta Selatan No 18", "Kristen Protestan", "Jl. Jakarta Utara No 20", "+628910111228", "henrymoore@gmail.com", "SMP 6", "Ayah 16", "PNS", "+62123456206", "Ibu 16", "PNS", "+62123456216"},
		{6, "Sofia Thompson", "9976777567", "Female", "Jl. Jakarta Timur No 19", "Kristen Protestan", "Jl. Jakarta Utara No 21", "+628910111229", "sofiathompson@gmail.com", "SMP 6", "Ayah 17", "PNS", "+62123456207", "Ibu 17", "PNS", "+62123456217"},
		{6, "Elijah Hall", "9976777568", "Male", "Jl. Jakarta Barat No 20", "Kristen Protestan", "Jl. Jakarta Utara No 22", "+628910111230", "elijahhall@gmail.com", "SMP 6", "Ayah 18", "PNS", "+62123456208", "Ibu 18", "PNS", "+62123456218"},
		{6, "Mila Perez", "9976777569", "Female", "Jl. Jakarta Pusat No 21", "Kristen Protestan", "Jl. Jakarta Utara No 23", "+628910111231", "milaperez@gmail.com", "SMP 6", "Ayah 19", "PNS", "+62123456209", "Ibu 19", "PNS", "+62123456219"},
		{6, "Oliver Turner", "9976777570", "Male", "Jl. Jakarta Selatan No 22", "Kristen Protestan", "Jl. Jakarta Utara No 24", "+628910111232", "oliverturner@gmail.com", "SMP 6", "Ayah 20", "PNS", "+62123456210", "Ibu 20", "PNS", "+62123456220"},

		{7, "Grace Young", "9976777571", "Female", "Jl. Jakarta Selatan No 23", "Kristen Protestan", "Jl. Jakarta Utara No 25", "+628910111233", "graceyoung@gmail.com", "SMP 6", "Ayah 21", "PNS", "+62123456211", "Ibu 21", "PNS", "+62123456221"},
		{7, "Jack Evans", "9976777572", "Male", "Jl. Jakarta Timur No 24", "Kristen Protestan", "Jl. Jakarta Utara No 26", "+628910111234", "jackevans@gmail.com", "SMP 6", "Ayah 22", "PNS", "+62123456212", "Ibu 22", "PNS", "+62123456222"},
		{7, "Hannah Lewis", "9976777573", "Female", "Jl. Jakarta Barat No 25", "Kristen Protestan", "Jl. Jakarta Utara No 27", "+628910111235", "hannahlewis@gmail.com", "SMP 6", "Ayah 23", "PNS", "+62123456213", "Ibu 23", "PNS", "+62123456223"},
		{7, "William Collins", "9976777574", "Male", "Jl. Jakarta Pusat No 26", "Kristen Protestan", "Jl. Jakarta Utara No 28", "+628910111236", "williamcollins@gmail.com", "SMP 6", "Ayah 24", "PNS", "+62123456214", "Ibu 24", "PNS", "+62123456224"},
		{7, "Ella Martinez", "9976777575", "Female", "Jl. Jakarta Utara No 29", "Kristen Protestan", "Jl. Jakarta Utara No 30", "+628910111237", "ellamartinez@gmail.com", "SMP 6", "Ayah 25", "PNS", "+62123456215", "Ibu 25", "PNS", "+62123456225"},
		{7, "Michael Lopez", "9976777576", "Male", "Jl. Jakarta Selatan No 27", "Kristen Protestan", "Jl. Jakarta Utara No 31", "+628910111238", "michaellopez@gmail.com", "SMP 6", "Ayah 26", "PNS", "+62123456216", "Ibu 26", "PNS", "+62123456226"},
		{7, "Victoria Hill", "9976777577", "Female", "Jl. Jakarta Timur No 28", "Kristen Protestan", "Jl. Jakarta Utara No 32", "+628910111239", "victoriahill@gmail.com", "SMP 6", "Ayah 27", "PNS", "+62123456217", "Ibu 27", "PNS", "+62123456227"},
		{7, "Daniel Gonzalez", "9976777578", "Male", "Jl. Jakarta Barat No 29", "Kristen Protestan", "Jl. Jakarta Utara No 33", "+628910111240", "danielgonzalez@gmail.com", "SMP 6", "Ayah 28", "PNS", "+62123456218", "Ibu 28", "PNS", "+62123456228"},
		{7, "Ava Robinson", "9976777579", "Female", "Jl. Jakarta Pusat No 30", "Kristen Protestan", "Jl. Jakarta Utara No 34", "+628910111241", "avarobinson@gmail.com", "SMP 6", "Ayah 29", "PNS", "+62123456219", "Ibu 29", "PNS", "+62123456229"},
		{7, "David Martinez", "9976777580", "Male", "Jl. Jakarta Selatan No 31", "Kristen Protestan", "Jl. Jakarta Utara No 35", "+628910111242", "davidmartinez@gmail.com", "SMP 6", "Ayah 30", "PNS", "+62123456220", "Ibu 30", "PNS", "+62123456230"},
	}

	for _, data := range studentDataList {
		students = append(students, createStudent(data))
	}

	// additional student data for testing class archive
	additionalStudents := []models.Student{
		{ClassNameID: 8, StudentName: "John Doe", StudentNISN: "9988776655", StudentGender: "Male", StudentPlaceOfBirth: "Jakarta", StudentDateOfBirth: time.Date(2005, 1, 10, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Kemang Timur No 12", StudentPhoneNumber: "+628910111111", StudentEmail: "johndoe@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 2", StudentFatherName: "Ayah 31", StudentFatherJob: "Wiraswasta", StudentFatherPhoneNumber: "+62123456111", StudentMotherName: "Ibu 31", StudentMotherJob: "Guru", StudentMotherPhoneNumber: "+62123456112"},
		{ClassNameID: 8, StudentName: "Jane Smith", StudentNISN: "9966443321", StudentGender: "Female", StudentPlaceOfBirth: "Bandung", StudentDateOfBirth: time.Date(2006, 5, 15, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Thamrin No 20", StudentPhoneNumber: "+628910222222", StudentEmail: "janesmith@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 3", StudentFatherName: "Ayah 32", StudentFatherJob: "Dokter", StudentFatherPhoneNumber: "+62123456222", StudentMotherName: "Ibu 32", StudentMotherJob: "Perawat", StudentMotherPhoneNumber: "+62123456223"},
		{ClassNameID: 8, StudentName: "Michael Johnson", StudentNISN: "9955332211", StudentGender: "Male", StudentPlaceOfBirth: "Surabaya", StudentDateOfBirth: time.Date(2005, 12, 20, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Senayan No 25", StudentPhoneNumber: "+628910333333", StudentEmail: "michaeljohnson@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 4", StudentFatherName: "Ayah 33", StudentFatherJob: "Karyawan Swasta", StudentFatherPhoneNumber: "+62123456333", StudentMotherName: "Ibu 33", StudentMotherJob: "Ibu Rumah Tangga", StudentMotherPhoneNumber: "+62123456334"},
		{ClassNameID: 8, StudentName: "Emily Davis", StudentNISN: "9944221100", StudentGender: "Female", StudentPlaceOfBirth: "Medan", StudentDateOfBirth: time.Date(2006, 7, 8, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Cilandak No 7", StudentPhoneNumber: "+628910444444", StudentEmail: "emilydavis@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 7", StudentFatherName: "Ayah 34", StudentFatherJob: "Pengusaha", StudentFatherPhoneNumber: "+62123456444", StudentMotherName: "Ibu 34", StudentMotherJob: "PNS", StudentMotherPhoneNumber: "+62123456445"},
		{ClassNameID: 8, StudentName: "Chris Brown", StudentNISN: "9933110099", StudentGender: "Male", StudentPlaceOfBirth: "Yogyakarta", StudentDateOfBirth: time.Date(2005, 2, 11, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Tebet Barat No 8", StudentPhoneNumber: "+628910555555", StudentEmail: "chrisbrown@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 5", StudentFatherName: "Ayah 35", StudentFatherJob: "Polisi", StudentFatherPhoneNumber: "+62123456555", StudentMotherName: "Ibu 35", StudentMotherJob: "Guru", StudentMotherPhoneNumber: "+62123456556"},
		{ClassNameID: 8, StudentName: "Sophia Wilson", StudentNISN: "9922008899", StudentGender: "Female", StudentPlaceOfBirth: "Malang", StudentDateOfBirth: time.Date(2006, 9, 5, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Kelapa Gading No 16", StudentPhoneNumber: "+628910666666", StudentEmail: "sophiawilson@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 8", StudentFatherName: "Ayah 36", StudentFatherJob: "TNI", StudentFatherPhoneNumber: "+62123456666", StudentMotherName: "Ibu 36", StudentMotherJob: "Dokter", StudentMotherPhoneNumber: "+62123456667"},
		{ClassNameID: 8, StudentName: "Liam Garcia", StudentNISN: "9911007788", StudentGender: "Male", StudentPlaceOfBirth: "Palembang", StudentDateOfBirth: time.Date(2005, 4, 1, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Menteng No 11", StudentPhoneNumber: "+628910777777", StudentEmail: "liamgarcia@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 9", StudentFatherName: "Ayah 37", StudentFatherJob: "Wiraswasta", StudentFatherPhoneNumber: "+62123456777", StudentMotherName: "Ibu 37", StudentMotherJob: "Perawat", StudentMotherPhoneNumber: "+62123456778"},
		{ClassNameID: 8, StudentName: "Isabella Martinos", StudentNISN: "9900099888", StudentGender: "Female", StudentPlaceOfBirth: "Makassar", StudentDateOfBirth: time.Date(2006, 6, 25, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Fatmawati No 6", StudentPhoneNumber: "+628910888888", StudentEmail: "isabelladiatercinta@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 10", StudentFatherName: "Ayah 38", StudentFatherJob: "Pengusaha", StudentFatherPhoneNumber: "+62123456888", StudentMotherName: "Ibu 38", StudentMotherJob: "Ibu Rumah Tangga", StudentMotherPhoneNumber: "+62123456889"},
		{ClassNameID: 8, StudentName: "Noah Thomas", StudentNISN: "9898987766", StudentGender: "Male", StudentPlaceOfBirth: "Bali", StudentDateOfBirth: time.Date(2005, 8, 12, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Pancoran No 2", StudentPhoneNumber: "+628910999999", StudentEmail: "noahthomas@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 11", StudentFatherName: "Ayah 39", StudentFatherJob: "Karyawan Swasta", StudentFatherPhoneNumber: "+62123456999", StudentMotherName: "Ibu 39", StudentMotherJob: "PNS", StudentMotherPhoneNumber: "+62123456990"},
		{ClassNameID: 8, StudentName: "Ava White", StudentNISN: "9888876655", StudentGender: "Female", StudentPlaceOfBirth: "Pontianak", StudentDateOfBirth: time.Date(2006, 11, 3, 0, 0, 0, 0, time.UTC), StudentReligion: "Konghucu", StudentAddress: "Jl. Setiabudi No 10", StudentPhoneNumber: "+628910000000", StudentEmail: "avawhite@gmail.com", StudentAcceptedDate: time.Date(2021, 7, 20, 0, 0, 0, 0, time.UTC), StudentSchoolOfOrigin: "SMP 12", StudentFatherName: "Ayah 40", StudentFatherJob: "Polisi", StudentFatherPhoneNumber: "+62123457000", StudentMotherName: "Ibu 40", StudentMotherJob: "Guru", StudentMotherPhoneNumber: "+62123457001"},
	}

	// for _, data := range additionalStudents {
	// 	// add deleted at for testing class archive
	// 	data.DeletedAt = gorm.DeletedAt{Time: time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC)}
	// }
	for i := range additionalStudents {
		// Add DeletedAt only for the additional students
		additionalStudents[i].StudentName = "Deleted " + additionalStudents[i].StudentName
		additionalStudents[i].DeletedAt = gorm.DeletedAt{
			Time: time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC), 
			Valid: true,
		}
	}

	students = append(students, additionalStudents...)

	return students
}
