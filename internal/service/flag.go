package service

import (
	"flagd/internal/domain"
	"flagd/internal/store/repository"
)


type FlagService struct{
	flagRepository repository.FlagRepository
}


func NewFlagService(flagRepository repository.FlagRepository) *FlagService{
	return &FlagService{flagRepository: flagRepository}
}

func (s *FlagService) GetById(id string) (*domain.Flag, error){
	return s.flagRepository.GetById(id)
}

func (s *FlagService) Create(flag *domain.Flag) error{
	return s.flagRepository.Create(flag)
}

func (s *FlagService) Update(flag *domain.Flag) error{
	return s.flagRepository.Update(flag)
}

func (s *FlagService) Delete(id string) error{
	return s.flagRepository.Delete(id)
}
