package request

import (
	"github.com/go-playground/validator/v10"
)

type InsertForgetPasswordRequest struct {
	UserEmail    string `json:"email" binding:"required" validate:"email"`
}

// Validate method
func (r *InsertForgetPasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
