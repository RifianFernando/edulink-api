package models

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type DaySchedule struct {
	DayScheduleID int64  `gorm:"primaryKey;autoIncrement"`
	DayName       string `json:"day_name" binding:"required" validate:"required,max=10"`
	Schedule      []Schedule
	lib.BaseModel
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (DaySchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "day_schedules")
}
