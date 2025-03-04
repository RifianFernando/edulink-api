package migration

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
	EventScheduleID        int64     `gorm:"primaryKey;autoIncrement"`
	EventScheduleName      string    `gorm:"not null"`
	EventScheduleDateStart time.Time `gorm:"not null"`
	EventScheduleDateEnd   time.Time `gorm:"not null"`
	lib.BaseModel                    /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (EventSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "event_schedules")
}
