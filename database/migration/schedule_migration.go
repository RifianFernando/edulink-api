package migration

import (
	"time"
)

type Schedule struct {
	ScheduleID         int64     `gorm:"primaryKey;autoIncrement"`
	ScheduleSchoolYear int32     `gorm:"not null"`
	ScheduleStartTime  time.Time `gorm:"not null"`
	ScheduleEndTime    time.Time `gorm:"not null"`
}
