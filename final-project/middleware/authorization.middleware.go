package middleware

import (
	"final-project/config"
	"final-project/helper"
	"final-project/repositories"

	"github.com/gofiber/fiber/v2"
)

func PhotoAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		photoIdParams, err := c.ParamsInt("photoId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helper.MessageUserParamsIdError})
		}
		db := config.GetDB()
		repo := repositories.New(db.DB)
		photoData, err := repo.GetPhotoById(c.Context(), uint32(photoIdParams))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helper.MessageDataNotFountError})
		}

		if photoData.UserID != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": helper.MessageUnauthorizedError})
		}
		return c.Next()
	}
}

func CommentAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		commentIdParams, err := c.ParamsInt("commentId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helper.MessageCommentParamsIdError})
		}
		db := config.GetDB()
		repo := repositories.New(db.DB)
		commentData, err := repo.GetCommentById(c.Context(), uint32(commentIdParams))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helper.MessageDataNotFountError})
		}

		if commentData.UserID != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": helper.MessageUnauthorizedError})
		}
		return c.Next()
	}
}

func SocialMediaAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getData := c.Locals("current_user").(*helper.JWTClaim)
		socialMediaIdParams, err := c.ParamsInt("socialMediaId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helper.MessageSocialMediaParamsIdError})
		}
		db := config.GetDB()
		repo := repositories.New(db.DB)
		socialMediaData, err := repo.GetSocialMediaById(c.Context(), uint32(socialMediaIdParams))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helper.MessageSocialMediaParamsIdError})
		}

		if socialMediaData.UserID != getData.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": helper.MessageUnauthorizedError})
		}
		return c.Next()
	}
}
