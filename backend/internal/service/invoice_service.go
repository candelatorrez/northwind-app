package service

import (
	"errors"
	"time"

	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

type InvoiceService struct {
	db *gorm.DB
}

func NewInvoiceService(db *gorm.DB) *InvoiceService {
	return &InvoiceService{
		db: db,
	}
}

func (s *InvoiceService) MarkAsPaid(invoiceID uint) (*domain.Invoice, error) {
	var updatedInvoice domain.Invoice

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var invoice domain.Invoice
		if err := tx.First(&invoice, invoiceID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrInvoiceNotFound
			}
			return err
		}

		if invoice.Status == domain.InvoicePaid {
			return ErrInvoiceAlreadyPaid
		}

		payment := domain.Payment{
			InvoiceID: invoice.ID,
			Amount:    invoice.Amount,
			PaidAt:    time.Now(),
		}

		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		invoice.Status = domain.InvoicePaid
		if err := tx.Save(&invoice).Error; err != nil {
			return err
		}

		updatedInvoice = invoice
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &updatedInvoice, nil
}
