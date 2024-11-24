package res

import (
	"net/http"

	"github.com/edulink-api/request"
	"github.com/go-playground/validator/v10"
)

var (
	Forbidden = http.StatusText(http.StatusForbidden)
	Success   = http.StatusText(http.StatusOK)
)

func ResponseMessage(err error) []map[string]string {
	var errorMessages []map[string]string
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range validationErrors {
				translatedError := ve.Translate(request.Trans)
				errorMessages = append(errorMessages, map[string]string{
					"field":   ve.Field(),
					"message": translatedError,
				})
			}
		} else {
			errorMessages = append(errorMessages, map[string]string{
				"error": err.Error(),
			})
		}
	}
	return errorMessages
}
