package controllers

import (
	"final-project/config"
	"final-project/dto"
	"final-project/helper"
	"final-project/services"
	validate "final-project/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service  services.UserService
	Validate *validator.Validate
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var payload dto.UserCreateRequest
	var err error
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validate.TranslateError(err)})
	}
	payload.Password, err = helper.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := controller.Service.Create(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.UserResponse{
		Status: fiber.StatusCreated,
		Data:   data,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var payload dto.UserLogin
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	token, err := controller.Service.Login(c.Context(), payload)
	config := config.Config()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(
		&fiber.Cookie{
			Name:     "access_token",
			Value:    token,
			MaxAge:   config.AccessTokenMaxAge + 60,
			Secure:   false,
			HTTPOnly: true,
			Domain:   "localhost",
			Path:     "/",
		},
	)

	c.Cookie(
		&fiber.Cookie{
			Name:     "logged_in",
			Value:    "true",
			MaxAge:   config.AccessTokenMaxAge + 60,
			Secure:   false,
			HTTPOnly: false,
			Domain:   "localhost",
			Path:     "/",
		},
	)

	response := dto.UserResponse{
		Status: fiber.StatusOK,
		Data:   fiber.Map{"token": token},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Update(c *fiber.Ctx) error {
	var payload dto.UserUpdateRequest
	getUser := c.Locals("current_user").(*helper.JWTClaim)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validate.TranslateError(err)})
	}
	result, err := controller.Service.Update(c.Context(), payload, getUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{
		Status: fiber.StatusOK,
		Data:   result,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Delete(c *fiber.Ctx) error {
	getUser := c.Locals("current_user").(*helper.JWTClaim)
	if err := controller.Service.Delete(c.Context(), uint32(getUser.ID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Your account has been successfuly deleted",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
