// requests/create_staff_request.go
package request

import (
	"time"

	"github.com/go-playground/validator/v10"
)

/*
* UpdateStaffRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type UpdateStaffRequest struct {
	UserName         string  `json:"name" binding:"required"`
	UserGender       string  `json:"gender" binding:"required,oneof=Male Female"`
	UserPlaceOfBirth string  `json:"place_of_birth" binding:"required"`
	UserReligion     string  `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	DateOfBirth      string  `json:"date_of_birth" binding:"required"`
	UserAddress      string  `json:"address" binding:"required"`
	UserPhoneNum     string  `json:"num_phone" binding:"required,e164"`
	UserEmail        string  `json:"email" binding:"required,email"`
	Position    	 string  `json:"position" binding:"required" validate:"required"`
}

// Validate method
func (r *UpdateStaffRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

/*
* see the documentation about time.Parse here:
* https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
 */
func (r *UpdateStaffRequest) ParseDates() (time.Time, error) {
	DateOfBirth, err := time.Parse("2006-01-02", r.DateOfBirth)
	if err != nil {
		return time.Time{}, err
	}

	return DateOfBirth, nil
}
