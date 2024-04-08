CREATE TABLE IF NOT EXISTS group_members (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES groups(id),
    member_id UUID REFERENCES users(id),
    status TEXT CHECK(status = 'invited' OR status = 'requesting' OR status = 'accepted' OR status = 'declined'),
    role TEXT CHECK(role = 'admin' OR role = 'user'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);