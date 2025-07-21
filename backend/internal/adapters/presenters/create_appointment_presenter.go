package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type createAppointmentPresenter struct{}

func (*createAppointmentPresenter) Present(ctx context.Context, appointment *domain.Appointment) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Appointment created successfully",
		"data": fiber.Map{
			"appointment": appointment,
		},
	})
}
func (*createAppointmentPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewCreateAppointmentPresenter() *createAppointmentPresenter {
	return &createAppointmentPresenter{}
}
