package seed

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/candelatorrez/northwind-app/internal/domain"
	"gorm.io/gorm"
)

var segments = []domain.ClientSegment{
	domain.SegmentEnterprise,
	domain.SegmentStartup,
	domain.SegmentStandard,
	domain.SegmentZombie,
}

var status = []domain.ClientStatus{
	domain.StatusActive,
	domain.StatusAtRisk,
	domain.StatusDelinquent,
	domain.StatusSuspended,
}

func Run(db *gorm.DB) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	clients := make([]domain.Client, 0, 420)

	for i := 1; i <= 420; i++ {
		segment := randomSegment(r)
		status := randomStatus(segment, r)

		client := domain.Client{
			Name:  fmt.Sprintf("client %d", i),
			Email: fmt.Sprintf("client%d@northwind.com", i),

			Segment: segment,
			Status:  status,

			MonthlyBilling: randomBilling(segment, r),
		}
		clients = append(clients, client)
	}

	if err := db.Create(&clients).Error; err != nil {
		return err
	}

	for _, c := range clients {
		invoices := generateInvoices(c.ID, r)
		if len(invoices) > 0 {
			if err := db.Create(&invoices).Error; err != nil {
				return err
			}
		}
	}

	fmt.Println("seed completed")

	return nil
}

func randomSegment(r *rand.Rand) domain.ClientSegment {
	return segments[r.Intn(len(segments))]
}

func randomStatus(segment domain.ClientSegment, r *rand.Rand) domain.ClientStatus {
	switch segment {
	case domain.SegmentEnterprise:
		return domain.StatusActive
	case domain.SegmentZombie:
		return domain.StatusDelinquent
	default:
		return status[r.Intn(len(status))]
	}
}

func randomBilling(segment domain.ClientSegment, r *rand.Rand) float64 {
	switch segment {
	case domain.SegmentEnterprise:
		return float64(r.Intn(10000) + 5000)
	case domain.SegmentStartup:
		return float64(r.Intn(3000) + 1000)
	default:
		return float64(r.Intn(1500) + 200)
	}
}

func generateInvoices(clientID uint, r *rand.Rand) []domain.Invoice {
	count := rand.Intn(5) + 1

	invoices := make([]domain.Invoice, 0, count)

	for i := 0; i < count; i++ {
		status := randomInvoiceStatus(r)

		invoice := domain.Invoice{
			ClientID: clientID,
			Amount:   float64(r.Intn(2000) + 100),
			DueDate:  time.Now().AddDate(0, 0, -r.Intn(90)),
			Status:   status,
		}

		invoices = append(invoices, invoice)
	}

	return invoices
}

func randomInvoiceStatus(r *rand.Rand) domain.InvoiceStatus {
	val := r.Intn(100)

	switch {
	case val < 60:
		return domain.InvoicePaid
	case val < 85:
		return domain.InvoicePending
	default:
		return domain.InvoiceOverdue
	}
}
