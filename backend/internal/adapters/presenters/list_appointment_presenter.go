package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type listAppointmentPresenter struct{}

func (*listAppointmentPresenter) Present(ctx context.Context, appointments []*domain.Appointment) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Appointments retrieved successfully",
		"data": fiber.Map{
			"appointments": appointments,
		},
	})
}
func (*listAppointmentPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewListAppointmentPresenter() application.ListAppointmentPresenter {
	return &listAppointmentPresenter{}
}
