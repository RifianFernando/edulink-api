package migration

import (
	"time"
)

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type Schedule struct {
	ScheduleID         int64     `gorm:"primaryKey;autoIncrement"`
	ScheduleSchoolYear int32     `gorm:"not null"`
	ScheduleStartTime  time.Time `gorm:"not null"`
	ScheduleEndTime    time.Time `gorm:"not null"`
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (Schedule) TableName() string {
	return GenerateTableName(Administration, "schedules")
}
