package repository

import (
	"flagd/internal/domain"
	"flagd/internal/store/postgres"
)

type FlagRepository interface {
	GetById(id string) (*domain.Flag, error)
	Create(flag *domain.Flag) error
	Update(flag *domain.Flag) error
	Delete(id string) error
}

func NewFlagRepository() *postgres.PostgresFlagRepository {
	return &postgres.PostgresFlagRepository{}
}
