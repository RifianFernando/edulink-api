package request

import (
	"github.com/edulink-api/database/models"
	"github.com/go-playground/validator/v10"
)

type InsertClassRequest struct {
	models.ClassName
}

// Validate method
func (r *InsertClassRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
