package migration

import (
	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Schedule struct {
	ScheduleID       int64              `gorm:"primaryKey;autoIncrement"`
	DayScheduleID    int64              `gorm:"not null"`
	HourScheduleID   int64              `gorm:"not null"`
	LearningSchedule []LearningSchedule `gorm:"foreignKey:ScheduleID;references:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:set null"`
	EventSchedule    []EventSchedule    `gorm:"foreignKey:ScheduleID;references:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:set null"`
	lib.BaseModel                       /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Schedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "schedules")
}
