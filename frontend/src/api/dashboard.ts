import type { DashboardMetrics } from "../types/dashboard";
import { api } from "./client";


export async function getDashboardMetricts(): Promise<DashboardMetrics> {
    const response = await api.get<DashboardMetrics>("/dashboard/metrics")

    return response.data
};

