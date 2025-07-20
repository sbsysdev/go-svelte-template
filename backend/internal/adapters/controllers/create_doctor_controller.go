package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type CreateDoctorController interface {
	Handle(*fiber.Ctx) error
}

type createDoctorController struct {
	doctorUseCase application.CreateDoctorUseCase
}

func (controller *createDoctorController) Handle(ctx *fiber.Ctx) error {
	var request application.CreateDoctorRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if request.Name == "" || len(request.Specialties) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and Specialties must be provided"})
	}
	return controller.doctorUseCase.Execute(ctx.Context(), request)
}

func NewCreateDoctorController(doctorUseCase application.CreateDoctorUseCase) CreateDoctorController {
	return &createDoctorController{
		doctorUseCase: doctorUseCase,
	}
}
