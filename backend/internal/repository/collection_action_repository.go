package repository

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

type CollectionActionRepository struct {
	db *gorm.DB
}

func NewCollectionActionRepository(db *gorm.DB) *CollectionActionRepository {
	return &CollectionActionRepository{
		db: db,
	}
}

func (r *CollectionActionRepository) Create(action *domain.CollectionAction) error {
	return r.db.Create(action).Error
}

func (r *CollectionActionRepository) FindByClientID(clientID uint) ([]domain.CollectionAction, error) {
	var actions []domain.CollectionAction

	err := r.db.Where("client_id = ?", clientID).Order("created_at DESC").Find(&actions).Error

	return actions, err
}
