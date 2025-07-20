package domain

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentState string

const (
	AppointmentStateScheduled AppointmentState = "scheduled"
	AppointmentStateCompleted AppointmentState = "completed"
	AppointmentStateCancelled AppointmentState = "cancelled"
)

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
	Save(appointment *Appointment) error
	FindByID(id uuid.UUID) (*Appointment, error)
	FindByDoctor(doctorID uuid.UUID, startDate time.Time, endDate time.Time) ([]*Appointment, error)
	FindByPatient(patientID uuid.UUID, startDate time.Time, endDate time.Time) ([]*Appointment, error)
}

type AppointmentGuard struct {
	appointmentRepository AppointmentRepository
}

func NewAppointmentGuard(appointmentRepository AppointmentRepository) *AppointmentGuard {
	return &AppointmentGuard{
		appointmentRepository: appointmentRepository,
	}
}

func (guard *AppointmentGuard) CheckAvailabilityByDate(appointment *Appointment, startDate time.Time, endDate time.Time) bool {
	// Check if the doctor is available at the requested time
	doctorAppointments, err := guard.appointmentRepository.FindByDoctor(appointment.Doctor.ID, startDate, endDate)
	if err != nil {
		return false
	}
	if !guard.checkAvailabilityByAppointments(doctorAppointments, appointment) {
		return false
	}

	// Check if the patient has any conflicting appointments
	patientAppointments, err := guard.appointmentRepository.FindByPatient(appointment.Patient.ID, startDate, endDate)
	if err != nil {
		return false
	}
	if !guard.checkAvailabilityByAppointments(patientAppointments, appointment) {
		return false
	}

	// If both checks pass, the appointment can be scheduled
	return true
}

func (guard *AppointmentGuard) checkAvailabilityByAppointments(appointments []*Appointment, appointment *Appointment) bool {
	for _, existingAppointment := range appointments {
		if existingAppointment.State != AppointmentStateScheduled {
			continue
		}

		existingStartDate := existingAppointment.Date
		existingEndDate := existingStartDate.Add(time.Duration(appointment.Specialty.Duration) * time.Minute)

		requestedStartDate := appointment.Date
		requestedEndDate := requestedStartDate.Add(time.Duration(appointment.Specialty.Duration) * time.Minute)

		if requestedStartDate.After(existingStartDate) && requestedStartDate.Before(existingEndDate) {
			return false // Requested time overlaps with an existing appointment
		}
		if requestedEndDate.After(existingStartDate) && requestedEndDate.Before(existingEndDate) {
			return false // Requested time overlaps with an existing appointment
		}
	}

	return true // No conflicts found
}
