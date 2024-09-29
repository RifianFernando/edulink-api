package migration

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Student struct {
	StudentID             int64     `gorm:"primaryKey;autoIncrement"`
	ClassID               int64     `gorm:"not null"`
	StudentName           string    `gorm:"unique;not null"`
	StudentNISN           string    `gorm:"unique;not null;size:10"`
	StudentGender         Gender    `gorm:"type:VARCHAR(6);not null;"`
	StudentPlaceOfBirth   string    `gorm:"not null"`
	StudentDateOfBirth    time.Time `gorm:"not null"`
	StudentReligion       string    `gorm:"not null"`
	StudentAddress        string    `gorm:"not null"`
	StudentNumPhone       string    `gorm:"unique;not null"`
	StudentEmail          string    `gorm:"unique;not null"`
	StudentAcceptedDate   time.Time `gorm:"not null"`
	StudentSchoolOrigin   string    `gorm:"not null"`
	StudentFatherName     string    `gorm:"not null"`
	StudentFatherJob      string    `gorm:"not null"`
	StudentFatherNumPhone string    `gorm:"not null"`
	StudentMotherName     string    `gorm:"not null"`
	StudentMotherJob      string    `gorm:"not null"`
	StudentMotherNumPhone string
	Grades                []Grade             `gorm:"foreignKey:StudentID;references:StudentID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Attendance            []Attendance        `gorm:"foreignKey:StudentID;references:StudentID"`
	AttendanceSummaries   []AttendanceSummary `gorm:"foreignKey:StudentID;references:StudentID"`
	lib.BaseModel                             /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Student) TableName() string {
	return lib.GenerateTableName(lib.Academic, "students")
}
