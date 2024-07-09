package migration

import (
	"time"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Student struct {
	StudentID            int64  `gorm:"primaryKey;autoIncrement"`
	StudentGender        Gender `gorm:"type:gender;not null"`
	StudentName          string `gorm:"unique;not null"`
	StudentPlaceOfBirth  string `gorm:"not null"`
	StudentDateOfBirth   time.Time `gorm:"not null"`
	StudentReligion      string `gorm:"not null"`
	StudentAddress       string `gorm:"not null"`
	StudentNumPhone      string `gorm:"unique;not null"`
	StudentEmail         string `gorm:"unique;not null"`
	StudentAcceptedDate  time.Time `gorm:"not null"`
	StudentSchoolOrigin  string `gorm:"unique;not null"`
	StudentFatherName    string `gorm:"not null"`
	StudentFatherJob     string `gorm:"not null"`
	StudentFatherNumPhone string `gorm:"not null"`
	StudentMotherName    string `gorm:"not null"`
	StudentMotherJob     string `gorm:"not null"`
	StudentMotherNumPhone string
	Grades               []Grade `gorm:"foreignKey:StudentID;references:StudentID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Attendance           []Attendance `gorm:"foreignKey:StudentID;references:StudentID"`
	AttendanceSummaries  []AttendanceSummary `gorm:"foreignKey:StudentID;references:StudentID"`
}
