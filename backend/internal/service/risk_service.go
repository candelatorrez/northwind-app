package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/candelatorrez/northwind-app/internal/domain"
	"github.com/candelatorrez/northwind-app/internal/repository"
	"gorm.io/gorm"
)

var ErrRiskSnapshotNotFound = errors.New("risk snapshot not found")

type RiskService struct {
	riskRepo    *repository.RiskSnapshotRepository
	invoiceRepo *repository.InvoiceRepository
	clientRepo  *repository.ClientRepository
}

func NewRiskService(
	riskRepo *repository.RiskSnapshotRepository,
	invoiceRepo *repository.InvoiceRepository,
	clientRepo *repository.ClientRepository,
) *RiskService {
	return &RiskService{
		riskRepo:    riskRepo,
		invoiceRepo: invoiceRepo,
		clientRepo:  clientRepo,
	}
}

func (s *RiskService) CalculateRiskLevel(daysOverdue int) domain.RiskLevel {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *RiskService) CalculateScoreWithSegment(daysOverdue int, segment domain.ClientSegment) int {
	base := s.CalculateScore(daysOverdue)

	switch segment {
	case domain.SegmentEnterprise:
		return max(0, base-20)
	case domain.SegmentStartup:
		return min(100, base+20)
	case domain.SegmentZombie:
		if daysOverdue >= 30 {
			return 95
		}
		return min(100, base+30)
	default:
		return base
	}
}

func (s *RiskService) CalculateDaysOverdue(invoices []domain.Invoice) int {
	now := time.Now()
	maxDays := 0

	for _, invoice := range invoices {
		if invoice.Status == domain.InvoicePaid {
			continue
		}

		if invoice.Status != domain.InvoiceOverdue && !invoice.DueDate.Before(now) {
			continue
		}

		days := int(now.Sub(invoice.DueDate).Hours() / 24)
		if days > maxDays {
			maxDays = days
		}
	}

	return maxDays
}

func (s *RiskService) buildSnapshot(clientID uint, daysOverdue int) *domain.RiskSnapshot {
	level := s.CalculateRiskLevel(daysOverdue)
	reason := "no overdue invoices"

	if daysOverdue > 0 {
		reason = fmt.Sprintf("%d days overdue on oldest unpaid invoice", daysOverdue)
	}

	return &domain.RiskSnapshot{
		ClientID: clientID,
		Score:    s.CalculateScore(daysOverdue),
		Level:    level,
		Reason:   reason,
	}
}

func (s *RiskService) SnapshotClient(clientID uint) (*domain.RiskSnapshot, error) {
	client, err := s.clientRepo.FindByID(clientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	invoices, err := s.invoiceRepo.FindClientByID(clientID)
	if err != nil {
		return nil, err
	}

	days := s.CalculateDaysOverdue(invoices)

	snapshot := s.buildSnapshot(clientID, days)
	// adjust score by client segment
	snapshot.Score = s.CalculateScoreWithSegment(days, client.Segment)

	if err := s.riskRepo.Create(snapshot); err != nil {
		return nil, err
	}

	return snapshot, nil
}

func (s *RiskService) SnapshotAllClients() error {
	clients, err := s.clientRepo.FindAll()
	if err != nil {
		return err
	}

	for _, client := range clients {
		if _, err := s.SnapshotClient(client.ID); err != nil {
			return err
		}
	}

	return nil
}

func (s *RiskService) GetLatestSnapshot(clientID uint) (*domain.RiskSnapshot, error) {
	snapshot, err := s.riskRepo.FindLatestByClientID(clientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRiskSnapshotNotFound
		}
		return nil, err
	}

	return snapshot, nil
}

func (s *RiskService) EnsureSnapshots() error {
	count, err := s.riskRepo.Count()
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	return s.SnapshotAllClients()
}
