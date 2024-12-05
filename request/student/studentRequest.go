// requests/create_student_request.go
package request

import (
	"fmt"
	"time"

	"github.com/edulink-api/database/models"
	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertStudentRequest struct {
	models.Student
	DateOfBirth  string `json:"date_of_birth" binding:"required" validate:"required,datetime=2006-01-02"`
	AcceptedDate string `json:"accepted_date" binding:"required" validate:"required,datetime=2006-01-02"`
}

// Validate method
func (r *InsertStudentRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}

/*
* see the documentation about time.Parse here:
* https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
 */
func (r *InsertStudentRequest) ParseDates() (time.Time, time.Time, error) {
	DateOfBirth, err := time.Parse("2006-01-02", r.DateOfBirth)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	AcceptedDate, err := time.Parse("2006-01-02", r.AcceptedDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return DateOfBirth, AcceptedDate, nil
}

/*
* InsertAllStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertAllStudentRequest struct {
	InsertStudentRequest []InsertStudentRequest `json:"student-data" binding:"required"`
}

// Validate method
// func (r *InsertAllStudentRequest) ValidateAllStudent() error {
// 	for _, data := range r.InsertStudentRequest {
// 		if err := data.Validate(); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// Validate method
func (r *InsertAllStudentRequest) ValidateAllStudent() []map[string]string {
	// Validate the struct
	var allErrors []map[string]string
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
