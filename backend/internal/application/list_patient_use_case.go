package application

import (
	"context"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type ListPatientPresenter interface {
	Present(context.Context, []*domain.Patient) error
	Error(context.Context, error) error
}

type ListPatientUseCase interface {
	Query(context.Context) error
}

type listPatientUseCase struct {
	patientRepository domain.PatientRepository
	patientPresenter  ListPatientPresenter
}

func (useCase *listPatientUseCase) Query(ctx context.Context) error {
	patients, err := useCase.patientRepository.FindAll(ctx)
	if err != nil {
		return useCase.patientPresenter.Error(ctx, err)
	}

	return useCase.patientPresenter.Present(ctx, patients)
}

func NewListPatientUseCase(
	patientRepository domain.PatientRepository,
	patientPresenter ListPatientPresenter,
) ListPatientUseCase {
	return &listPatientUseCase{
		patientRepository: patientRepository,
		patientPresenter:  patientPresenter,
	}
}
