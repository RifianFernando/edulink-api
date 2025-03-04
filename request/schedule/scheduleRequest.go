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
	ScheduleRequest []ScheduleRequest `json:"schedule" binding:"required" validate:"required,min=1,dive"`
}

// Validate method
func (r *InsertScheduleRequest) Validate() []map[string]string {
	// Validate the struct
	var allErrors []map[string]string

	// Create a map to track unique TeacherIDs
	teacherIDMap := make(map[int64]bool)

	err := res.ResponseMessage(req.Validate.Struct(r))
	if len(err) > 0 {
		return err
	}
	for i, data := range r.ScheduleRequest {
		// Check if TeacherID is unique
		if _, exists := teacherIDMap[data.TeacherID]; exists {
			allErrors = append(allErrors, map[string]string{
				"row-error":  fmt.Sprintf("%d", i + 1),
				"field":      "teacher_id",
				"message":    "TeacherID must be unique across all schedule requests",
			})
		} else {
			teacherIDMap[data.TeacherID] = true
		}
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
