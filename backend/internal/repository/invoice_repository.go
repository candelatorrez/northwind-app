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

func (r *InvoiceRepository) FindByID(id uint) (*domain.Invoice, error) {
	var invoice domain.Invoice

	err := r.db.First(&invoice, id).Error
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (r *InvoiceRepository) Update(invoice *domain.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *InvoiceRepository) FindClientByID(clientID uint) ([]domain.Invoice, error) {
	var invoices []domain.Invoice

	err := r.db.Where("client_id = ?", clientID).Find(&invoices).Error

	return invoices, err
}

func (r *InvoiceRepository) CountOverdue() (int64, error) {
	var count int64

	err := r.db.Model(&domain.Invoice{}).Where("status = ?", domain.InvoiceOverdue).Count(&count).Error

	return count, err
}

func (r *InvoiceRepository) SumOutstandingAmount() (float64, error) {
	var total float64

	err := r.db.Model(&domain.Invoice{}).Where("status IN ?", []string{
		string(domain.InvoicePending),
		string(domain.InvoiceOverdue),
	}).Select("COALESCE(SUM(amount), 0)").Scan(&total).Error

	return total, err
}
