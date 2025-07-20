package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type listPatientPresenter struct{}

func (*listPatientPresenter) Present(ctx context.Context, patients []*domain.Patient) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Patients retrieved successfully",
		"data": fiber.Map{
			"patients": patients,
		},
	})
}
func (*listPatientPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewListPatientPresenter() application.ListPatientPresenter {
	return &listPatientPresenter{}
}
