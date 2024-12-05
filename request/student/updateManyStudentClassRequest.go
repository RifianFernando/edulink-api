package request

import (
	"github.com/edulink-api/database/models"
	"github.com/go-playground/validator/v10"
)

/*
* InsertStudentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type UpdateManyStudentClassRequest struct {
	UpdateStudentClass []models.UpdateManyStudentClass `json:"student-data" binding:"required"`
}

// Validate method
func (r *UpdateManyStudentClassRequest) ValidateAllData() error {
	for _, data := range r.UpdateStudentClass {
		if err := validator.New().Struct(data); err != nil {
			return err
		}
	}
	return nil
}
