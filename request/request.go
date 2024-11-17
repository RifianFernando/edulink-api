package request

import (
	"github.com/go-playground/locales/id" // Indonesian locale
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	// Initialize the Indonesian locale
	idLocale := id.New()
	uni = ut.New(idLocale, idLocale) // Universal Translator setup

	// Get the Indonesian translator
	var found bool
	Trans, found = uni.GetTranslator("id")
	if !found {
		panic("Indonesian translator not found")
	}
	Validate = validator.New()

	registerTranslations()
}

func registerTranslations() {
	_ = registerValidationTranslation("required", "{field} wajib diisi")
	_ = registerValidationTranslation("e164", "nomor telepon harus valid dengan kode negara (misalnya +62)")
	_ = registerValidationTranslation("oneof", "{field} tidak sesuai kriteria: {value}")
	_ = registerValidationTranslation("email", "format email memiliki @ dan .")
	_ = registerValidationTranslation("min", "{field} minimal {min} karakter")
	_ = registerValidationTranslation("max", "{field} maksimal {max} karakter")
	_ = registerValidationTranslation("len", "panjang {field} harus {len} karakter")
}

func registerValidationTranslation(rule, defaultMessage string) error {
	return Validate.RegisterTranslation(rule, Trans, func(ut ut.Translator) error {
		// Add the custom message for each rule
		return ut.Add(rule, defaultMessage, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		// Translate based on rule
		return translateMessage(ut, rule, fe)
	})
}

func translateMessage(ut ut.Translator, rule string, fe validator.FieldError) string {
	// Prepare the map of placeholders for each rule
	params := make([]interface{}, 0)

	// Check the rule and populate params
	params = append(params, fe.Field()) // field name

	switch rule {
	case "oneof":
		params = append(params, fe.Param()) // value for "oneof" rule
	case "min", "max":
		params = append(params, fe.Param()) // min/max values
	case "len":
		params = append(params, fe.Param()) // length value
	}

	// Convert []interface{} to []string for ut.T
	strParams := make([]string, len(params))
	for i, v := range params {
		strParams[i] = v.(string) // Convert each interface{} to string
	}

	// Translate the message and use the params
	translatedMessage, _ := ut.T(rule, strParams...)
	return translatedMessage
}
