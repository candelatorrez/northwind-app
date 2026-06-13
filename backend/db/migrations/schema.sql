CREATE TABLE clients (
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,

    segment VARCHAR(50) NOT NULL,

    monthly_billing NUMERIC(12,2) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE invoices (
    id BIGSERIAL PRIMARY KEY,

    client_id BIGINT NOT NULL REFERENCES clients(id),

    amount NUMERIC(12,2) NOT NULL,

    due_date DATE NOT NULL,

    status VARCHAR(20) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE payments (
    id BIGSERIAL PRIMARY KEY,

    invoice_id BIGINT NOT NULL REFERENCES invoices(id),

    amount NUMERIC(12,2) NOT NULL,

    paid_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE collection_actions (
    id BIGSERIAL PRIMARY KEY,

    client_id BIGINT NOT NULL REFERENCES clients(id),

    action_type VARCHAR(20) NOT NULL,

    notes TEXT,

    performed_by VARCHAR(100),

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE risk_snapshots (
    id BIGSERIAL PRIMARY KEY,

    client_id BIGINT NOT NULL REFERENCES clients(id),

    score INTEGER NOT NULL,

    level VARCHAR(20) NOT NULL,

    reason TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);