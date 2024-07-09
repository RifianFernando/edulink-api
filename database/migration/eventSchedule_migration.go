package migration

import (
	"time"
)

type EventSchedule struct {
	EventScheduleID   int64     `gorm:"primaryKey;autoIncrement"`
	ScheduleID        int64     `gorm:"not null"`
	EventScheduleName string    `gorm:"not null"`
	EventScheduleDate time.Time `gorm:"not null"`
	Schedule          Schedule  `gorm:"foreignKey:ScheduleID;references:ScheduleID"`
}
