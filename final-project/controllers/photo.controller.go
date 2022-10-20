package controllers

import (
	"final-project/dto"
	"final-project/helper"
	"final-project/services"
	validate "final-project/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PhotoController struct {
	Service  services.PhotoService
	Validate *validator.Validate
}

func (controller *PhotoController) Create(c *fiber.Ctx) error {
	current_user := c.Locals("current_user").(*helper.JWTClaim)
	var payload dto.PhotoCreateRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validate.TranslateError(err))
	}
	result, err := controller.Service.Create(c.Context(), payload, current_user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.PhotoResponse{
		Code:    fiber.StatusCreated,
		Status:  "Created",
		Message: "Success create a data",
		Data:    result,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller *PhotoController) GetAll(c *fiber.Ctx) error {
	result, err := controller.Service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.PhotoResponse{
		Code:    fiber.StatusOK,
		Status:  "Ok",
		Message: "success get data",
		Data:    result,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller PhotoController) Update(c *fiber.Ctx) error {
	photoId, _ := c.ParamsInt("photoId")
	var payload dto.PhotoCreateRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validate.TranslateError(err))
	}

	result, err := controller.Service.Update(c.Context(), payload, int32(photoId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.PhotoResponse{
		Code:    fiber.StatusOK,
		Status:  "Ok",
		Message: "success update data",
		Data:    result,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *PhotoController) Delete(c *fiber.Ctx) error {
	photoId, _ := c.ParamsInt("photoId")
	if err := controller.Service.Delete(c.Context(), int32(photoId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{Code: fiber.StatusOK, Status: "Ok", Message: "Your photo has been successfuly deleted"}
	return c.Status(fiber.StatusOK).JSON(response)
}
