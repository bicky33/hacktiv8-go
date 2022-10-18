package validator_test

import (
	"final-project/validator"
	"fmt"
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validate "github.com/go-playground/validator/v10"
)

func TestValidate(t *testing.T) {
	var test struct {
		Username string `validate:"required"`
		Age      int    `validate:"min=8"`
		Data     struct {
			Email string `validate:"email,required"`
		}
	}
	en := en.New()
	uni := ut.New(en, en)
	validator.Trans, _ = uni.GetTranslator("en")
	validator.Validate = validate.New()
	err := validator.Validate.Struct(test)
	if err != nil {
		heh := validator.TranslateError(err)
		fmt.Println(heh)
	}
}
