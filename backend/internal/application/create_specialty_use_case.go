package application

import (
	"context"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type CreateSpecialtyRequest struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"` // in minutes
}

type CreateSpecialtyPresenter interface {
	Present(context.Context, *domain.Specialty) error
	Error(context.Context, error) error
}

type CreateSpecialtyUseCase interface {
	Execute(context.Context, CreateSpecialtyRequest) error
}

type createSpecialtyUseCase struct {
	specialtyRepository domain.SpecialtyRepository
	specialtyPresenter  CreateSpecialtyPresenter
}

func (useCase *createSpecialtyUseCase) Execute(ctx context.Context, dto CreateSpecialtyRequest) error {
	specialty := domain.NewSpecialty(dto.Name, dto.Duration)

	if err := useCase.specialtyRepository.Save(ctx, specialty); err != nil {
		return useCase.specialtyPresenter.Error(ctx, err)
	}

	return useCase.specialtyPresenter.Present(ctx, specialty)
}

func NewCreateSpecialtyUseCase(
	specialtyRepository domain.SpecialtyRepository,
	specialtyPresenter CreateSpecialtyPresenter,
) CreateSpecialtyUseCase {
	return &createSpecialtyUseCase{
		specialtyRepository: specialtyRepository,
		specialtyPresenter:  specialtyPresenter,
	}
}
