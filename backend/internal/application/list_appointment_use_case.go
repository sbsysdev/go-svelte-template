package application

import (
	"context"

	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type ListAppointmentPresenter interface {
	Present(context.Context, []*domain.Appointment) error
	Error(context.Context, error) error
}

type ListAppointmentUseCase interface {
	Query(context.Context) error
}

type listAppointmentUseCase struct {
	appointmentRepository domain.AppointmentRepository
	appointmentPresenter  ListAppointmentPresenter
}

func (useCase *listAppointmentUseCase) Query(ctx context.Context) error {
	appointments, err := useCase.appointmentRepository.FindAll(ctx)
	if err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	return useCase.appointmentPresenter.Present(ctx, appointments)
}

func NewListAppointmentUseCase(
	appointmentRepository domain.AppointmentRepository,
	appointmentPresenter ListAppointmentPresenter,
) ListAppointmentUseCase {
	return &listAppointmentUseCase{
		appointmentRepository: appointmentRepository,
		appointmentPresenter:  appointmentPresenter,
	}
}
