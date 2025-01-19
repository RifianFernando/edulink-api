package migration

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Student struct {
	StudentID                int64        `gorm:"primaryKey;autoIncrement"`
	ClassNameID              *int64       // nullable foreign key
	StudentName              string       `gorm:"not null"`
	StudentNISN              string       `gorm:"not null;unique;type:CHAR(10)"`
	StudentGender            Gender       `gorm:"type:VARCHAR(6);not null"`
	StudentPlaceOfBirth      string       `gorm:"not null;type:VARCHAR(35)"`
	StudentDateOfBirth       time.Time    `gorm:"not null"`
	StudentReligion          string       `gorm:"not null;type:VARCHAR(17)"`
	StudentAddress           string       `gorm:"not null;type:VARCHAR(200)"`
	StudentPhoneNumber       string       `gorm:"unique;not null;type:VARCHAR(20)"`
	StudentEmail             string       `gorm:"unique;not null;type:VARCHAR(40)"`
	StudentAcceptedDate      time.Time    `gorm:"not null"`
	StudentSchoolOfOrigin    string       `gorm:"not null"`
	StudentFatherName        string       `gorm:"not null"`
	StudentFatherJob         string       `gorm:"not null"`
	StudentFatherPhoneNumber string       `gorm:"not null"`
	StudentMotherName        string       `gorm:"not null"`
	StudentMotherJob         string       `gorm:"not null"`
	StudentMotherPhoneNumber string       `gorm:"not null"`
	Scores                   []Score      `gorm:"foreignKey:StudentID;references:StudentID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Attendance               []Attendance `gorm:"foreignKey:StudentID;references:StudentID"`
	lib.BaseModel                         /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Student) TableName() string {
	return lib.GenerateTableName(lib.Academic, "students")
}
