// requests/create_student_request.go
package request

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type InsertStudentRequest struct {
	Name              string `json:"name" binding:"required"`
	Gender            string `json:"gender" binding:"required,oneof=male female other"`
	PlaceOfBirth      string `json:"place_of_birth" binding:"required"`
	DateOfBirth       string `json:"date_of_birth" binding:"required"`
	Religion          string `json:"religion" binding:"required"`
	Address           string `json:"address" binding:"required"`
	NumberPhone       string `json:"number_phone" binding:"required,e164"`
	Email             string `json:"email" binding:"required,email"`
	AcceptedDate      string `json:"accepted_date" binding:"required"`
	SchoolOrigin      string `json:"school_origin" binding:"required"`
	IDClass           uint   `json:"id_class" binding:"required"`
	FatherName        string `json:"father_name" binding:"required"`
	FatherJob         string `json:"father_job" binding:"required"`
	FatherNumberPhone string `json:"father_number_phone" binding:"required,e164"`
	MotherName        string `json:"mother_name" binding:"required"`
	MotherJob         string `json:"mother_job" binding:"required"`
	MotherNumberPhone string `json:"mother_number_phone" binding:"required,e164"`
}

// Validate method
func (r *InsertStudentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Method to parse dates
func (r *InsertStudentRequest) ParseDates() (time.Time, time.Time, error) {
	dateOfBirth, err := time.Parse("2006-01-02", r.DateOfBirth)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	acceptedDate, err := time.Parse("2006-01-02", r.AcceptedDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return dateOfBirth, acceptedDate, nil
}
