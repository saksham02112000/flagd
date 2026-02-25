package postgres

import (
	"context"
	"flagd/internal/domain"

	"github.com/jackc/pgx/v5"
)

type PostgresFlagRepository struct {
	db *pgx.Conn
}

func NewPostgresFlagRepository(db *pgx.Conn) *PostgresFlagRepository {
	return &PostgresFlagRepository{db: db}
}

func (r *PostgresFlagRepository) GetById(ctx context.Context, id string) (*domain.Flag, error) {
	query := `SELECT id, key, name, description, created_at, updated_at, archived_at 
			  FROM flags WHERE id = $1`

	var flag domain.Flag
	err := r.db.QueryRow(ctx, query, id).Scan(
		&flag.ID, &flag.Key, &flag.Name, &flag.Description,
		&flag.CreatedAt, &flag.UpdatedAt, &flag.ArchivedAt,
	)
	if err != nil {
		return nil, err
	}
	return &flag, nil
}

func (r *PostgresFlagRepository) Create(ctx context.Context, key string, name string, description string) (*domain.Flag, error) {
	query := `INSERT INTO flags (key, name, description) 
			  VALUES ($1, $2, $3) 
			  RETURNING id, created_at, updated_at`

	flag := &domain.Flag{
		Key:         key,
		Name:        name,
		Description: description,
	}

	err := r.db.QueryRow(ctx, query, flag.Key, flag.Name, flag.Description).Scan(
		&flag.ID, &flag.CreatedAt, &flag.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return flag, nil
}

func (r *PostgresFlagRepository) Update(ctx context.Context, flag *domain.Flag) error {
	query := `UPDATE flags SET key = $1, name = $2, description = $3, updated_at = NOW() 
			  WHERE id = $4`
	_, err := r.db.Exec(ctx, query, flag.Key, flag.Name, flag.Description, flag.ID)
	return err
}

func (r *PostgresFlagRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM flags WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
