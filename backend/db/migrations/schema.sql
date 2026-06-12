CREATE TABLE clients {
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,

    segment VARCHAR(255) NOT NULL,

    mothly_billing NUMERIC(12, 2) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()

};

CREATE TABLE invoices {
    id BIGSERIAL PRIMARY KEY,

    client_id BIGINT NOT NULL,

    amount NUMERIC(12, 2) NOT NULL,

    due_date DATE NOT NULL,

    status VARCHAR(20) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_invoice_client
        FOREIGN KEY (client_id)
        REFERENCES clients(id)
};

CREATE TABLE payments {
    id BIGSERIAL PRIMARY KEY,

    invoice_id BIGINT NOT NULL,

    mount NUMERIC(12, 2) NOT NULL,

    paid_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_payment_invoice
        FOREIGN KEY (invoice_id)
        REFERENCES invoices(id)
};

CREATE TABLE collection_actions {
    id BIGSERIAL PRIMARY KEY,

    client_id BIGINT NOT NULL,

    action_type VARCHAR(20) NOT NULL,

    notes TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_action_client
        FOREIGN KEY (client_id)
        REFERENCES clients(id)
};

CREATE TABLE risk_snapshots {
    id BIGSERIAL PRIMARY KEY,

    client_id BIGINT NOT NULL,

    score INTEGER NOT NULL,

    level VARCHAR(20) NOT NULL,

    reason TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_risk_client
        FOREIGN KEY (client_id)
        REFERENCES clients(id)
}