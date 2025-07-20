package application

import (
	"context"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type ListSpecialityPresenter interface {
	Present(context.Context, []*domain.Speciality) error
	Error(context.Context, error) error
}

type ListSpecialityUseCase interface {
	Query(context.Context) error
}

type listSpecialityUseCase struct {
	specialityRepository domain.SpecialityRepository
	specialityPresenter  ListSpecialityPresenter
}

func (useCase *listSpecialityUseCase) Query(ctx context.Context) error {
	specialties, err := useCase.specialityRepository.FindAll(ctx)
	if err != nil {
		return useCase.specialityPresenter.Error(ctx, err)
	}

	return useCase.specialityPresenter.Present(ctx, specialties)
}

func NewListSpecialityUseCase(
	specialityRepository domain.SpecialityRepository,
	specialityPresenter ListSpecialityPresenter,
) ListSpecialityUseCase {
	return &listSpecialityUseCase{
		specialityRepository: specialityRepository,
		specialityPresenter:  specialityPresenter,
	}
}
