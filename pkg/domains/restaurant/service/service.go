package service

import (
	"fmt"

	"my-restaurant/pkg/domains/restaurant/model"
	"my-restaurant/pkg/domains/restaurant/repository"
)

// ServiceI is a interface to communicate with Repository
type ServiceI interface {
	LoadRestaurant() (model.Restaurant, error)
	SetupRestaurant() (model.Restaurant, error)
	TakeOrder(model.Restaurant) ([]model.Table, error)
	PrepareOrder(model.Table, int64) error
}

type Service struct {
	repository repository.RepositoryI
}

// NewService generates a struct of type Service
func NewService(repository repository.RepositoryI) (*Service, error) {
	if repository == nil {
		return nil, fmt.Errorf("repository is nil")
	}
	return &Service{
		repository: repository,
	}, nil
}

func (s *Service) LoadRestaurant() (model.Restaurant, error) {
	return s.repository.LoadRestaurant()
}

func (s *Service) SetupRestaurant() (model.Restaurant, error) {
	return s.repository.SetupRestaurant()
}

func (s *Service) TakeOrder(rest model.Restaurant) ([]model.Table, error) {
	return s.repository.TakeOrder(rest)
}

func (s *Service) PrepareOrder(table model.Table, chefs int64) error {
	return s.repository.PrepareOrder(table, chefs)
}
