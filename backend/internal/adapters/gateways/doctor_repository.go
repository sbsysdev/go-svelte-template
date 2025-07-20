package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/domain"
)

type doctorRepository struct {
	pool                *pgxpool.Pool
	specialtyRepository domain.SpecialtyRepository
}

func (repository *doctorRepository) Save(ctx context.Context, doctor *domain.Doctor) error {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, execErr := tx.Exec(ctx, createDoctorSQL, doctor.ID, doctor.Name)
	if execErr != nil {
		return execErr
	}

	for _, specialty := range doctor.Specialties {
		_, execErr = tx.Exec(ctx, createDoctorSpecialtySQL, doctor.ID, specialty.ID)
		if execErr != nil {
			return execErr
		}
	}

	return tx.Commit(ctx)
}
func (repository *doctorRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error) {
	row := repository.pool.QueryRow(ctx, findDoctorByIDSQL, id)

	var doctor domain.Doctor
	if err := row.Scan(&doctor.ID, &doctor.Name); err != nil {
		return nil, err
	}

	specialties, err := repository.specialtyRepository.FindByDoctorID(ctx, id)
	if err != nil {
		return nil, err
	}
	doctor.Specialties = specialties

	return &doctor, nil
}
func (repository *doctorRepository) FindBySpecialtyID(ctx context.Context, specialtyID uuid.UUID) ([]*domain.Doctor, error) {
	rows, err := repository.pool.Query(ctx, findDoctorsBySpecialtyIDSQL, specialtyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []*domain.Doctor
	for rows.Next() {
		var doctor domain.Doctor
		if err := rows.Scan(&doctor.ID, &doctor.Name); err != nil {
			return nil, err
		}

		specialties, err := repository.specialtyRepository.FindByDoctorID(ctx, doctor.ID)
		if err != nil {
			return nil, err
		}
		doctor.Specialties = specialties

		doctors = append(doctors, &doctor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}

func NewDoctorRepository(
	pool *pgxpool.Pool,
	specialtyRepository domain.SpecialtyRepository,
) domain.DoctorRepository {
	return &doctorRepository{
		pool:                pool,
		specialtyRepository: specialtyRepository,
	}
}
