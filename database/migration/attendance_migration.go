package migration

import (
	"time"
)

type Attendance struct {
	AttendanceID int64 `gorm:"primaryKey;autoIncrement"`
	StudentID    int64 `gorm:"not null"`
	AttendanceDate time.Time `gorm:"not null"`
	AttendanceStatus string `gorm:"not null"`
}
