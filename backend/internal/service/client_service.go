package service

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"github.com/candelatorrez/northwind-app/internal/repository"
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
