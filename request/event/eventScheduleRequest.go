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
	EventName      string    `json:"event_name" binding:"required" validate:"required,min=3"`
	EventDateStart time.Time `json:"event_date_start" binding:"required" validate:"required"`
	EventDateEnd   time.Time `json:"event_date_end" binding:"required" validate:"required"`
}

// Validate method
func (r *InsertEventScheduleRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))
	if (r.EventDateStart).After(r.EventDateEnd) {
		err = append(err, map[string]string{
			"error": "Event date start must be before event date end",
		})
	}
	if (r.EventDateStart).Equal(r.EventDateEnd) {
		err = append(err, map[string]string{
			"error": "Event date start must be different with event date end",
		})
	}
	if (r.EventDateStart).Before(time.Now()) {
		err = append(err, map[string]string{
			"error": "Event date start must be after today",
		})
	}
	if (r.EventDateEnd).Before(time.Now()) {
		err = append(err, map[string]string{
			"error": "Event date end must be after today",
		})
	}

	return err
}

/*
* see the documentation about time.Parse here:
* https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
 */
// func (r *InsertEventScheduleRequest) ParseDates() (time.Time, time.Time, error) {
// 	EventDateStart, err := time.Parse("2006-01-02", r.EventDateStart)
// 	if err != nil {
// 		return time.Time{}, time.Time{}, fmt.Errorf("invalid date format ma bro")
// 	}

// 	EventDateEnd, err := time.Parse("2006-01-02", r.EventDateEnd)
// 	if err != nil {
// 		return time.Time{}, time.Time{}, fmt.Errorf("invalid date format ma bro")
// 	}

// 	return EventDateStart, EventDateEnd, nil
// }
