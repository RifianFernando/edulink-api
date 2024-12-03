// requests/create_teacher_request.go
package request

import (
	"time"

	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertTeacherRequest struct
*
*
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertTeacherRequest struct {
	UserName         string `json:"name" binding:"required"`
	UserGender       string `json:"gender" binding:"required" validate:"oneof='Male' 'Female'"`
	UserPlaceOfBirth string `json:"place_of_birth" binding:"required"`
	UserDateOfBirth  time.Time
	UserReligion     string  `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	UserAddress      string  `json:"address" binding:"required"`
	UserPhoneNum     string  `json:"num_phone" binding:"required" validate:"required,e164"`
	UserEmail        string  `json:"email" binding:"required" validate:"required,email"`
	DateOfBirth      string  `json:"date_of_birth" binding:"required" validate:"required,datetime=2006-01-02"`
	TeachingHour     string  `json:"teaching_hour" binding:"required" validate:"required"`
	TeachingSubject  []int64 `json:"teaching_subject" binding:"required" validate:"required,min=1,dive,min=1"`
}

// Validate method
func (r *InsertTeacherRequest) ValidateTeacher() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
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
*
*
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertAllTeacherRequest struct {
	InsertTeacherRequest []InsertTeacherRequest `json:"teacher-data" binding:"required"`
}

// Validate method
func (r *InsertAllTeacherRequest) ValidateAllTeacher() []map[string]string {
	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}
