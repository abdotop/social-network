CREATE TABLE IF NOT EXISTS invitations (
    id UUID PRIMARY KEY,
    inviting_user_id UUID REFERENCES users(id),
    invited_user_id UUID REFERENCES users(id),
    group_member_id UUID REFERENCES group_members(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);