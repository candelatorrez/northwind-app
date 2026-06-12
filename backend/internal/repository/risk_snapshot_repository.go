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
	if err != nil {
		return nil, err
	}

	return &risk, nil
}

func (r *RiskSnapshotRepository) Count() (int64, error) {
	var count int64

	err := r.db.Model(&domain.RiskSnapshot{}).Count(&count).Error

	return count, err
}

func (r *RiskSnapshotRepository) CountHighRisk() (int64, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(*) FROM (
			SELECT DISTINCT ON (client_id) level
			FROM risk_snapshots
			WHERE deleted_at IS NULL
			ORDER BY client_id, created_at DESC
		) latest
		WHERE level = ?
	`, domain.RiskHigh).Scan(&count).Error

	return count, err
}
