package migration

import (
	"time"
)

type Subject struct {
	SubjectID       int64     `gorm:"primaryKey;autoIncrement"`
	SubjectName     string    `gorm:"unique;not null"`
	SubjectDuration time.Time `gorm:"not null"`
}
