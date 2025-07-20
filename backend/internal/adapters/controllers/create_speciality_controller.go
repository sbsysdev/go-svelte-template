package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type CreateSpecialityController interface {
	Handle(*fiber.Ctx) error
}

type createSpecialityController struct {
	specialityUseCase application.CreateSpecialityUseCase
}

func (controller *createSpecialityController) Handle(ctx *fiber.Ctx) error {
	var request application.CreateSpecialityRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if request.Name == "" || request.Duration <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and Duration must be provided"})
	}
	return controller.specialityUseCase.Execute(ctx.Context(), request)
}

func NewCreateSpecialityController(specialityUseCase application.CreateSpecialityUseCase) CreateSpecialityController {
	return &createSpecialityController{
		specialityUseCase: specialityUseCase,
	}
}
