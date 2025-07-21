package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AppointmentState string

const (
	AppointmentStateScheduled AppointmentState = "scheduled"
	AppointmentStateCompleted AppointmentState = "completed"
	AppointmentStateCancelled AppointmentState = "cancelled"
)

const errAppointmentDateOverlap = "appointment date overlaps with an existing appointment"

type Appointment struct {
	ID        uuid.UUID
	Patient   *Patient
	Doctor    *Doctor
	Specialty *Specialty
	Date      time.Time
	State     AppointmentState
}

func NewAppointment(patient *Patient, doctor *Doctor, specialty *Specialty, date time.Time) *Appointment {
	return &Appointment{
		ID:        uuid.New(),
		Patient:   patient,
		Doctor:    doctor,
		Specialty: specialty,
		Date:      date,
		State:     AppointmentStateScheduled,
	}
}

type AppointmentRepository interface {
	Save(context.Context, *Appointment) error
	FindByID(context.Context, uuid.UUID) (*Appointment, error)
	FindByDoctor(context.Context, uuid.UUID, time.Time, time.Time) ([]*Appointment, error)
	FindByPatient(context.Context, uuid.UUID, time.Time, time.Time) ([]*Appointment, error)
}

type AppointmentGuard struct {
	appointmentRepository AppointmentRepository
}

func NewAppointmentGuard(appointmentRepository AppointmentRepository) *AppointmentGuard {
	return &AppointmentGuard{
		appointmentRepository: appointmentRepository,
	}
}

func (guard *AppointmentGuard) CheckAvailabilityByDate(ctx context.Context, appointment *Appointment, startDate time.Time, endDate time.Time) error {
	// Check if the doctor is available at the requested time
	doctorAppointments, doctorErr := guard.appointmentRepository.FindByDoctor(ctx, appointment.Doctor.ID, startDate, endDate)
	if doctorErr != nil {
		return doctorErr
	}
	if err := guard.checkAvailabilityByAppointments(doctorAppointments, appointment); err != nil {
		return err
	}

	// Check if the patient has any conflicting appointments
	patientAppointments, patientErr := guard.appointmentRepository.FindByPatient(ctx, appointment.Patient.ID, startDate, endDate)
	if patientErr != nil {
		return patientErr
	}
	if err := guard.checkAvailabilityByAppointments(patientAppointments, appointment); err != nil {
		return err
	}

	// If no conflicts found, the appointment can be scheduled
	return nil
}
func (guard *AppointmentGuard) checkAvailabilityByAppointments(appointments []*Appointment, appointment *Appointment) error {
	for _, existingAppointment := range appointments {
		if existingAppointment.State != AppointmentStateScheduled {
			continue
		}

		existingStartDate := existingAppointment.Date
		existingEndDate := existingStartDate.Add(time.Duration(appointment.Specialty.Duration) * time.Minute)

		requestedStartDate := appointment.Date
		requestedEndDate := requestedStartDate.Add(time.Duration(appointment.Specialty.Duration) * time.Minute)

		if requestedStartDate.Compare(existingStartDate) >= 0 && requestedStartDate.Before(existingEndDate) {
			return errors.New(errAppointmentDateOverlap)
		}
		if requestedEndDate.After(existingStartDate) && requestedEndDate.Before(existingEndDate) {
			return errors.New(errAppointmentDateOverlap)
		}
	}

	return nil // No conflicts found
}
