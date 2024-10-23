// requests/create_teacher_request.go
package request

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/skripsi-be/models"
)

/*
* InsertTeacherRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertTeacherRequest struct {
	models.User
	TeachingHour int32       `json:"teaching_hour" binding:"required"`
	DateOfBirth  string `json:"date_of_birth" binding:"required"`
}

// Validate method
func (r *InsertTeacherRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

/*
* see the documentation about time.Parse here:
* https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
 */
func (r *InsertTeacherRequest) ParseDates() (time.Time, error) {
	DateOfBirth, err := time.Parse("2006-01-02", r.DateOfBirth)
	if err != nil {
		return time.Time{}, err
	}

	return DateOfBirth, nil
}

/*
* InsertAllTeacherRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertAllTeacherRequest struct {
	InsertTeacherRequest []InsertTeacherRequest `json:"teacher-data" binding:"required"`
}

// Validate method
func (r *InsertAllTeacherRequest) ValidateAllTeacher() error {
	validate := validator.New()
	return validate.Struct(r)
}
