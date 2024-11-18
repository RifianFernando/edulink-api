package models

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

type Attendance struct {
	AttendanceID     int64     `gorm:"primaryKey;autoIncrement"`
	StudentID        int64     `gorm:"not null"`
	AttendanceDate   time.Time `gorm:"not null"`
	AttendanceStatus string    `gorm:"not null" validate:"oneof='Absent' 'Leave' 'Sick' 'Present'"`
}

type AttendanceModel struct {
	Attendance
	Student Student `gorm:"foreignKey:StudentID;references:StudentID"`
}

func (Attendance) TableName() string {
	return lib.GenerateTableName(lib.Administration, "attendances")
}
