package models

import (
	"fmt"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type EventSchedule struct {
	EventScheduleID        int64     `gorm:"primaryKey"`
	EventScheduleName      string    `json:"event_name" binding:"required" validate:"required"`
	EventScheduleDateStart time.Time `json:"schedule_date_start" binding:"required" validate:"required,datetime=2006-01-02"`
	EventScheduleDateEnd   time.Time `json:"schedule_date_end" binding:"required" validate:"required,datetime=2006-01-02"`
	lib.BaseModel
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (EventSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "event_schedules")
}

func (e *EventSchedule) CreateEventSchedule() error {
	result := connections.DB.Create(&e)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (e *EventSchedule) UpdateEventSchedule() error {
	result := connections.DB.Save(&e)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("eventSchedule not found")
	}

	return nil
}

func (e *EventSchedule) DeleteEventSchedule() error {
	result := connections.DB.Delete(&e)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("eventSchedule not found")
	}

	return nil
}

func (e *EventSchedule) GetAllEvent() (event []EventSchedule, err error) {
	result := connections.DB.Find(&event)
	if result.Error != nil {
		return []EventSchedule{}, result.Error
	} else if result.RowsAffected == 0 {
		return []EventSchedule{}, fmt.Errorf("eventSchedule not found")
	}

	return event, nil
}
