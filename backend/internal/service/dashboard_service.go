package service

import (
	"github.com/candelatorrez/northwind-app/internal/domain"
	"github.com/candelatorrez/northwind-app/internal/repository"
)

type DashboardService struct {
	clientRepo  *repository.ClientRepository
	invoiceRepo *repository.InvoiceRepository
	riskRepo    *repository.RiskSnapshotRepository
}

func NewDashboardService(
	clientRepo *repository.ClientRepository,
	invoiceRepo *repository.InvoiceRepository,
	riskRepo *repository.RiskSnapshotRepository,
) *DashboardService {
	return &DashboardService{
		clientRepo:  clientRepo,
		invoiceRepo: invoiceRepo,
		riskRepo:    riskRepo,
	}
}

func (s *DashboardService) GetMetrics() (*domain.DashboardMetrics, error) {
	clients, err := s.clientRepo.FindAll()
	if err != nil {
		return nil, err
	}

	highRisk, err := s.riskRepo.CountHighRisk()
	if err != nil {
		return nil, err
	}

	overdue, err := s.invoiceRepo.CountOverdue()
	if err != nil {
		return nil, err
	}

	outstanding, err := s.invoiceRepo.SumOutstandingAmount()
	if err != nil {
		return nil, err
	}

	return &domain.DashboardMetrics{
		TotalClients:      len(clients),
		HighRiskClients:   int(highRisk),
		OverdueInvoices:   int(overdue),
		OutstandingAmount: outstanding,
	}, nil
}
