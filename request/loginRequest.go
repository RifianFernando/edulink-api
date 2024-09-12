package request

import (
	"github.com/go-playground/validator/v10"
)

type InsertLoginRequest struct {
	UserEmail    string `json:"user_email" binding:"required" validate:"email"`
	UserPassword string `json:"user_password" binding:"required"`
}

// Validate method
func (r *InsertLoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
