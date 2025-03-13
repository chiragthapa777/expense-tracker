-- Enable the UUID extension if not enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create login_logs table
CREATE TABLE login_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    method VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE ,
    CONSTRAINT fk_login_logs_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes for faster queries
CREATE INDEX idx_login_logs_created_at ON login_logs (created_at);
CREATE INDEX idx_login_logs_user_id ON login_logs (user_id);
