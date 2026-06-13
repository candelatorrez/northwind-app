export interface CollectionAction {
    id: number;
    clientId: number;
    type: "call" | "email" | "note";
    notes: string;
    performedBy: string;
    createdAt: string;
}

export interface CreateActionRequest {
    type: "call" | "email" | "note";
    notes: string;
}