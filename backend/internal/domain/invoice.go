package domain

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceStatus string

const (
	InvoicePending InvoiceStatus = "pending"
	InvoicePaid    InvoiceStatus = "paid"
	InvoiceOverdue InvoiceStatus = "overdue"
)

type Invoice struct {
	ID uint `gorm:"primaryKey"`

	ClientID uint `gorm:"index;not null"`

	Amount  float64 `gorm:"not null"`
	DueDate time.Time

	Status InvoiceStatus `gorm:"type:varchar(20);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
