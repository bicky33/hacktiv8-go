package controllers

import (
	"final-project/dto"
	"final-project/helper"
	"final-project/services"
	validate "final-project/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CommentController struct {
	Service  services.CommentService
	Validate *validator.Validate
}

func (controller *CommentController) Create(c *fiber.Ctx) error {
	var payload dto.CommentCreateRequest
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

func (controller *CommentController) GetAll(c *fiber.Ctx) error {
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

func (controller *CommentController) Update(c *fiber.Ctx) error {
	commentId, _ := c.ParamsInt("commentId")
	var payload dto.CommentUpdateRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := controller.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validate.TranslateError(err))
	}

	result, err := controller.Service.Update(c.Context(), payload, uint32(commentId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.CommentResponse{
		Status: fiber.StatusOK,
		Data:   result,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *CommentController) Delete(c *fiber.Ctx) error {
	commentId, _ := c.ParamsInt("commentId")
	if err := controller.Service.Delete(c.Context(), uint32(commentId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Your comment has been successfuly deleted"}
	return c.Status(fiber.StatusOK).JSON(response)
}
