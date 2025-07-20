package application

import (
	"context"
	"time"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type CreatePatientRequest struct {
	Name  string `json:"name"`
	Birth string `json:"birth"`
}

type CreatePatientPresenter interface {
	Present(context.Context, *domain.Patient) error
	Error(context.Context, error) error
}

type CreatePatientUseCase interface {
	Execute(context.Context, CreatePatientRequest) error
}

type createPatientUseCase struct {
	patientRepository domain.PatientRepository
	patientPresenter  CreatePatientPresenter
}

func (useCase *createPatientUseCase) Execute(ctx context.Context, dto CreatePatientRequest) error {
	birth, err := time.Parse(time.DateOnly, dto.Birth)
	if err != nil {
		return useCase.patientPresenter.Error(ctx, err)
	}
	patient := domain.NewPatient(dto.Name, birth)

	if err := useCase.patientRepository.Save(ctx, patient); err != nil {
		return useCase.patientPresenter.Error(ctx, err)
	}

	return useCase.patientPresenter.Present(ctx, patient)
}

func NewCreatePatientUseCase(
	patientRepository domain.PatientRepository,
	patientPresenter CreatePatientPresenter,
) CreatePatientUseCase {
	return &createPatientUseCase{
		patientRepository: patientRepository,
		patientPresenter:  patientPresenter,
	}
}
