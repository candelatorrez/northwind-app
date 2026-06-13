import type {
  DashboardMetrics,
  DashboardMetricsDTO,
} from "../types/dashboard";

export function toDashboardMetrics(
  dto: DashboardMetricsDTO,
): DashboardMetrics {

  return {
    totalClients: dto.total_clients,

    highRiskClients:
      dto.high_risk_clients,

    overdueInvoices:
      dto.overdue_invoices,

    outstandingAmount:
      dto.outstanding_amount,
  };
}