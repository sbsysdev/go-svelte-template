package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type CreateAppointmentRequest struct {
	PatientID   string `json:"patient_id"`
	DoctorID    string `json:"doctor_id"`
	SpecialtyID string `json:"specialty_id"`
	Date        string `json:"date"`
}

type CreateAppointmentPresenter interface {
	Present(context.Context, *domain.Appointment) error
	Error(context.Context, error) error
}

type CreateAppointmentUseCase interface {
	Execute(context.Context, CreateAppointmentRequest) error
}

type createAppointmentUseCase struct {
	patientRepository     domain.PatientRepository
	doctorRepository      domain.DoctorRepository
	specialtyRepository   domain.SpecialtyRepository
	appointmentRepository domain.AppointmentRepository
	appointmentPresenter  CreateAppointmentPresenter
	appointmentGuard      *domain.AppointmentGuard
}

func (useCase *createAppointmentUseCase) Execute(ctx context.Context, dto CreateAppointmentRequest) error {
	patient, err := useCase.patientRepository.FindByID(ctx, uuid.MustParse(dto.PatientID))
	if err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	doctor, err := useCase.doctorRepository.FindByID(ctx, uuid.MustParse(dto.DoctorID))
	if err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	specialty, err := useCase.specialtyRepository.FindByID(ctx, uuid.MustParse(dto.SpecialtyID))
	if err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	if err := doctor.HasSpecialty(specialty); err != nil {
		return err
	}

	date, err := time.Parse(time.RFC3339, dto.Date)
	if err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	appointment := domain.NewAppointment(patient, doctor, specialty, date)

	startDate, endDate := domain.GetDayRangeFromDate(date)
	if err := useCase.appointmentGuard.CheckAvailabilityByDate(ctx, appointment, startDate, endDate); err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	if err := useCase.appointmentRepository.Save(ctx, appointment); err != nil {
		return useCase.appointmentPresenter.Error(ctx, err)
	}

	return useCase.appointmentPresenter.Present(ctx, appointment)
}

func NewCreateAppointmentUseCase(
	patientRepository domain.PatientRepository,
	doctorRepository domain.DoctorRepository,
	specialtyRepository domain.SpecialtyRepository,
	appointmentRepository domain.AppointmentRepository,
	appointmentPresenter CreateAppointmentPresenter,
	appointmentGuard *domain.AppointmentGuard,
) CreateAppointmentUseCase {
	return &createAppointmentUseCase{
		patientRepository:     patientRepository,
		doctorRepository:      doctorRepository,
		specialtyRepository:   specialtyRepository,
		appointmentRepository: appointmentRepository,
		appointmentPresenter:  appointmentPresenter,
		appointmentGuard:      appointmentGuard,
	}
}
