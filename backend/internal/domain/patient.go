package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID    uuid.UUID
	Name  string
	Birth time.Time
}

func NewPatient(name string, birth time.Time) *Patient {
	return &Patient{
		ID:    uuid.New(),
		Name:  name,
		Birth: birth,
	}
}

type PatientRepository interface {
	Save(context.Context, *Patient) error
	FindAll(context.Context) ([]*Patient, error)
	FindByID(context.Context, uuid.UUID) (*Patient, error)
}
