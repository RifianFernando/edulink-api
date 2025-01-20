package models

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type EventSchedule struct {
	ScheduleID        int64     `json:"schedule_id" binding:"required" validate:"required"`
	EventScheduleName string    `json:"event_name" binding:"required" validate:"required"`
	EventScheduleDate time.Time `json:"schedule_date"`
	lib.BaseModel
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (EventSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "event_schedules")
}
