package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/skripsi-be/models"
)

type InsertClassRequest struct {
	models.ClassName
}

// Validate method
func (r *InsertClassRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
