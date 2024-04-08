CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    group_id UUID,
    member_id UUID,
    is_invite BOOLEAN DEFAULT FALSE,
    type TEXT CHECK(type = 'follow_request'OR type = 'follow_accepted' OR type = 'follow_declined' OR type = 'unfollow' OR type = 'group_invitation' OR type = 'new_message' OR type = 'new_event'),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);