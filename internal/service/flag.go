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

func (s *FlagService) GetAll(ctx context.Context) ([]*domain.Flag, error) {
	return s.flagRepository.GetAll(ctx)
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

func (s *FlagService) DeleteFlag(ctx context.Context, id string) error {
	log := logger.FromContext(ctx)
	log.InfoContext(ctx, "Deleting flag", "id", id)

	err := s.flagRepository.Delete(ctx, id)
	if err != nil {
		log.Error("failed to delete flag", "id", id, "error", err)
		return err
	}

	log.Info("flag deleted", "id", id)
	return nil
}

func (s *FlagService) ToggleFlag(ctx context.Context, id string, envSlug string) (*domain.FlagEnvironment, error) {
	log := logger.FromContext(ctx)
	log.InfoContext(ctx, "Toggling flag", "id", id, "envSlug", envSlug)

	fe, err := s.flagRepository.Toggle(ctx, id, envSlug)
	if err != nil {
		log.Error("failed to toggle flag", "id", id, "envSlug", envSlug, "error", err)
		return nil, err
	}

	log.Info("flag toggled", "id", id, "envSlug", envSlug, "enabled", fe.Enabled)
	return fe, nil
}
