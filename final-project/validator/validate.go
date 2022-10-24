package validator

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate
var Trans ut.Translator

// doesnt use for custom error or tag
// func RegisterTranslation() {
// 	Validate.RegisterTranslation("required", Trans, func(ut ut.Translator) error {
// 		return ut.Add("required", "{0} must not blank", true) // see universal-translator for details
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("required", fe.Field())

// 		return t
// 	})

// 	Validate.RegisterTranslation("email", Trans, func(ut ut.Translator) error {
// 		return ut.Add("email", "{0} must valid email format", true) // see universal-translator for details
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("email", fe.Field())

// 		return t
// 	})

// 	Validate.RegisterTranslation("min", Trans, func(ut ut.Translator) error {
// 		return ut.Add("min", "{0} require minimum {1}", true) // see universal-translator for details
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("min", fe.Field(), fe.Param())

// 		return t
// 	})
// }

func TranslateError(err error) map[string][]string {
	errs := make(map[string][]string)
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(Trans))
		errs[e.Field()] = append(errs[e.Field()], translatedErr.Error())
	}
	return errs
}
