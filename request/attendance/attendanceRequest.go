// requests/create_teacher_request.go
package request

import (
	"fmt"

	"github.com/edulink-api/models"
	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* AttendanceRequest struct
*
*
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type AttendanceRequest struct {
	models.UpdateClassDateAttendanceStudent
}

// Validate method
func (r *AttendanceRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}

/*
* AllAttendanceRequest struct
*
*
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type AllAttendanceRequest struct {
	AttendanceRequest []AttendanceRequest `json:"student" binding:"required"`
}

// Validate method
func (r *AllAttendanceRequest) ValidateAllAttendance() []map[string]string {
	// Validate the struct
	var allErrors []map[string]string
	for i, data := range r.AttendanceRequest {
		if err := data.Validate(); err != nil {
			// index with error
			errorMap := map[string]string{
				"row-error": fmt.Sprintf("%d", i+1),
				"field":     err[0]["field"],
				"message":   err[0]["message"],
			}
			allErrors = append(allErrors, errorMap)
		}
	}

	return allErrors
}
