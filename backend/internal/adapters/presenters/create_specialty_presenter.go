package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type createSpecialtyPresenter struct{}

func (*createSpecialtyPresenter) Present(ctx context.Context, specialty *domain.Specialty) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Specialty created successfully",
		"data": fiber.Map{
			"specialty": fiber.Map{
				"id":       specialty.ID,
				"name":     specialty.Name,
				"duration": specialty.Duration,
			},
		},
	})
}
func (*createSpecialtyPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewCreateSpecialtyPresenter() application.CreateSpecialtyPresenter {
	return &createSpecialtyPresenter{}
}
