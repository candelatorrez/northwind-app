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

type Client struct {
	ID             uint          `gorm:"primaryKey"`
	Name           string        `gorm:"not null"`
	Email          string        `gorm:"not null"`
	Segment        ClientSegment `gorm:"type:varchar(20);notnull"`
	MonthlyBilling float64       `gorm:"not null"`

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
