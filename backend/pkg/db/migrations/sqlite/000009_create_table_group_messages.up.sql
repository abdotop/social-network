CREATE TABLE IF NOT EXISTS group_messages (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES groups(id),
    sender_id UUID REFERENCES users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);