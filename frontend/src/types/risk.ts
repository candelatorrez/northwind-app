export interface RiskSnapshot {
    id: number;
    clientId: number;
    score: number;
    level: string;
    reason: string;
    createdAt: string;
}