package gateways

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type appointmentRepository struct {
	pool                *pgxpool.Pool
	specialtyRepository domain.SpecialtyRepository
	doctorRepository    domain.DoctorRepository
	patientRepository   domain.PatientRepository
}

func (repository *appointmentRepository) Save(ctx context.Context, appointment *domain.Appointment) error {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, execErr := tx.Exec(ctx, createAppointmentSQL, appointment.ID, appointment.Patient.ID, appointment.Doctor.ID, appointment.Specialty.ID, appointment.Date, appointment.State)
	if execErr != nil {
		return execErr
	}

	return tx.Commit(ctx)
}
func (repository *appointmentRepository) FindAll(ctx context.Context) ([]*domain.Appointment, error) {
	rows, err := repository.pool.Query(ctx, findAllAppointmentsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*domain.Appointment
	for rows.Next() {
		var appointment domain.Appointment
		appointment.Patient = &domain.Patient{}
		appointment.Doctor = &domain.Doctor{}
		appointment.Specialty = &domain.Specialty{}
		if err := rows.Scan(&appointment.ID, &appointment.Patient.ID, &appointment.Doctor.ID, &appointment.Specialty.ID, &appointment.Date, &appointment.State); err != nil {
			return nil, err
		}
		patient, err := repository.patientRepository.FindByID(ctx, appointment.Patient.ID)
		if err != nil {
			return nil, err
		}
		appointment.Patient = patient

		doctor, err := repository.doctorRepository.FindByID(ctx, appointment.Doctor.ID)
		if err != nil {
			return nil, err
		}
		appointment.Doctor = doctor

		specialty, err := repository.specialtyRepository.FindByID(ctx, appointment.Specialty.ID)
		if err != nil {
			return nil, err
		}
		appointment.Specialty = specialty

		appointments = append(appointments, &appointment)
	}

	return appointments, nil
}
func (repository *appointmentRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Appointment, error) {
	row := repository.pool.QueryRow(ctx, findAppointmentByIDSQL, id)

	var appointment domain.Appointment
	appointment.Patient = &domain.Patient{}
	appointment.Doctor = &domain.Doctor{}
	appointment.Specialty = &domain.Specialty{}
	if err := row.Scan(&appointment.ID, &appointment.Patient.ID, &appointment.Doctor.ID, &appointment.Specialty.ID, &appointment.Date, &appointment.State); err != nil {
		return nil, err
	}

	patient, err := repository.patientRepository.FindByID(ctx, appointment.Patient.ID)
	if err != nil {
		return nil, err
	}
	appointment.Patient = patient

	doctor, err := repository.doctorRepository.FindByID(ctx, appointment.Doctor.ID)
	if err != nil {
		return nil, err
	}
	appointment.Doctor = doctor

	specialty, err := repository.specialtyRepository.FindByID(ctx, appointment.Specialty.ID)
	if err != nil {
		return nil, err
	}
	appointment.Specialty = specialty

	return &appointment, nil
}
func (repository *appointmentRepository) FindByDoctor(ctx context.Context, doctorID uuid.UUID, startDate time.Time, endDate time.Time) ([]*domain.Appointment, error) {
	rows, err := repository.pool.Query(ctx, findAppointmentsByDoctorSQL, doctorID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*domain.Appointment
	for rows.Next() {
		var appointment domain.Appointment
		appointment.Patient = &domain.Patient{}
		appointment.Doctor = &domain.Doctor{}
		appointment.Specialty = &domain.Specialty{}
		if err := rows.Scan(&appointment.ID, &appointment.Patient.ID, &appointment.Doctor.ID, &appointment.Specialty.ID, &appointment.Date, &appointment.State); err != nil {
			return nil, err
		}

		patient, err := repository.patientRepository.FindByID(ctx, appointment.Patient.ID)
		if err != nil {
			return nil, err
		}
		appointment.Patient = patient

		doctor, err := repository.doctorRepository.FindByID(ctx, appointment.Doctor.ID)
		if err != nil {
			return nil, err
		}
		appointment.Doctor = doctor

		specialty, err := repository.specialtyRepository.FindByID(ctx, appointment.Specialty.ID)
		if err != nil {
			return nil, err
		}
		appointment.Specialty = specialty

		appointments = append(appointments, &appointment)
	}

	return appointments, nil
}
func (repository *appointmentRepository) FindByPatient(ctx context.Context, patientID uuid.UUID, startDate time.Time, endDate time.Time) ([]*domain.Appointment, error) {
	rows, err := repository.pool.Query(ctx, findAppointmentsByPatientSQL, patientID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*domain.Appointment
	for rows.Next() {
		var appointment domain.Appointment
		appointment.Patient = &domain.Patient{}
		appointment.Doctor = &domain.Doctor{}
		appointment.Specialty = &domain.Specialty{}
		if err := rows.Scan(&appointment.ID, &appointment.Patient.ID, &appointment.Doctor.ID, &appointment.Specialty.ID, &appointment.Date, &appointment.State); err != nil {
			return nil, err
		}

		patient, err := repository.patientRepository.FindByID(ctx, appointment.Patient.ID)
		if err != nil {
			return nil, err
		}
		appointment.Patient = patient

		doctor, err := repository.doctorRepository.FindByID(ctx, appointment.Doctor.ID)
		if err != nil {
			return nil, err
		}
		appointment.Doctor = doctor

		specialty, err := repository.specialtyRepository.FindByID(ctx, appointment.Specialty.ID)
		if err != nil {
			return nil, err
		}
		appointment.Specialty = specialty

		appointments = append(appointments, &appointment)
	}

	return appointments, nil
}

func NewAppointmentRepository(pool *pgxpool.Pool, specialtyRepo domain.SpecialtyRepository, doctorRepo domain.DoctorRepository, patientRepo domain.PatientRepository) domain.AppointmentRepository {
	return &appointmentRepository{
		pool:                pool,
		specialtyRepository: specialtyRepo,
		doctorRepository:    doctorRepo,
		patientRepository:   patientRepo,
	}
}
