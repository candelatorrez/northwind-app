package service

import "github.com/candelatorrez/northwind-app/internal/domain"

type RiskService struct{}

func NewRiskService() *RiskService {
	return &RiskService{}
}

func (s *RiskService) CalculateRiskSLevel(daysOverdue int) domain.RiskLevel {
	switch {
	case daysOverdue >= 60:
		return domain.RiskHigh
	case daysOverdue >= 30:
		return domain.RiskMedium
	default:
		return domain.RiskLow
	}
}

func (s *RiskService) CalculateScore(daysOverdue int) int {
	switch {
	case daysOverdue >= 60:
		return 90
	case daysOverdue >= 30:
		return 60
	default:
		return 20
	}
}
