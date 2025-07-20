package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type ListPatientController interface {
	Handle(*fiber.Ctx) error
}

type listPatientController struct {
	patientUseCase application.ListPatientUseCase
}

func (controller *listPatientController) Handle(ctx *fiber.Ctx) error {
	return controller.patientUseCase.Query(ctx.Context())
}

func NewListPatientController(patientUseCase application.ListPatientUseCase) ListPatientController {
	return &listPatientController{
		patientUseCase: patientUseCase,
	}
}
