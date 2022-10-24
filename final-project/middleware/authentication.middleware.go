package middleware

import (
	"final-project/config"
	"final-project/helper"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		if authorizationHeader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helper.MessageAuthorizedHeaderError})
		}
		if !strings.Contains(authorizationHeader, "Bearer") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helper.MessageInvalidFormatHeaderError})
		}
		config := config.Config()
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		data, err := helper.ValidateToken(tokenString, config.AccessTokenPublicKey)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		c.Locals("current_user", data)
		c.Next()
		return nil
	}
}
