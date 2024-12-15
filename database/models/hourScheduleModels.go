package models

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type HourSchedule struct {
	HourScheduleID int64 `gorm:"primaryKey;autoIncrement"`
	StartHour      int   `json:"start_hour" binding:"required" validate:"required"`
	EndHour        int   `json:"end_hour" binding:"required" validate:"required"`
	Schedule       []Schedule
	lib.BaseModel
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (HourSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "hour_schedules")
}
