package gateways

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type specialityRepository struct {
	pool *pgxpool.Pool
}

func (r *specialityRepository) Save(ctx context.Context, specialty *domain.Speciality) error {
	// Implementation for saving a speciality
	return nil
}
func (r *specialityRepository) FindAll(ctx context.Context) ([]*domain.Speciality, error) {
	// Implementation for finding all specialities
	return []*domain.Speciality{}, nil
}

func NewSpecialityRepository(pool *pgxpool.Pool) domain.SpecialityRepository {
	return &specialityRepository{
		pool: pool,
	}
}
