package domain

import (
	"context"

	"github.com/google/uuid"
)

type Specialty struct {
	ID       uuid.UUID
	Name     string
	Duration int // in minutes
}

func NewSpecialty(name string, duration int) *Specialty {
	return &Specialty{
		ID:       uuid.New(),
		Name:     name,
		Duration: duration,
	}
}

type SpecialtyRepository interface {
	Save(context.Context, *Specialty) error
	FindByID(context.Context, uuid.UUID) (*Specialty, error)
	FindByDoctorID(context.Context, uuid.UUID) ([]*Specialty, error)
	FindAll(context.Context) ([]*Specialty, error)
}
