package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type ListAppointmentController interface {
	Handle(*fiber.Ctx) error
}

type listAppointmentController struct {
	appointmentUseCase application.ListAppointmentUseCase
}

func (controller *listAppointmentController) Handle(ctx *fiber.Ctx) error {
	return controller.appointmentUseCase.Query(ctx.Context())
}

func NewListAppointmentController(appointmentUseCase application.ListAppointmentUseCase) ListAppointmentController {
	return &listAppointmentController{
		appointmentUseCase: appointmentUseCase,
	}
}
