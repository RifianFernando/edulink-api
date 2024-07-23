package migration

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
 */
type EventSchedule struct {
	EventScheduleID   int64     `gorm:"primaryKey;autoIncrement"`
	ScheduleID        int64     `gorm:"not null"`
	EventScheduleName string    `gorm:"not null"`
	EventScheduleDate time.Time `gorm:"not null"`
	// Schedule          Schedule  `gorm:"foreignKey:ScheduleID;references:ScheduleID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Schedule      Schedule
	lib.BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (EventSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "event_schedules")
}
