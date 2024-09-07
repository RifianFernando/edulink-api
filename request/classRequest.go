package request

import (
	"github.com/go-playground/validator/v10"
)

type InsertClassRequest struct {
	IDTeacher uint   `json:"id_teacher" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Grade     int `json:"grade" binding:"required"`
}

// Validate method
func (r *InsertClassRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

