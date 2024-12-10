package request

import (
	"fmt"

	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertStudentScoreRequest struct {
	StudentID int64 `json:"studentID" binding:"required"`
	Score     int   `json:"score" binding:"required"`
}

// Validate method
func (r *InsertStudentScoreRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}

/*
* InsertAllStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertAllStudentScoreRequest struct {
	AssignmentID         int64                       `json:"assignment_id" binding:"required"`
	InsertStudentRequest []InsertStudentScoreRequest `json:"scores" binding:"required"`
}

// Validate method
func (r *InsertAllStudentScoreRequest) ValidateAllStudentScore() []map[string]string {
	// Validate the struct
	var allErrors []map[string]string
	err := res.ResponseMessage(req.Validate.Struct(r))
	if err != nil {
		return err
	}

	for i, data := range r.InsertStudentRequest {
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
