package domain

import (
	"context"

	"github.com/google/uuid"
)

type Speciality struct {
	ID       uuid.UUID
	Name     string
	Duration int // in minutes
}

func NewSpeciality(name string, duration int) *Speciality {
	return &Speciality{
		ID:       uuid.New(),
		Name:     name,
		Duration: duration,
	}
}

type SpecialityRepository interface {
	Save(context.Context, *Speciality) error
	FindAll(context.Context) ([]*Speciality, error)
}
