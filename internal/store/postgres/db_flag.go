package postgres

import (
	"context"
	"errors"
	"flagd/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresFlagRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFlagRepository(db *pgxpool.Pool) *PostgresFlagRepository {
	return &PostgresFlagRepository{db: db}
}

func (r *PostgresFlagRepository) GetAll(ctx context.Context) ([]*domain.Flag, error) {
	query := `SELECT id, key, name, description, created_at, updated_at, archived_at 
			  FROM flags ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flags []*domain.Flag
	for rows.Next() {
		var flag domain.Flag
		err := rows.Scan(
			&flag.ID, &flag.Key, &flag.Name, &flag.Description,
			&flag.CreatedAt, &flag.UpdatedAt, &flag.ArchivedAt,
		)
		if err != nil {
			return nil, err
		}
		flags = append(flags, &flag)
	}
	return flags, nil
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrFlagNotFound
		}
		return nil, err
	}
	return &flag, nil
}

func (r *PostgresFlagRepository) Create(ctx context.Context, key string, name string, description string) (*domain.Flag, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// 1. Insert the flag
	flagQuery := `INSERT INTO flags (key, name, description) 
				  VALUES ($1, $2, $3) 
				  RETURNING id, created_at, updated_at`

	flag := &domain.Flag{
		Key:         key,
		Name:        name,
		Description: description,
	}

	err = tx.QueryRow(ctx, flagQuery, flag.Key, flag.Name, flag.Description).Scan(
		&flag.ID, &flag.CreatedAt, &flag.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// 2. Seed flag_environments for every existing environment (all disabled)
	seedQuery := `
		INSERT INTO flag_environments (flag_id, environment_id, enabled)
		SELECT $1, id, false
		FROM environments
	`
	_, err = tx.Exec(ctx, seedQuery, flag.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
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
	query := `UPDATE flags SET archived_at = NOW() WHERE id = $1 AND archived_at IS NULL`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrFlagNotFound
	}
	return nil
}

func (r *PostgresFlagRepository) Toggle(ctx context.Context, id string, envSlug string) (*domain.FlagEnvironment, error) {
	query := `
		UPDATE flag_environments fe
		SET enabled = NOT enabled, updated_at = NOW()
		FROM environments e
		WHERE fe.environment_id = e.id
		  AND fe.flag_id = $1
		  AND e.slug = $2
		RETURNING fe.flag_id, fe.environment_id, e.slug, fe.enabled, fe.updated_at
	`

	var fe domain.FlagEnvironment
	err := r.db.QueryRow(ctx, query, id, envSlug).Scan(
		&fe.FlagID, &fe.EnvironmentID, &fe.EnvironmentSlug, &fe.Enabled, &fe.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrEnvNotFound
		}
		return nil, err
	}
	return &fe, nil
}
