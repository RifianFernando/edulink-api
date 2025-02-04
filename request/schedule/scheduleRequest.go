package request

import (
	"fmt"

	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertScheduleRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type ScheduleRequest struct {
	TeacherID    int64 `json:"teacher_id" binding:"required" validate:"required"`
	TeachingHour int32 `json:"teaching_hour" binding:"required" validate:"required"`
	DataTeaching []struct {
		SubjectID   int64   `json:"subject_id" binding:"required" validate:"required,min=1"`
		ClassNameID []int64 `json:"class_name_id" binding:"required" validate:"required,min=1,dive"`
	} `json:"data_teaching" binding:"required" validate:"required,min=1,dive"`
}

// Validate method
func (r *ScheduleRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}

type InsertScheduleRequest struct {
	ScheduleRequest []ScheduleRequest `json:"schedule" binding:"required"`
}

// Validate method
func (r *InsertScheduleRequest) Validate() []map[string]string {
	// Validate the struct
	var allErrors []map[string]string
	for i, data := range r.ScheduleRequest {
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
