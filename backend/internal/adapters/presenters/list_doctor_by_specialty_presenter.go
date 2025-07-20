package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type listDoctorBySpecialtyPresenter struct{}

func (*listDoctorBySpecialtyPresenter) Present(ctx context.Context, doctors []*domain.Doctor) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Doctors retrieved successfully",
		"data": fiber.Map{
			"doctors": doctors,
		},
	})
}
func (*listDoctorBySpecialtyPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewListDoctorBySpecialtyPresenter() application.ListDoctorBySpecialtyPresenter {
	return &listDoctorBySpecialtyPresenter{}
}
