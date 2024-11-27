package vldtor

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func RegisterIsMobile(trans ut.Translator, validate *validator.Validate) {
	base("is_mobile", "{0} 不正确", isMobile, trans, validate)
}

func isMobile(field validator.FieldLevel) bool {
	var (
		matched bool
		err     error
	)

	if matched, err = regexp.Match("1(3|4|5|6|7|8|9)[0-9]{9}", []byte(field.Field().String())); err != nil {
		return false
	}
	return matched
}
