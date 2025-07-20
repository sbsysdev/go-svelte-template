package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type ListDoctorBySpecialtyPresenter interface {
	Present(context.Context, []*domain.Doctor) error
	Error(context.Context, error) error
}

type ListDoctorBySpecialtyUseCase interface {
	Query(ctx context.Context, specialtyID string) error
}

type listDoctorBySpecialtyUseCase struct {
	doctorRepository domain.DoctorRepository
	doctorPresenter  ListDoctorBySpecialtyPresenter
}

func (useCase *listDoctorBySpecialtyUseCase) Query(ctx context.Context, specialtyID string) error {
	doctors, err := useCase.doctorRepository.FindBySpecialtyID(ctx, uuid.MustParse(specialtyID))
	if err != nil {
		return useCase.doctorPresenter.Error(ctx, err)
	}

	return useCase.doctorPresenter.Present(ctx, doctors)
}

func NewListDoctorBySpecialtyUseCase(
	doctorRepository domain.DoctorRepository,
	doctorPresenter ListDoctorBySpecialtyPresenter,
) ListDoctorBySpecialtyUseCase {
	return &listDoctorBySpecialtyUseCase{
		doctorRepository: doctorRepository,
		doctorPresenter:  doctorPresenter,
	}
}
