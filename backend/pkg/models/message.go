package models

import (
	"database/sql"
	"html"
	"time"

	"github.com/google/uuid"
)

type PrivateMessages []PrivateMessage
type GroupMessages []GroupMessage
type PrivateMessage struct {
	ID         uuid.UUID `sql:"type:uuid;primary key"`
	SenderID   uuid.UUID `sql:"type:uuid"`
	ReceiverID uuid.UUID `sql:"type:uuid"`
	Content    string    `sql:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}
type GroupMessage struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	SenderID  uuid.UUID `sql:"type:uuid"`
	Content   string    `sql:"type:text"`
	Sender    User      `sql:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (m *PrivateMessage) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	query := `INSERT INTO private_messages (id, sender_id, receiver_id, content, created_at) 
		VALUES ($1, $2, $3, $4, $5)`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(m.ID, m.SenderID, m.ReceiverID, html.EscapeString(m.Content), m.CreatedAt)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"ID":         m.ID,
		"SenderID":   m.SenderID,
		"ReceiverID": m.ReceiverID,
		"Content":    m.Content,
		"CreatedAt":  m.CreatedAt,
	}

	Data <- map[string]interface{}{
		"key":  "private_message",
		"data": data,
	}

	return nil
}
func (m *GroupMessage) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	query := `INSERT INTO group_messages (id, group_id, sender_id, content, created_at) 
		VALUES ($1, $2, $3, $4, $5)`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(m.ID, m.GroupID, m.SenderID, html.EscapeString(m.Content), m.CreatedAt)
	if err != nil {
		return err
	}
	// Data.Store("group_message_id_"+id, map[string]interface{}{
	// 	"ID":        m.ID,
	// 	"GroupID":   m.GroupID,
	// 	"SenderID":  m.SenderID,
	// 	"Content":   m.Content,
	// 	"Sender":    m.Sender,
	// 	"CreatedAt": m.CreatedAt,
	// })
	return nil
}
func (m *PrivateMessage) Get(db *sql.DB, id uuid.UUID) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, sender_id, receiver_id, content, created_at, updated_at, deleted_at FROM private_messages WHERE id = $1 AND deleted_at IS NULL`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	err = stm.QueryRow(id).Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *PrivateMessage) GetLastMessage(db *sql.DB, senderID, receiverID uuid.UUID) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, sender_id, receiver_id, content, created_at, updated_at, deleted_at FROM private_messages WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1) AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 1`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	err = stm.QueryRow(senderID, receiverID).Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *GroupMessage) Get(db *sql.DB, id uuid.UUID) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, group_id, sender_id, content, created_at, updated_at, deleted_at FROM group_messages WHERE id = $1 AND deleted_at IS NULL`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	err = stm.QueryRow(id).Scan(&m.ID, &m.GroupID, &m.SenderID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}
func (m *PrivateMessage) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	m.UpdatedAt = time.Now()
	query := `UPDATE private_messages SET content = $1, updated_at = $2 WHERE id = $3`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(html.EscapeString(m.Content), m.UpdatedAt, m.ID)
	if err != nil {
		return err
	}
	return nil
}
func (m *GroupMessage) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	m.UpdatedAt = time.Now()
	query := `UPDATE group_messages SET content = $1, updated_at = $2 WHERE id = $3`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(html.EscapeString(m.Content), m.UpdatedAt, m.ID)
	if err != nil {
		return err
	}
	return nil
}
func (m *PrivateMessage) Delete(db *sql.DB) error {
	query := `UPDATE private_messages SET deleted_at = $1 WHERE id = $2`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(time.Now(), m.ID)
	if err != nil {
		return err
	}
	return nil
}
func (m *GroupMessage) Delete(db *sql.DB) error {
	query := `UPDATE group_messages SET deleted_at = $1 WHERE id = $2`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(time.Now(), m.ID)
	if err != nil {
		return err
	}
	return nil
}
func (ms *PrivateMessages) GetPrivateMessages(db *sql.DB, receiverID, senderID uuid.UUID) error {
	query := `
        SELECT id, sender_id, receiver_id, content, created_at, updated_at, deleted_at 
        FROM private_messages 
        WHERE 
            (receiver_id = $1 AND sender_id = $2) OR 
            (receiver_id = $2 AND sender_id = $1) AND 
            deleted_at IS NULL
    `
	rows, err := db.Query(query, receiverID, senderID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var m PrivateMessage
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return err
		}
		*ms = append(*ms, m)
	}
	return nil
}
func (ms *GroupMessages) GetGroupMessages(db *sql.DB, groupID uuid.UUID) error {
	query := `SELECT id, group_id, sender_id, content, created_at, updated_at, deleted_at FROM group_messages WHERE group_id = $1 AND deleted_at IS NULL`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var m GroupMessage
		if err := rows.Scan(&m.ID, &m.GroupID, &m.SenderID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return err
		}
		*ms = append(*ms, m)
	}
	// get user of each message
	for i, m := range *ms {
		user := User{}
		err := user.Get(db, m.SenderID)
		if err != nil {
			return err
		}
		(*ms)[i].Sender = user
	}
	return nil
}
func (ms *PrivateMessages) GetPrivateMessagesBetween(db *sql.DB, senderID, receiverID uuid.UUID) error {
	query := `SELECT id, sender_id, receiver_id, content, created_at, updated_at, deleted_at FROM private_messages WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1) AND deleted_at IS NULL`
	rows, err := db.Query(query, senderID, receiverID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var m PrivateMessage
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return err
		}
		*ms = append(*ms, m)
	}
	return nil
}
func (ms *GroupMessages) ClearGroupMessages(db *sql.DB, groupID uuid.UUID) error {
	query := `UPDATE group_messages SET deleted_at = $1 WHERE group_id = $2`
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(time.Now(), groupID)
	if err != nil {
		return err
	}
	return nil
}
