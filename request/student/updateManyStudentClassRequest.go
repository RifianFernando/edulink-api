package request

import (
	"github.com/go-playground/validator/v10"
)

/*
* InsertStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type UpdateManyStudentClassRequest struct {
	StudentID int64 `json:"student_id" binding:"required" validate:"required,numeric"`
	ClassID   int64 `json:"class_id" binding:"required" validate:"required,numeric"`
}

// Validate method
func (r *UpdateManyStudentClassRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
