package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type CreatePatientController interface {
	Handle(*fiber.Ctx) error
}

type createPatientController struct {
	patientUseCase application.CreatePatientUseCase
}

func (controller *createPatientController) Handle(ctx *fiber.Ctx) error {
	var request application.CreatePatientRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if request.Name == "" || request.Birth == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and Birth must be provided"})
	}
	return controller.patientUseCase.Execute(ctx.Context(), request)
}

func NewCreatePatientController(patientUseCase application.CreatePatientUseCase) CreatePatientController {
	return &createPatientController{
		patientUseCase: patientUseCase,
	}
}
