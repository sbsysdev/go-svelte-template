package gateways

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type specialityRepository struct {
	pool *pgxpool.Pool
}

func (repository *specialityRepository) Save(ctx context.Context, specialty *domain.Speciality) error {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, execErr := tx.Exec(ctx, createSpecialitySQL, specialty.ID, specialty.Name, specialty.Duration)
	if execErr != nil {
		return execErr
	}

	return tx.Commit(ctx)
}
func (repository *specialityRepository) FindAll(ctx context.Context) ([]*domain.Speciality, error) {
	rows, err := repository.pool.Query(ctx, findAllSpecialitiesSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialities []*domain.Speciality
	for rows.Next() {
		var speciality domain.Speciality
		if err := rows.Scan(&speciality.ID, &speciality.Name, &speciality.Duration); err != nil {
			return nil, err
		}
		specialities = append(specialities, &speciality)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return specialities, nil
}

func NewSpecialityRepository(pool *pgxpool.Pool) domain.SpecialityRepository {
	return &specialityRepository{
		pool: pool,
	}
}
