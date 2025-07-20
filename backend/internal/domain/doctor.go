package domain

import (
	"context"

	"github.com/google/uuid"
)

type Doctor struct {
	ID          uuid.UUID
	Name        string
	Specialties []*Specialty
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
