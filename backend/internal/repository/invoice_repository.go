package repository

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{
		db: db,
	}
}

func (r *InvoiceRepository) Create(invoice *domain.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *InvoiceRepository) FindClientByID(clientID uint) ([]domain.Invoice, error) {
	var invoices []domain.Invoice

	err := r.db.Where("client_id = ?", clientID).Find(&invoices).Error

	return invoices, err
}
