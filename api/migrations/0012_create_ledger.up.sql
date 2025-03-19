CREATE TABLE ledgers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    debit DECIMAL(10,2) NOT NULL DEFAULT 0,
    credit DECIMAL(10,2) NOT NULL DEFAULT 0,
    account_id UUID REFERENCES user_accounts(id) ON DELETE SET NULL,
    transaction_id VARCHAR,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_ledgers_deleted_at ON ledgers(deleted_at);
CREATE INDEX idx_ledgers_user_id ON ledgers(user_id);
CREATE INDEX idx_ledgers_account_id ON ledgers(account_id);
CREATE INDEX idx_ledgers_date ON ledgers(date); 