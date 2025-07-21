package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

const errDoctorDoesNotHoldSpecialty = "doctor does not hold the specialty"

type Doctor struct {
	ID          uuid.UUID
	Name        string
	Specialties []*Specialty
}

func (d *Doctor) HasSpecialty(specialty *Specialty) error {
	for _, s := range d.Specialties {
		if s.ID == specialty.ID {
			return nil
		}
	}
	return errors.New(errDoctorDoesNotHoldSpecialty)
}

func NewDoctor(name string, specialties []*Specialty) *Doctor {
	return &Doctor{
		ID:          uuid.New(),
		Name:        name,
		Specialties: specialties,
	}
}

type DoctorRepository interface {
	Save(context.Context, *Doctor) error
	FindByID(context.Context, uuid.UUID) (*Doctor, error)
	FindBySpecialtyID(context.Context, uuid.UUID) ([]*Doctor, error)
}
