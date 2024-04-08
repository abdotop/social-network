package models

import (
	"database/sql"
	"fmt"
	"html"
	"time"

	"github.com/google/uuid"
)

type Events []Event
type EventParticipants []EventParticipant

type Event struct {
	ID           uuid.UUID `sql:"type:uuid;primary key"`
	GroupID      uuid.UUID `sql:"type:uuid"`
	CreatorID    uuid.UUID `sql:"type:uuid"`
	Title        string    `sql:"type:varchar(255)" json:"title"`
	Description  string    `sql:"type:text" json:"description"`
	DateTime     time.Time `json:"date_time"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
	Participants EventParticipants
}

type EventResponse string

const (
	ResponseGoing    EventResponse = "going"
	ResponseNotGoing EventResponse = "not_going"
)

type EventParticipant struct {
	ID        uuid.UUID     `sql:"type:uuid;primary key"`
	MemberID  uuid.UUID     `sql:"type:uuid"`
	EventID   uuid.UUID     `sql:"type:uuid"`
	Response  EventResponse `json:"response"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	User      User
}

func (e *Event) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	e.ID = uuid.New()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()

	query := `INSERT INTO events (id, group_id, creator_id, title, description, date_time, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(e.ID, e.GroupID, e.CreatorID, html.EscapeString(e.Title), html.EscapeString(e.Description), e.DateTime, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Get(db *sql.DB, id uuid.UUID, getParticipants, getUser bool) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, group_id, creator_id, title, description, date_time, created_at, updated_at FROM events WHERE id = $1`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stm.Close()

	err = stm.QueryRow(id).Scan(&e.ID, &e.GroupID, &e.CreatorID, &e.Title, &e.Description, &e.DateTime, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error getting event: %v", err)
	}

	if getParticipants {
		p := EventParticipants{}
		if err := p.GetEventParticipants(db, e.ID, getUser); err != nil {
			return err
		}
	}

	return nil
}

func (e *Event) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	e.UpdatedAt = time.Now()

	query := `UPDATE events SET title = $1, description = $2, date_time = $3, updated_at = $4 WHERE id = $5`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(html.EscapeString(e.Title), html.EscapeString(e.Description), e.DateTime, e.UpdatedAt, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete(db *sql.DB) error {
	query := `UPDATE events SET deleted_at = $1 WHERE id = $2`

	if err := e.Participants.DeleteEventParticipants(db, e.ID); err != nil {
		return err
	}

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec(time.Now(), e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *EventParticipant) CreateParticipant(db *sql.DB, eventID, memberID uuid.UUID) error {
	p.ID = uuid.New()
	p.EventID = eventID
	p.MemberID = memberID
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	query := `INSERT INTO events_participants (id, event_id, member_id, response, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec(p.ID, p.EventID, p.MemberID, p.Response, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *EventParticipant) GetParticipant(db *sql.DB, eventID, memberID, UserID uuid.UUID, getUser bool) error {
	query := `SELECT id, event_id, member_id, response, created_at, updated_at FROM events_participants WHERE event_id = $1 AND member_id = $2`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stm.Close()

	err = stm.QueryRow(eventID, memberID).Scan(&p.ID, &p.EventID, &p.MemberID, &p.Response, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error getting event participant: %v", err)
	}

	if getUser {
		u := User{}
		if err := u.Get(db, UserID); err != nil {
			return err
		}
		p.User = u
	}

	return nil
}

func (p *EventParticipant) UpdateParticipant(db *sql.DB) error {
	p.UpdatedAt = time.Now()

	query := `UPDATE events_participants SET response = $1, updated_at = $2 WHERE id = $3`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(p.Response, p.UpdatedAt, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *EventParticipant) Delete(db *sql.DB) error {
	query := `UPDATE events_participants SET deleted_at = $1 WHERE id = $2`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec(time.Now(), p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *EventParticipants) GetEventParticipants(db *sql.DB, eventID uuid.UUID, getUser bool) error {
	query := `SELECT id, event_id, member_id, response, created_at, updated_at FROM events_participants WHERE event_id = $1`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(eventID)
	if err != nil {
		return fmt.Errorf("error getting event participants: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		participant := EventParticipant{}
		err = rows.Scan(&participant.ID, &participant.EventID, &participant.MemberID, &participant.Response, &participant.CreatedAt, &participant.UpdatedAt)
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("error scanning event participants: %v", err)
		}

		if getUser {
			m := GroupMember{}
			if err := m.GetMemberById(db, participant.MemberID, true); err != nil {
				return err
			}

			participant.User = m.User
		}

		*p = append(*p, participant)
	}

	return nil
}

func (e *Events) GetGroupEvents(db *sql.DB, groupID uuid.UUID, getParticipants, getUser bool) error {
	query := `SELECT id, group_id, creator_id, title, description, date_time, created_at, updated_at FROM events WHERE group_id = $1`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(groupID)
	if err != nil {
		return fmt.Errorf("error getting group events: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := Event{}
		err = rows.Scan(&event.ID, &event.GroupID, &event.CreatorID, &event.Title, &event.Description, &event.DateTime, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			return fmt.Errorf("error scanning group events: %v", err)
		}

		if getParticipants {
			p := EventParticipants{}
			if err := p.GetEventParticipants(db, event.ID, getUser); err != nil {
				return err
			}
			event.Participants = p
		}

		*e = append(*e, event)
	}

	return nil
}

func (p *EventParticipants) DeleteEventParticipants(db *sql.DB, eventID uuid.UUID) error {
	query := `UPDATE events_participants SET deleted_at = $1 WHERE group_id = $2`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec(time.Now(), eventID)
	if err != nil {
		return err
	}

	return nil
}
