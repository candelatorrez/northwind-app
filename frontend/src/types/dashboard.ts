export interface DashboardMetricsDTO {
  total_clients: number;
  high_risk_clients: number;
  overdue_invoices: number;
  outstanding_amount: number;
}

export interface DashboardMetrics {
  totalClients: number;
  highRiskClients: number;
  overdueInvoices: number;
  outstandingAmount: number;
}

export interface ClientOverview {
  id: number;
  name: string;
  segment: string;
  status: string;
  monthlyBilling: number;
  riskScore: number;
  riskLevel: string;
}