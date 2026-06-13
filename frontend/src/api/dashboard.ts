import axios from "axios";

import type {
  DashboardMetricsDTO,
  DashboardMetrics,
} from "../types/dashboard";

import { toDashboardMetrics }
from "../adapters/dashboardAdapter";

export async function getDashboardMetrics():
Promise<DashboardMetrics> {

  const response =
    await axios.get<DashboardMetricsDTO>(
      "http://localhost:8080/dashboard/metrics"
    );

  return toDashboardMetrics(
    response.data,
  );
}