-- Enable uuid-ossp extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the files table
CREATE TABLE files (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    mime_type VARCHAR NOT NULL,
    file_name VARCHAR UNIQUE NOT NULL,
    path_name VARCHAR NOT NULL,
    alt_text TEXT,
    is_private BOOLEAN NOT NULL DEFAULT FALSE, -- Added NOT NULL for clarity
    variants JSONB
);

-- Indexes for common queries
CREATE INDEX idx_files_deleted_at ON files (deleted_at); -- For soft deletes
CREATE INDEX idx_files_created_at ON files (created_at); -- For sorting by creation time