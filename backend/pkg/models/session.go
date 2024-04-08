package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const SessionDuration = 24 * time.Hour

type Sessions []Session

type Session struct {
	ID           uuid.UUID `sql:"type:uuid;primary key"`
	UserID       uuid.UUID `sql:"type:uuid"`
	SessionToken string    `sql:"type:varchar(255)"`
	CreatedAt    time.Time
	ExpiresAt    time.Time
	DeletedAt    sql.NullTime
	User         User
}

// Create a new session
func (s *Session) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	s.ID = uuid.New()
	s.CreatedAt = time.Now()
	s.ExpiresAt = time.Now().Add(SessionDuration)

	query := `INSERT INTO sessions (id, user_id, session_token, created_at, expires_at) 
		VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(query, s.ID, s.UserID, s.SessionToken, s.CreatedAt, s.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}

// Get a session by its ID
func (s *Session) Get(db *sql.DB, id uuid.UUID, getUser bool) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, user_id, session_token, created_at, expires_at, deleted_at FROM sessions WHERE id = $1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	err = stm.QueryRow(id).Scan(&s.ID, &s.UserID, &s.SessionToken, &s.CreatedAt, &s.ExpiresAt, &s.DeletedAt)
	if err != nil {
		return err
	}

	if getUser {
		err = s.User.Get(db, s.UserID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Get a session by its token
func (s *Session) GetByToken(db *sql.DB, token string, getUser bool) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, user_id, session_token, created_at, expires_at, deleted_at FROM sessions WHERE session_token = $1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	err = stm.QueryRow(token).Scan(&s.ID, &s.UserID, &s.SessionToken, &s.CreatedAt, &s.ExpiresAt, &s.DeletedAt)
	if err != nil {
		return err
	}

	if getUser {
		err = s.User.Get(db, s.UserID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete a session
func (s *Session) Delete(db *sql.DB) error {
	query := `UPDATE sessions SET deleted_at = $1 WHERE id = $2`

	_, err := db.Exec(query, time.Now(), s.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete all sessions for a user
func (s *Session) DeleteByUser(db *sql.DB) error {
	query := `UPDATE sessions SET deleted_at = $1 WHERE user_id = $2`

	_, err := db.Exec(query, time.Now(), s.UserID)
	if err != nil {
		return err
	}
	return nil
}

// Delete all expired sessions
func (s *Session) DeleteExpired(db *sql.DB) error {
	query := `UPDATE sessions SET deleted_at = $1 WHERE expires_at < $2`

	_, err := db.Exec(query, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// Get all sessions for a user
func (s *Sessions) GetByUser(db *sql.DB, userID uuid.UUID) error {
	query := `SELECT id, user_id, session_token, created_at, expires_at, deleted_at FROM sessions WHERE user_id = $1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	rows, err := stm.Query(userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var session Session
		err = rows.Scan(&session.ID, &session.UserID, &session.SessionToken, &session.CreatedAt, &session.ExpiresAt, &session.DeletedAt)
		if err != nil {
			return err
		}
		*s = append(*s, session)
	}
	return nil
}

// Get all sessions
func (s *Sessions) GetAll(db *sql.DB, getUser bool) error {
	query := `SELECT id, user_id, session_token, created_at, expires_at, deleted_at FROM sessions WHERE deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var session Session
		err = rows.Scan(&session.ID, &session.UserID, &session.SessionToken, &session.CreatedAt, &session.ExpiresAt, &session.DeletedAt)
		if err != nil {
			return err
		}

		if getUser {
			err = session.User.Get(db, session.UserID)
			if err != nil {
				return err
			}
		}

		*s = append(*s, session)
	}
	return nil
}
