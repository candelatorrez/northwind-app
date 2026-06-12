package repository

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) Create(payment *domain.Payment) error {
	return r.db.Create(payment).Error
}
