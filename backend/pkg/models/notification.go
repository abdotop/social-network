package models

import (
	"database/sql"
	"html"
	"time"

	"github.com/google/uuid"
)

type NotificationType string
type Notifications []Notification

const (
	TypeFollowRequest   NotificationType = "follow_request"
	TypeFollowAccepted  NotificationType = "follow_accepted"
	TypeFollowDeclined  NotificationType = "follow_declined"
	TypeUnFollow        NotificationType = "unfollow"
	TypeGroupInvitation NotificationType = "group_invitation"
	TypeNewMessage      NotificationType = "new_message"
	TypeNewEvent        NotificationType = "new_event"
	// Add more types as needed
)

type Notification struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	UserID    uuid.UUID `sql:"type:uuid"`
	GroupId   uuid.UUID `sql:"type:uuid"`
	MemberId  uuid.UUID
	Is_invite bool
	ConcernID uuid.UUID `sql:"type:uuid"`
	Type      NotificationType
	Message   string `sql:"type:text"`
	CreatedAt time.Time
	DeletedAt sql.NullTime
}

// Create a new notification
func (n *Notification) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	n.ID = uuid.New()
	n.CreatedAt = time.Now()

	query := `INSERT INTO notifications (id, user_id, concern_id, group_id, member_id,is_invite,type, message, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6,$7,$8, $9)`

	_, err := db.Exec(query, n.ID, n.UserID, n.ConcernID, n.GroupId, n.MemberId, n.Is_invite, n.Type, html.EscapeString(n.Message), n.CreatedAt)
	if err != nil {
		return err
	}
	// member := new(GroupMember)
	// is_invite := false
	// if n.Type == TypeGroupInvitation {
	// 	if n.GroupId != uuid.Nil {
	// 		member.GetMember(db, n.UserID, n.GroupId, false)
	// 	} else {
	// 		is_invite = true
	// 	}
	// }

	user := new(User)
	user.Get(db, n.UserID)
	user.Password = ""
	data := map[string]interface{}{
		"id":         n.ID,
		"type":       n.Type,
		"concernID":  n.ConcernID,
		"user":       user,
		"message":    n.Message,
		"created_at": n.CreatedAt,
		"group_id":   n.GroupId,
		"member_id":  n.MemberId,
		"is_invite":  n.Is_invite,
	}

	Data <- map[string]interface{}{
		"key":  "notification",
		"data": data,
	}

	return nil
}

// Get a notification by its ID
func (n *Notification) Get(db *sql.DB, id ...uuid.UUID) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	if len(id) > 0 {
		query := `SELECT id, user_id, concern_id, group_id, member_id, is_invite,type, message, created_at, deleted_at FROM notifications WHERE id = $1 AND deleted_at IS NULL`

		stm, err := db.Prepare(query)

		if err != nil {
			return err
		}

		defer stm.Close()

		err = stm.QueryRow(id[0]).Scan(&n.ID, &n.UserID, &n.ConcernID, &n.GroupId, &n.MemberId, &n.Is_invite, &n.Type, &n.Message, &n.CreatedAt, &n.DeletedAt)
		if err != nil {
			return err
		}
	} else {
		query := `SELECT id, group_id, member_id,is_invite,message, created_at, deleted_at FROM notifications WHERE user_id = $1 AND concern_id = $2  AND type = $3 AND deleted_at IS NULL`

		stm, err := db.Prepare(query)

		if err != nil {
			return err
		}

		defer stm.Close()

		err = stm.QueryRow(n.UserID, n.ConcernID, n.Type).Scan(&n.ID, &n.GroupId, &n.MemberId, &n.Is_invite, &n.Type, &n.Message, &n.CreatedAt, &n.DeletedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete a notification
func (n *Notification) Delete(db *sql.DB) error {
	query := `UPDATE notifications SET deleted_at = $1 WHERE id = $2`

	_, err := db.Exec(query, time.Now(), n.ID)
	if err != nil {
		return err
	}
	return nil
}

// Get all notifications for a user
func (n *Notifications) GetByUser(db *sql.DB, userID uuid.UUID) error {
	query := `SELECT id, user_id, group_id, member_id,is_invite,type, message, created_at, deleted_at FROM notifications WHERE concern_id = $1 AND deleted_at IS NULL`

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
		var notif Notification
		err = rows.Scan(&notif.ID, &notif.UserID, &notif.GroupId, &notif.MemberId, &notif.Is_invite, &notif.Type, &notif.Message, &notif.CreatedAt, &notif.DeletedAt)
		if err != nil {
			return err
		}
		*n = append(*n, notif)
	}

	return nil
}

// ClearByUser deletes all notifications for a user
func (n *Notifications) ClearByUser(db *sql.DB, userID uuid.UUID) error {
	query := `UPDATE notifications SET deleted_at = $1 WHERE user_id = $2`

	_, err := db.Exec(query, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
