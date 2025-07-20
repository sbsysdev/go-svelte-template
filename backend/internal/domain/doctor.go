package domain

import "github.com/google/uuid"

type Doctor struct {
	ID          uuid.UUID
	Name        string
	Specialties []*Speciality
}

func NewDoctor(name string, specialties []*Speciality) *Doctor {
	return &Doctor{
		ID:          uuid.New(),
		Name:        name,
		Specialties: specialties,
	}
}

type DoctorRepository interface {
	Save(doctor *Doctor) error
	FindByID(id uuid.UUID) (*Doctor, error)
	FindBySpecialty(specialtyID uuid.UUID) ([]*Doctor, error)
}
