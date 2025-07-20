package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type CreateDoctorRequest struct {
	Name        string   `json:"name"`
	Specialties []string `json:"specialties"` // IDs of the specialties
}

type CreateDoctorPresenter interface {
	Present(context.Context, *domain.Doctor) error
	Error(context.Context, error) error
}

type CreateDoctorUseCase interface {
	Execute(context.Context, CreateDoctorRequest) error
}

type createDoctorUseCase struct {
	doctorRepository    domain.DoctorRepository
	doctorPresenter     CreateDoctorPresenter
	specialtyRepository domain.SpecialtyRepository
}

func (useCase *createDoctorUseCase) Execute(ctx context.Context, dto CreateDoctorRequest) error {
	specialties := make([]*domain.Specialty, 0, len(dto.Specialties))
	for _, specialtyID := range dto.Specialties {
		specialty, err := useCase.specialtyRepository.FindByID(ctx, uuid.MustParse(specialtyID))
		if err != nil {
			return useCase.doctorPresenter.Error(ctx, err)
		}
		specialties = append(specialties, specialty)
	}

	doctor := domain.NewDoctor(dto.Name, specialties)

	if err := useCase.doctorRepository.Save(ctx, doctor); err != nil {
		return useCase.doctorPresenter.Error(ctx, err)
	}

	return useCase.doctorPresenter.Present(ctx, doctor)
}

func NewCreateDoctorUseCase(
	doctorRepository domain.DoctorRepository,
	doctorPresenter CreateDoctorPresenter,
	specialtyRepository domain.SpecialtyRepository,
) CreateDoctorUseCase {
	return &createDoctorUseCase{
		doctorRepository:    doctorRepository,
		doctorPresenter:     doctorPresenter,
		specialtyRepository: specialtyRepository,
	}
}
