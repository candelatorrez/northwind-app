package domain

import (
	"time"

	"gorm.io/gorm"
)

type ActionType string

const (
	ActionCall  ActionType = "call"
	ActionEmail ActionType = "email"
	ActionNote  ActionType = "note"
)

type CollectionAction struct {
	ID uint `gorm:"primaryKey"`

	ClientID uint `gorm:"index;not null"`

	Type ActionType `gorm:"type:varchar(20);not null"`

	Notes string `gorm:"type:text"`

	PerformedBy string `gorm:"size:100"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
