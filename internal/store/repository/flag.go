package repository

import (
	"context"
	"flagd/internal/domain"
)

type FlagRepository interface {
	GetById(ctx context.Context, id string) (*domain.Flag, error)
	Create(ctx context.Context, key string, name string, description string) (*domain.Flag, error)
	Update(ctx context.Context, flag *domain.Flag) error
	Delete(ctx context.Context, id string) error
}
