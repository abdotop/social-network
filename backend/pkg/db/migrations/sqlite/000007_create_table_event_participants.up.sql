CREATE TABLE IF NOT EXISTS events_participants (
    id UUID PRIMARY KEY,
    event_id UUID REFERENCES events(id),
    member_id UUID REFERENCES users(id),
    response TEXT CHECK(response = 'going' OR response = 'not_going'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
