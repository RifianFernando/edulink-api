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
	_ = registerValidationTranslation("required", "{0} wajib diisi", Trans, Validate)
	_ = registerValidationTranslation("e164", "nomor telepon harus valid dengan kode negara (misalnya +62)", Trans, Validate)
	_ = registerValidationTranslation("oneof", "{0} tidak sesuai kriteria: {1}", Trans, Validate)
	_ = registerValidationTranslation("email", "format email memiliki @ dan .", Trans, Validate)
}

func registerValidationTranslation(rule, defaultMessage string, trans ut.Translator, validate *validator.Validate) error {
	return validate.RegisterTranslation(rule, trans, func(ut ut.Translator) error {
		return ut.Add(rule, defaultMessage, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		translatedMessage, _ := ut.T("oneof", fe.Field(), fe.Param())
		return translatedMessage
	})
}
