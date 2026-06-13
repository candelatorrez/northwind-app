export type ClientSegment =
  | "enterprise"
  | "startup"
  | "standard"
  | "zombie";

export type ClientStatus =
  | "active"
  | "at_risk"
  | "delinquent"
  | "suspended";

export interface Client {
  id: number;
  name: string;
  email: string;
  segment: ClientSegment;
  status: ClientStatus;
  monthlyBilling: number;
  createdAt: string;
  updatedAt: string;
}