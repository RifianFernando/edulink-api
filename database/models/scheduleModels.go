package models

import (
	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Schedule struct {
	ScheduleID       int64 `gorm:"primaryKey;autoIncrement"`
	DayScheduleID    int64 `json:"day_schedule_id" binding:"required" validate:"required"`
	HourScheduleID   int64 `json:"hour_schedule_id" binding:"required" validate:"required"`
	AcademicYearID   int64 `json:"academic_year_id" binding:"required" validate:"required"`
	LearningSchedule []LearningSchedule
	lib.BaseModel
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Schedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "schedules")
}
