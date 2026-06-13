import KPICard from "../components/KPICard";
import ErrorState from "../components/ErrorState";
import { useDashboard } from "../hooks/useDashboard";
import LoadingState from "../components/LoadingState";

export default function Dashboard() {
  const {
    metrics,
    loading,
    error,
  } = useDashboard();

  if (loading) {
    return <LoadingState />;
  }

  if (!metrics) {
    return (
      <ErrorState message="No metrics found" />
    );
  }

   if (error) {
    return (
      <ErrorState message={error} />
    );
  }

  console.log(metrics)

  return (
    <div
      style={{
        padding: "24px",
      }}
    >
      <h1>
        Northwind Collections
      </h1>

      <div
        style={{
          display: "grid",
          gridTemplateColumns:
            "repeat(4,1fr)",
          gap: "16px",
          marginTop: "24px",
        }}
      >
        <KPICard
          title="Clients"
          value={metrics.totalClients}
        />

        <KPICard
          title="High Risk"
          value={metrics.highRiskClients}
        />

        <KPICard
          title="Overdue Invoices"
          value={metrics.overdueInvoices}
        />

        <KPICard
          title="Outstanding Amount"
          value={`$${metrics.outstandingAmount}`}
        />
      </div>
    </div>
  );
}