package domain

import (
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
	Save(patient *Patient) error
	FindByID(id uuid.UUID) (*Patient, error)
	FindByName(name string) ([]*Patient, error)
	FindAll() ([]*Patient, error)
}
