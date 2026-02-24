package postgres

import (
	"database/sql"
	"flagd/internal/domain"
)

type PostgresFlagRepository struct {
	db *sql.DB
}

func (r *PostgresFlagRepository) GetById(id string) (*domain.Flag, error) {
	return nil, nil
}

func (r *PostgresFlagRepository) Create(flag *domain.Flag) error {
	return nil
}

func (r *PostgresFlagRepository) Update(flag *domain.Flag) error {
	return nil
}

func (r *PostgresFlagRepository) Delete(id string) error {
	return nil
}
