package request

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

/*
* ResetPasswordRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type ResetPasswordRequest struct {
	Token           string `json:"token" binding:"required"`
	Email           string `json:"email" binding:"required"`
	NewPassword     string `json:"password" binding:"required"`
	ConfirmPassword string `json:"password_confirmation" binding:"required"`
}

// Validate method
func (r *ResetPasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

/*
* see the documentation about time.Parse here:
* https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
 */
func (r *ResetPasswordRequest) ValidatePasswords() error {
	if strings.Compare(r.NewPassword, r.ConfirmPassword) != 0 {
		return errors.New("passwords do not match")
	}
	return nil
}
