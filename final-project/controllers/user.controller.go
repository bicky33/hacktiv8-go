package controllers

import (
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
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := controller.Validate.Struct(payload)
	if err != nil {
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

	responseData := dto.UserCreateResponse{ID: int(data.ID), Age: int(data.Age), Username: data.Username, Email: data.Email}

	response := dto.UserResponse{Code: fiber.StatusOK, Message: "succes create data", Status: "Ok", Data: responseData}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var payload dto.UserLogin
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
}
