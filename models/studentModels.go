package models

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

type Student struct {
	StudentID             int64     `gorm:"primaryKey"`
	ClassID               int64     `json:"id_class" binding:"required"`
	StudentName           string    `json:"name" binding:"required"`
	StudentGender         string    `json:"gender" binding:"required,oneof=Male Female other"`
	StudentPlaceOfBirth   string    `json:"place_of_birth" binding:"required"`
	StudentDateOfBirth    time.Time `json:"date_of_birth"`
	StudentReligion       string    `json:"religion" binding:"required"`
	StudentAddress        string    `json:"address" binding:"required"`
	StudentNumPhone       string    `json:"number_phone" binding:"required,e164"`
	StudentEmail          string    `json:"email" binding:"required,email"`
	StudentAcceptedDate   time.Time `json:"accepted_date"`
	StudentSchoolOrigin   string    `json:"school_origin" binding:"required"`
	StudentFatherName     string    `json:"father_name" binding:"required"`
	StudentFatherJob      string    `json:"father_job" binding:"required"`
	StudentFatherNumPhone string    `json:"father_number_phone" binding:"required,e164"`
	StudentMotherName     string    `json:"mother_name" binding:"required"`
	StudentMotherJob      string    `json:"mother_job" binding:"required"`
	StudentMotherNumPhone string    `json:"mother_number_phone" binding:"required,e164"`
	lib.BaseModel
}

func (Student) TableName() string {
	return lib.GenerateTableName(lib.Academic, "students")
}
