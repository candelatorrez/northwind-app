package repository

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) Create(client *domain.Client) error {
	return r.db.Create(client).Error
}

func (r *ClientRepository) FindByID(id uint) (*domain.Client, error) {
	var client domain.Client

	err := r.db.First(&client, id).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) FindAll() ([]domain.Client, error) {
	var clients []domain.Client

	err := r.db.Order("id ASC").Find(&clients).Error

	return clients, err
}

func (r *ClientRepository) Update(client *domain.Client) error {
	return r.db.Save(client).Error
}
