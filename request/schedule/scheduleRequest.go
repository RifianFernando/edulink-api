package request

import (
	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertScheduleRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertScheduleRequest struct {
}

// Validate method
func (r *InsertScheduleRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}
