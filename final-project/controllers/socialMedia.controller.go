package controllers

import (
	"final-project/dto"
	"final-project/helper"
	"final-project/services"
	validate "final-project/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SocialMediaController struct {
	Service  services.SocialMedia
	Validate *validator.Validate
}

func (controller *SocialMediaController) Create(c *fiber.Ctx) error {
	var payload dto.SocialMediaCreateRequest
	var err error
	current_user := c.Locals("current_user").(*helper.JWTClaim)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validate.TranslateError(err)})
	}

	result, err := controller.Service.Create(c.Context(), payload, current_user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.CommentResponse{
		Status: fiber.StatusCreated,
		Data:   result,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller *SocialMediaController) GetAll(c *fiber.Ctx) error {
	result, err := controller.Service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.CommentResponse{
		Status: fiber.StatusOK,
		Data:   result,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *SocialMediaController) Update(c *fiber.Ctx) error {
	socialMediaId, _ := c.ParamsInt("socialMediaId")
	var payload dto.SocialMediaCreateRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validate.TranslateError(err))
	}

	result, err := controller.Service.Update(c.Context(), payload, uint32(socialMediaId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.CommentResponse{
		Status: fiber.StatusOK,
		Data:   result,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *SocialMediaController) Delete(c *fiber.Ctx) error {
	socialMediaId, _ := c.ParamsInt("socialMediaId")
	if err := controller.Service.Delete(c.Context(), uint32(socialMediaId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Your Social Media has been successfuly deleted",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
