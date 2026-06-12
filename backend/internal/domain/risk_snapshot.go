package domain

import (
	"time"

	"gorm.io/gorm"
)

type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

type RiskSnapshot struct {
	ID uint `gorm:"primaryKey"`

	ClientID uint `gorm:"index;not null"`

	Score int `gorm:"not null"`

	Level RiskLevel `gorm:"type:varchar(20);not null"`

	Reason string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
