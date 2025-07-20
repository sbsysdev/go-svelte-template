package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type ListSpecialityController interface {
	Handle(*fiber.Ctx) error
}

type listSpecialityController struct {
	specialityUseCase application.ListSpecialityUseCase
}

func (controller *listSpecialityController) Handle(ctx *fiber.Ctx) error {
	return controller.specialityUseCase.Query(ctx.Context())
}

func NewListSpecialityController(specialityUseCase application.ListSpecialityUseCase) ListSpecialityController {
	return &listSpecialityController{
		specialityUseCase: specialityUseCase,
	}
}
