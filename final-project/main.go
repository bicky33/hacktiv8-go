package main

import (
	"final-project/config"
	"final-project/routes"
	"final-project/validator"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validate "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	en := en.New()
	uni := ut.New(en, en)
	validator.Trans, _ = uni.GetTranslator("en")
	validator.Validate = validate.New()
	en_translations.RegisterDefaultTranslations(validator.Validate, validator.Trans)
	// validator.RegisterTranslation()
	routes.Router(app, db)
	app.Listen(":3000")
}
