CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR NOT NULL,
    last_name VARCHAR,
    email VARCHAR UNIQUE,
    email_verified_at TIMESTAMP WITH TIME ZONE ,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE  NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);