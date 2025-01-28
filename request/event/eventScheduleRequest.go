package request

import (
	"time"

	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertEventScheduleRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertEventScheduleRequest struct {
	EventDate string `json:"event_date" binding:"required" validate:"required,datetime=2006-01-02"`
	EventName string `json:"event_name" binding:"required" validate:"required"`
	StartTIme string `json:"start_time" binding:"required" validate:"required"`
	EndTime   string `json:"end_time" binding:"required" validate:"required"`
}

// Validate method
func (r *InsertEventScheduleRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}

/*
* see the documentation about time.Parse here:
* https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
 */
func (r *InsertEventScheduleRequest) ParseDates() (time.Time, error) {
	EventDate, err := time.Parse("2006-01-02", r.EventDate)
	if err != nil {
		return time.Time{}, err
	}

	return EventDate, nil
}
