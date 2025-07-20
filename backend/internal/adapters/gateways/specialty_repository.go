package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type specialtyRepository struct {
	pool *pgxpool.Pool
}

func (repository *specialtyRepository) Save(ctx context.Context, specialty *domain.Specialty) error {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, execErr := tx.Exec(ctx, createSpecialtySQL, specialty.ID, specialty.Name, specialty.Duration)
	if execErr != nil {
		return execErr
	}

	return tx.Commit(ctx)
}
func (repository *specialtyRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Specialty, error) {
	row := repository.pool.QueryRow(ctx, findSpecialtyByIDSQL, id)

	var specialty domain.Specialty
	if err := row.Scan(&specialty.ID, &specialty.Name, &specialty.Duration); err != nil {
		return nil, err
	}

	return &specialty, nil
}
func (repository *specialtyRepository) FindByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]*domain.Specialty, error) {
	rows, err := repository.pool.Query(ctx, findSpecialtiesByDoctorIDSQL, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialties []*domain.Specialty
	for rows.Next() {
		var specialty domain.Specialty
		if err := rows.Scan(&specialty.ID, &specialty.Name, &specialty.Duration); err != nil {
			return nil, err
		}
		specialties = append(specialties, &specialty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return specialties, nil
}
func (repository *specialtyRepository) FindAll(ctx context.Context) ([]*domain.Specialty, error) {
	rows, err := repository.pool.Query(ctx, findAllSpecialtiesSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialties []*domain.Specialty
	for rows.Next() {
		var specialty domain.Specialty
		if err := rows.Scan(&specialty.ID, &specialty.Name, &specialty.Duration); err != nil {
			return nil, err
		}
		specialties = append(specialties, &specialty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return specialties, nil
}

func NewSpecialtyRepository(pool *pgxpool.Pool) domain.SpecialtyRepository {
	return &specialtyRepository{
		pool: pool,
	}
}
