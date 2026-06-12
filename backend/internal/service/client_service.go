package service

import (
	"errors"

	"github.com/candelatorrez/northwind-app/internal/domain"
	"github.com/candelatorrez/northwind-app/internal/repository"
	"gorm.io/gorm"
)

type ClientService struct {
	clientRepo *repository.ClientRepository
}

func NewClientService(clientRepo *repository.ClientRepository) *ClientService {
	return &ClientService{
		clientRepo: clientRepo,
	}
}

func (s *ClientService) GetAllClients() ([]domain.Client, error) {
	return s.clientRepo.FindAll()
}

func (s *ClientService) GetClientByID(id uint) (*domain.Client, error) {
	return s.clientRepo.FindByID(id)
}

func (s *ClientService) UpdateClient(client *domain.Client) error {
	return s.clientRepo.Update(client)
}

func (s *ClientService) UpdateClientStatus(id uint, status domain.ClientStatus) (*domain.Client, error) {
	if !isValidClientStatus(status) {
		return nil, ErrInvalidClientStatus
	}

	client, err := s.clientRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	client.Status = status

	if err := s.clientRepo.Update(client); err != nil {
		return nil, err
	}

	return client, nil
}

func isValidClientStatus(status domain.ClientStatus) bool {
	switch status {
	case domain.StatusActive, domain.StatusAtRisk, domain.StatusDelinquent, domain.StatusSuspended:
		return true
	default:
		return false
	}
}
