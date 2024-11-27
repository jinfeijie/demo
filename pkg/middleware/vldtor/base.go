package vldtor

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func base(tag string, text string, f func(field validator.FieldLevel) bool, trans ut.Translator, validate *validator.Validate) {
	_ = validate.RegisterValidation(tag, f)
	_ = validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, text, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}

func RegisterCustomRule(trans ut.Translator, validate *validator.Validate) {
	RegisterIsMobile(trans, validate)
}
