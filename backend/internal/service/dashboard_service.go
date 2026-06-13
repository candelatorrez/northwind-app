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

func (s *DashboardService) GetClientsOverview() ([]domain.ClientOverview, error) {

	clients, err := s.clientRepo.FindAll()

	if err != nil {
		return nil, err
	}

	result := make([]domain.ClientOverview, 0)

	for _, client := range clients {

		snapshot, err := s.riskRepo.FindLatestByClientID(client.ID)

		if err != nil {
			continue
		}

		result = append(result, domain.ClientOverview{
			ID:             client.ID,
			Name:           client.Name,
			Segment:        string(client.Segment),
			Status:         string(client.Status),
			MonthlyBilling: client.MonthlyBilling,
			RiskScore:      snapshot.Score,
			RiskLevel:      string(snapshot.Level),
		})
	}

	return result, nil
}
