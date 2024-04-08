package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FollowerStatus string
type Followers []Follower

const (
	StatusRequested FollowerStatus = "requested"
	StatusAccepted  FollowerStatus = "accepted"
	StatusDeclined  FollowerStatus = "declined"
)

type Follower struct {
	ID         uuid.UUID `sql:"type:uuid;primary key"`
	FollowerID uuid.UUID `sql:"type:uuid"`
	FolloweeID uuid.UUID `sql:"type:uuid"`
	Status     FollowerStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Create a new follower
func (follower *Follower) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	// Define the follow default properties
	follower.ID = uuid.New()
	follower.CreatedAt = time.Now()
	follower.UpdatedAt = time.Now()

	// Define the query
	query := `INSERT INTO followers (id, follower_id, followee_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	// Prepare the query
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	// Execute the query
	_, err = stmt.Exec(
		follower.ID,
		follower.FollowerID,
		follower.FolloweeID,
		follower.Status,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Get a follower
func (follower *Follower) Get(db *sql.DB, reverse ...bool) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	if len(reverse) > 0 && reverse[0] {
		follower.FollowerID, follower.FolloweeID = follower.FolloweeID, follower.FollowerID
	}

	// Define the query
	query := `SELECT id, follower_id, followee_id, status, created_at, updated_at FROM followers WHERE follower_id = $1 AND followee_id = $2 AND deleted_at IS NULL`

	// Execute the query
	err := db.QueryRow(query, follower.FollowerID, follower.FolloweeID).Scan(
		&follower.ID,
		&follower.FollowerID,
		&follower.FolloweeID,
		&follower.Status,
		&follower.CreatedAt,
		&follower.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Update a follower
func (follower *Follower) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	// Define the query
	query := `UPDATE followers SET status = $1, updated_at = $2 WHERE id = $3`

	// Prepare the query
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	// Execute the query
	_, err = stmt.Exec(follower.Status, time.Now(), follower.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Delete a follower
func (follower *Follower) Delete(db *sql.DB) error {
	// Define the query
	query := `UPDATE followers SET deleted_at = $1 WHERE id = $2`

	// Prepare the query
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	// Execute the query
	_, err = stmt.Exec(time.Now(), follower.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Get all followers by followee id
func (followers *Followers) GetAllByFolloweeID(db *sql.DB, followeeID uuid.UUID) error {
	// Define the query
	query := `SELECT id, follower_id, followee_id, status, created_at, updated_at FROM followers WHERE followee_id = $1 AND status= "accepted" AND  deleted_at IS NULL`

	// Execute the query
	rows, err := db.Query(query, followeeID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var follower Follower
		err := rows.Scan(
			&follower.ID,
			&follower.FollowerID,
			&follower.FolloweeID,
			&follower.Status,
			&follower.CreatedAt,
			&follower.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to execute the query. %v", err)
		}
		*followers = append(*followers, follower)
	}

	return nil
}



// count all followers by followee id
func (followers *Followers) CountAllByFolloweeID(db *sql.DB, followeeID uuid.UUID) int {
	// Define the query
	query := `SELECT COUNT(id) FROM followers WHERE followee_id = $1 AND status = $2 AND deleted_at IS NULL`
	var count int
	// Execute the query
	err := db.QueryRow(query, followeeID, StatusAccepted).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

// Get all followers by follower id
func (followers *Followers) GetAllByFollowerID(db *sql.DB, followerID uuid.UUID) error {
	// Define the query
	query := `SELECT id, follower_id, followee_id, status, created_at, updated_at FROM followers WHERE follower_id = $1 AND status= "accepted" AND deleted_at IS NULL`

	// Execute the query
	rows, err := db.Query(query, followerID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var follower Follower
		err := rows.Scan(
			&follower.ID,
			&follower.FollowerID,
			&follower.FolloweeID,
			&follower.Status,
			&follower.CreatedAt,
			&follower.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to execute the query. %v", err)
		}
		*followers = append(*followers, follower)
	}

	return nil
}

// count all followers by follower id
func (followers *Followers) CountAllByFollowerID(db *sql.DB, followerID uuid.UUID) int {
	// Define the query
	query := `SELECT COUNT(id) FROM followers WHERE follower_id = $1 AND status = $2 AND deleted_at IS NULL`
	var count int
	// Execute the query
	err := db.QueryRow(query, followerID, StatusAccepted).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}
