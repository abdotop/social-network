CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES groups(id),
    creator_id UUID REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    title TEXT,
    description TEXT,
    date_time DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);