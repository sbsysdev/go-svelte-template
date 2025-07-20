package application

import (
	"context"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type CreateSpecialityRequest struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"` // in minutes
}

type CreateSpecialityPresenter interface {
	Present(context.Context, *domain.Speciality) error
	Error(context.Context, error) error
}

type CreateSpecialityUseCase interface {
	Execute(context.Context, CreateSpecialityRequest) error
}

type createSpecialityUseCase struct {
	specialityRepository domain.SpecialityRepository
	specialityResponse   CreateSpecialityPresenter
}

func (useCase *createSpecialityUseCase) Execute(ctx context.Context, dto CreateSpecialityRequest) error {
	speciality := domain.NewSpeciality(dto.Name, dto.Duration)

	if err := useCase.specialityRepository.Save(ctx, speciality); err != nil {
		return useCase.specialityResponse.Error(ctx, err)
	}

	return useCase.specialityResponse.Present(ctx, speciality)
}

func NewCreateSpecialityUseCase(
	specialityRepository domain.SpecialityRepository,
	specialityResponse CreateSpecialityPresenter,
) CreateSpecialityUseCase {
	return &createSpecialityUseCase{
		specialityRepository: specialityRepository,
		specialityResponse:   specialityResponse,
	}
}
