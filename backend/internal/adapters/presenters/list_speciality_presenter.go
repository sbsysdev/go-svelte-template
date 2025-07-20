package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type listSpecialityPresenter struct{}

func (p *listSpecialityPresenter) Present(ctx context.Context, specialities []*domain.Speciality) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Specialities retrieved successfully",
		"data": fiber.Map{
			"specialities": specialities,
		},
	})
}

func (p *listSpecialityPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewListSpecialityPresenter() *listSpecialityPresenter {
	return &listSpecialityPresenter{}
}
