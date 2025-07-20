package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type ListDoctorBySpecialtyController interface {
	Handle(*fiber.Ctx) error
}

type listDoctorBySpecialtyController struct {
	doctorUseCase application.ListDoctorBySpecialtyUseCase
}

func (controller *listDoctorBySpecialtyController) Handle(ctx *fiber.Ctx) error {
	specialtyID := ctx.Params("specialtyID")
	if specialtyID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Specialty ID is required",
		})
	}

	return controller.doctorUseCase.Query(ctx.Context(), specialtyID)
}

func NewListDoctorBySpecialtyController(doctorUseCase application.ListDoctorBySpecialtyUseCase) ListDoctorBySpecialtyController {
	return &listDoctorBySpecialtyController{
		doctorUseCase: doctorUseCase,
	}
}
