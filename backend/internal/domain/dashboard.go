package domain

type DashboardMetrics struct {
	TotalClients      int     `json:"total_clients"`
	HighRiskClients   int     `json:"high_risk_clients"`
	OverdueInvoices   int     `json:"overdue_invoices"`
	OutstandingAmount float64 `json:"outstanding_amount"`
}
