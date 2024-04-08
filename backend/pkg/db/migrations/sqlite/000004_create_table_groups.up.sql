CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY,
    title TEXT,
    description TEXT,
    banner_url TEXT,
    creator_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);