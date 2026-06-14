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
	ID uint `gorm:"primaryKey" json:"id"`

	ClientID uint `gorm:"index;not null" json:"clientId"`

	Type ActionType `gorm:"type:varchar(20);not null" json:"type"`

	Notes string `gorm:"type:text" json:"notes"`

	PerformedBy string `gorm:"size:100;not null" json:"performedBy"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deteledAt"`
}
