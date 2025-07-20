package application

import (
	"context"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type ListSpecialtyPresenter interface {
	Present(context.Context, []*domain.Specialty) error
	Error(context.Context, error) error
}

type ListSpecialtyUseCase interface {
	Query(context.Context) error
}

type listSpecialtyUseCase struct {
	specialtyRepository domain.SpecialtyRepository
	specialtyPresenter  ListSpecialtyPresenter
}

func (useCase *listSpecialtyUseCase) Query(ctx context.Context) error {
	specialties, err := useCase.specialtyRepository.FindAll(ctx)
	if err != nil {
		return useCase.specialtyPresenter.Error(ctx, err)
	}

	return useCase.specialtyPresenter.Present(ctx, specialties)
}

func NewListSpecialtyUseCase(
	specialtyRepository domain.SpecialtyRepository,
	specialtyPresenter ListSpecialtyPresenter,
) ListSpecialtyUseCase {
	return &listSpecialtyUseCase{
		specialtyRepository: specialtyRepository,
		specialtyPresenter:  specialtyPresenter,
	}
}
