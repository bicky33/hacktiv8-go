package routes

import (
	"final-project/config"
	"final-project/controllers"
	"final-project/middleware"
	"final-project/repositories"
	"final-project/services"
	"final-project/validator"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, db *config.Postgres) {
	// User Service
	repositories := repositories.New(db.DB)
	userService := services.NewUserServie(repositories)
	userController := controllers.UserController{Service: userService, Validate: validator.Validate}
	// Photo Service
	photoService := services.NewPhotoService(repositories)
	photoController := controllers.PhotoController{Service: photoService, Validate: validator.Validate}
	// Comment Service
	commentService := services.NewCommentService(repositories)
	commentController := controllers.CommentController{Service: commentService, Validate: validator.Validate}
	// Social Media Service
	socialMediaService := services.NewCommentService(repositories)
	socialMediaController := controllers.CommentController{Service: socialMediaService, Validate: validator.Validate}

	// Auth
	app.Post("/users/login", userController.Login)
	app.Post("/users/register", userController.Register)
	// user Router
	userRoute := app.Group("/users", middleware.Authentication())
	userRoute.Put("/:userId", middleware.UserAuthorization(), userController.Update)
	userRoute.Delete("/:userId", middleware.UserAuthorization(), userController.Delete)
	// photo Route
	photoRoute := app.Group("/photos", middleware.Authentication())
	photoRoute.Post("/", photoController.Create)
	photoRoute.Get("/", photoController.GetAll)
	photoRoute.Put("/:photoId", middleware.PhotoAuthorization(), photoController.Update)
	photoRoute.Delete("/:photoId", middleware.PhotoAuthorization(), photoController.Delete)
	// commment Route
	commentRoute := app.Group("/comments", middleware.Authentication())
	commentRoute.Post("/", commentController.Create)
	commentRoute.Get("/", commentController.GetAll)
	commentRoute.Put("/:commentId", middleware.CommentAuthorization(), commentController.Update)
	commentRoute.Delete("/:commentId", middleware.CommentAuthorization(), commentController.Delete)
	// social media Route
	socialMediaRoute := app.Group("/socialMedias", middleware.Authentication())
	socialMediaRoute.Post("/", socialMediaController.Create)
	socialMediaRoute.Get("/", socialMediaController.GetAll)
	socialMediaRoute.Put("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaController.Update)
	socialMediaRoute.Delete("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaController.Delete)

}
