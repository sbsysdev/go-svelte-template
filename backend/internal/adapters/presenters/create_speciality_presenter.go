package presenters

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type createSpecialityPresenter struct{}

func (presenter *createSpecialityPresenter) Present(ctx context.Context, speciality *domain.Speciality) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Speciality created successfully",
		"data": fiber.Map{
			"speciality": fiber.Map{
				"id":       speciality.ID,
				"name":     speciality.Name,
				"duration": speciality.Duration,
			},
		},
	})
}
func (presenter *createSpecialityPresenter) Error(ctx context.Context, err error) error {
	fiberCtx := ctx.Value("fiberContext").(*fiber.Ctx)
	return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func NewCreateSpecialityPresenter() *createSpecialityPresenter {
	return &createSpecialityPresenter{}
}
