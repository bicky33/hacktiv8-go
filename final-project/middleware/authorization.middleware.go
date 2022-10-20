package middleware

import (
	"errors"
	"final-project/config"
	"final-project/helper"
	"final-project/repositories"

	"github.com/gofiber/fiber/v2"
)

func UserAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		userIdParams, err := c.ParamsInt("userId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "params userId must be integer"})
		}
		if int32(userIdParams) != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized action"})
		}
		return c.Next()
	}
}

func PhotoAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		photoIdParams, err := c.ParamsInt("photoId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "params userId must be integer"})
		}
		db := config.GetDB()
		repo := repositories.New(db.DB)
		photoData, err := repo.GetPhotoById(c.Context(), int32(photoIdParams))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("data not found").Error()})
		}

		if photoData.UserID != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized action"})
		}
		return c.Next()
	}
}

func CommentAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		commentIdParams, err := c.ParamsInt("commentId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "params commentId must be integer"})
		}
		db := config.GetDB()
		repo := repositories.New(db.DB)
		commentData, err := repo.GetCommentById(c.Context(), int32(commentIdParams))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("data not found").Error()})
		}

		if commentData.UserID != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized action"})
		}
		return c.Next()
	}
}

func SocialMediaAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		socialMediaIdParams, err := c.ParamsInt("socialMediaId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "params socialMediaId must be integer"})
		}
		db := config.GetDB()
		repo := repositories.New(db.DB)
		socialMediaData, err := repo.GetSocialMediaById(c.Context(), int32(socialMediaIdParams))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.New("data not found").Error()})
		}

		if socialMediaData.UserID != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized action"})
		}
		return c.Next()
	}
}
