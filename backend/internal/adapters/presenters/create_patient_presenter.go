package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type createPatientPresenter struct{}

func (*createPatientPresenter) Present(ctx context.Context, patient *domain.Patient) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Patient created successfully",
		"data": fiber.Map{
			"patient": patient,
		},
	})
}
func (*createPatientPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewCreatePatientPresenter() application.CreatePatientPresenter {
	return &createPatientPresenter{}
}
