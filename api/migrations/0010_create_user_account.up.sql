CREATE TABLE user_accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bank_id UUID REFERENCES banks(id) ON DELETE SET NULL,
    wallet_id UUID REFERENCES wallets(id) ON DELETE SET NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_number VARCHAR,
    phone_number VARCHAR,
    name VARCHAR NOT NULL,
    balance DECIMAL(10,2) NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_user_accounts_deleted_at ON user_accounts(deleted_at);
CREATE INDEX idx_user_accounts_user_id ON user_accounts(user_id);
CREATE INDEX idx_user_accounts_bank_id ON user_accounts(bank_id);
CREATE INDEX idx_user_accounts_wallet_id ON user_accounts(wallet_id); 