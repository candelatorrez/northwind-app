package domain

import (
	"time"

	"gorm.io/gorm"
)

type ClientSegment string

const (
	SegmentEnterprise ClientSegment = "enterprise"
	SegmentStartup    ClientSegment = "startup"
	SegmentStandard   ClientSegment = "standard"
	SegmentZombie     ClientSegment = "zombie"
)

type ClientStatus string

const (
	StatusActive     ClientStatus = "active"
	StatusAtRisk     ClientStatus = "at_risk"
	StatusDelinquent ClientStatus = "delinquent"
	StatusSuspended  ClientStatus = "suspended"
)

type Client struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`

	Segment ClientSegment `gorm:"type:varchar(20);not null" json:"segment"`
	Status  ClientStatus  `gorm:"type:varchar(20);not null" json:"status"`

	MonthlyBilling float64 `gorm:"not null" json:"monthlyBilling"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
