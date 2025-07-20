package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type patientRepository struct {
	pool *pgxpool.Pool
}

func (repository *patientRepository) Save(ctx context.Context, patient *domain.Patient) error {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, execErr := tx.Exec(ctx, createPatientSQL, patient.ID, patient.Name, patient.Birth)
	if execErr != nil {
		return execErr
	}

	return tx.Commit(ctx)
}
func (repository *patientRepository) FindAll(ctx context.Context) ([]*domain.Patient, error) {
	rows, err := repository.pool.Query(ctx, findAllPatientsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []*domain.Patient
	for rows.Next() {
		var patient domain.Patient
		if err := rows.Scan(&patient.ID, &patient.Name, &patient.Birth); err != nil {
			return nil, err
		}
		patients = append(patients, &patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}
func (repository *patientRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Patient, error) {
	row := repository.pool.QueryRow(ctx, findPatientByIDSQL, id)

	var patient domain.Patient
	if err := row.Scan(&patient.ID, &patient.Name, &patient.Birth); err != nil {
		return nil, err
	}

	return &patient, nil
}

func NewPatientRepository(pool *pgxpool.Pool) domain.PatientRepository {
	return &patientRepository{pool: pool}
}
