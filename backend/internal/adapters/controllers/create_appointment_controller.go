package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type CreateAppointmentController interface {
	Handle(*fiber.Ctx) error
}

type createAppointmentController struct {
	appointmentUseCase application.CreateAppointmentUseCase
}

func (controller *createAppointmentController) Handle(ctx *fiber.Ctx) error {
	var request application.CreateAppointmentRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if request.PatientID == "" || request.DoctorID == "" || request.SpecialtyID == "" || request.Date == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields must be provided"})
	}
	return controller.appointmentUseCase.Execute(ctx.Context(), request)
}

func NewCreateAppointmentController(appointmentUseCase application.CreateAppointmentUseCase) CreateAppointmentController {
	return &createAppointmentController{
		appointmentUseCase: appointmentUseCase,
	}
}
