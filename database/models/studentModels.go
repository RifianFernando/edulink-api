package models

import (
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type Student struct {
	StudentID                int64     `gorm:"primaryKey"`
	ClassNameID              int64     `json:"class_id" binding:"required"`
	StudentName              string    `json:"name" binding:"required"`
	StudentNISN              string    `json:"nisn" binding:"required" validate:"len=10"`
	StudentGender            string    `json:"gender" binding:"required,oneof='Male' 'Female'"`
	StudentPlaceOfBirth      string    `json:"place_of_birth" binding:"required"`
	StudentDateOfBirth       time.Time `json:"date_of_birth"`
	StudentReligion          string    `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	StudentAddress           string    `json:"address" binding:"required" validate:"required,min=1,max=200"`
	StudentPhoneNumber       string    `json:"number_phone" binding:"required" validate:"required,e164"`
	StudentEmail             string    `json:"email" binding:"required" validate:"required,email"`
	StudentAcceptedDate      time.Time `json:"accepted_date"`
	StudentSchoolOfOrigin    string    `json:"school_origin" binding:"required"`
	StudentFatherName        string    `json:"father_name" binding:"required"`
	StudentFatherJob         string    `json:"father_job" binding:"required"`
	StudentFatherPhoneNumber string    `json:"father_number_phone" binding:"required,e164"`
	StudentMotherName        string    `json:"mother_name" binding:"required"`
	StudentMotherJob         string    `json:"mother_job" binding:"required"`
	StudentMotherPhoneNumber string    `json:"mother_number_phone" binding:"required,e164"`
	lib.BaseModel
}

type StudentModel struct {
	Student
	ClassName ClassNameGrade `gorm:"foreignKey:ClassNameID;references:ClassNameID"`
}

type StudentAttendance struct {
	Student
	Attendance []Attendance `gorm:"foreignKey:StudentID;references:StudentID"`
}

func (Student) TableName() string {
	return lib.GenerateTableName(lib.Academic, "students")
}

func (StudentModel) TableName() string {
	return lib.GenerateTableName(lib.Academic, "students")
}

// Create student
func (student *Student) CreateStudent() error {
	result := connections.DB.Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all students
func (student *StudentModel) GetAllStudents() (
	students []StudentModel,
	msg string,
) {
	result := connections.DB.Preload("ClassName.Grade").Find(&students)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No students found"
	}

	return students, ""
}

func (student *Student) GetAllStudentByClassNameID() (
	students []Student,
	msg string,
) {
	result := connections.DB.
		Where("class_name_id = ?", student.ClassNameID).
		Find(&students)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No students found"
	}

	return students, ""
}

func (student *Student) GetStudent() (Student, error) {
	// get result by model
	result := connections.DB.Where(&student).First(&student)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		if result.Error == gorm.ErrRecordNotFound {
			return Student{}, nil
		}
		return Student{}, result.Error
	}

	return *student, nil
}

// get student by id
func (student *StudentModel) GetStudentById(id string) (StudentModel, error) {
	var students StudentModel
	result := connections.DB.Preload("ClassName.Grade").First(&students, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return students, nil
		}
		return students, result.Error
	}

	return students, nil
}

// update student by id
func (student *Student) UpdateStudentById(students *Student) error {
	result := connections.DB.Model(&student).Updates(&students)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// delete student by id
func (student *Student) DeleteStudentById(id string) error {
	result := connections.DB.Delete(&student, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CreateAllStudents(students []Student) error {
	result := connections.DB.Create(&students)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type UpdateManyStudentClass struct {
	StudentID   int64 `json:"student_id" binding:"required" validate:"required,numeric,gte=1"`
	ClassNameID int64 `json:"class_name_id" binding:"required" validate:"required,numeric,gte=1"`
}

func UpdateManyStudentClassID(studentData []UpdateManyStudentClass) error {
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, data := range studentData {
		result := tx.Model(&Student{
			StudentID: data.StudentID,
		}).Updates(Student{
			ClassNameID: data.ClassNameID,
		})
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	return tx.Commit().Error
}

func (student *Student) GetAllStudentsByClassID(classID string) (students []Student, err error) {
	result := connections.DB.Where("class_name_id = ?", classID).Find(&students)
	if result.Error != nil {
		return students, result.Error
	}

	return students, nil
}
