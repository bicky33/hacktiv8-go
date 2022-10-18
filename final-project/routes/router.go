package routes

import (
	"final-project/config"
	"final-project/controllers"
	"final-project/repositories"
	"final-project/services"
	"final-project/validator"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, db *config.Postgres) {
	// User
	repositories := repositories.New(db.DB)
	userService := services.NewUserServie(repositories)
	userController := controllers.UserController{Service: userService, Validate: validator.Validate}
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("hello guys")
	})
	app.Post("/users", userController.Register)

}
