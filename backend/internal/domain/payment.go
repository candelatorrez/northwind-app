package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID uint `gorm:"primaryKey"`

	InvoiceID uint `gorm:"index;not null"`

	Amount float64 `gorm:"not null"`

	PaidAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
