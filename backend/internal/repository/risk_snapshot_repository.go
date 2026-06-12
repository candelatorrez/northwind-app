package repository

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

type RiskSnapshotRepository struct {
	db *gorm.DB
}

func NewRiskSnapshotRepository(db *gorm.DB) *RiskSnapshotRepository {
	return &RiskSnapshotRepository{
		db: db,
	}
}

func (r *RiskSnapshotRepository) Create(risk *domain.RiskSnapshot) error {
	return r.db.Create(risk).Error
}

func (r *RiskSnapshotRepository) FindLatestByClientID(clientID uint) (*domain.RiskSnapshot, error) {
	var risk domain.RiskSnapshot

	err := r.db.Where("client_id = ?", clientID).Order("created_at DESC").First(&risk).Error

	return &risk, err
}
