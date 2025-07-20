package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type CreateSpecialtyController interface {
	Handle(*fiber.Ctx) error
}

type createSpecialtyController struct {
	specialtyUseCase application.CreateSpecialtyUseCase
}

func (controller *createSpecialtyController) Handle(ctx *fiber.Ctx) error {
	var request application.CreateSpecialtyRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if request.Name == "" || request.Duration <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and Duration must be provided"})
	}
	return controller.specialtyUseCase.Execute(ctx.Context(), request)
}

func NewCreateSpecialtyController(specialtyUseCase application.CreateSpecialtyUseCase) CreateSpecialtyController {
	return &createSpecialtyController{
		specialtyUseCase: specialtyUseCase,
	}
}
