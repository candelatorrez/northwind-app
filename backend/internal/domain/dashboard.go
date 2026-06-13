package domain

type DashboardMetrics struct {
	TotalClients      int     `json:"total_clients"`
	HighRiskClients   int     `json:"high_risk_clients"`
	OverdueInvoices   int     `json:"overdue_invoices"`
	OutstandingAmount float64 `json:"outstanding_amount"`
}

type ClientOverview struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Segment        string  `json:"segment"`
	Status         string  `json:"status"`
	MonthlyBilling float64 `json:"monthlyBilling"`
	RiskScore      int     `json:"riskScore"`
	RiskLevel      string  `json:"riskLevel"`
}
