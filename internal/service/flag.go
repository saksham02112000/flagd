package service

import (
	"context"
	"flagd/internal/domain"
	"flagd/internal/logger"
	"flagd/internal/store/repository"
)

type FlagService struct {
	flagRepository repository.FlagRepository
}

func NewFlagService(flagRepository repository.FlagRepository) *FlagService {
	return &FlagService{flagRepository: flagRepository}
}

func (s *FlagService) GetById(ctx context.Context, id string) (*domain.Flag, error) {
	return s.flagRepository.GetById(ctx, id)
}

func (s *FlagService) CreateFlag(ctx context.Context, key string, name string, description string) (*domain.Flag, error) {
	log := logger.FromContext(ctx)
	log.InfoContext(ctx, "Creating flag", "key", key, "name", name, "description", description)

	flag, err := s.flagRepository.Create(ctx, key, name, description)
	if err != nil {
		log.Error("failed to create flag", "key", key, "error", err)
		return nil, err
	}

	log.Info("flag created", "id", flag.ID, "key", flag.Key)
	return flag, nil
}
