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

	response := dto.UserResponse{Code: fiber.StatusCreated, Message: "succes create data", Status: "Created", Data: data}
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

	response := dto.UserResponse{Code: fiber.StatusOK, Status: "Ok", Message: "Success create token", Data: fiber.Map{"token": token}}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Update(c *fiber.Ctx) error {
	var payload dto.UserUpdateRequest
	userId, _ := c.ParamsInt("userId")
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	payload.ID = int32(userId)
	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validate.TranslateError(err)})
	}
	result, err := controller.Service.Update(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{Code: fiber.StatusOK, Status: "Ok", Message: "Success update data", Data: result}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Delete(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("userId")
	if err := controller.Service.Delete(c.Context(), int32(userId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{Code: fiber.StatusOK, Status: "Ok", Message: "Your account has been successfuly deleted"}
	return c.Status(fiber.StatusOK).JSON(response)
}
