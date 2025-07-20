package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type ListSpecialtyController interface {
	Handle(*fiber.Ctx) error
}

type listSpecialtyController struct {
	specialtyUseCase application.ListSpecialtyUseCase
}

func (controller *listSpecialtyController) Handle(ctx *fiber.Ctx) error {
	return controller.specialtyUseCase.Query(ctx.Context())
}

func NewListSpecialtyController(specialtyUseCase application.ListSpecialtyUseCase) ListSpecialtyController {
	return &listSpecialtyController{
		specialtyUseCase: specialtyUseCase,
	}
}
