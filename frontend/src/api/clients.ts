import type { CollectionAction, CreateActionRequest } from "../types/action";
import type { Client } from "../types/client";
import type { RiskSnapshot } from "../types/risk";
import { api } from "./client";


export async function getClients(): Promise<Client[]> {
    const response = await api.get<Client[]>("/clients");

    return response.data;
}

export async function getClient(id: number): Promise<Client> {
    const response = await api.get<Client>(`/clients/${id}`);
    
    return response.data;
}

export async function getClientActions(id: number): Promise<CollectionAction[]> {
    const response = await api.get<CollectionAction[]>(`/clients/${id}/actions`)

    return response.data;
}

export async function createAction(id: number, payload: CreateActionRequest): Promise<CollectionAction> {
    const response = await api.post<CollectionAction>(`/clients/${id}/actions`, payload);

    return response.data;
}

export async function getRisk(id: number): Promise<RiskSnapshot> {
    const response = await api.get<RiskSnapshot>(`/clients/${id}/risk`);

    return response.data;
}