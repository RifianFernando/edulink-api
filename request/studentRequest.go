// requests/create_student_request.go
package request

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/skripsi-be/models"
)

/*
* InsertStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertStudentRequest struct {
	models.Student
	DateOfBirth  string `json:"date_of_birth" binding:"required"`
	AcceptedDate string `json:"accepted_date" binding:"required"`
}

// Validate method
func (r *InsertStudentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
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
func (r *InsertAllStudentRequest) ValidateAllStudent() error {
	validate := validator.New()
	return validate.Struct(r)
}
