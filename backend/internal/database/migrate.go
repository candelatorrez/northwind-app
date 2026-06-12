package database

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Client{},
		&domain.Invoice{},
		&domain.Payment{},
		&domain.CollectionAction{},
		&domain.RiskSnapshot{},
	)
}
