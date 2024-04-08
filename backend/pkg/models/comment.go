package models

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Comment struct represents a comment in the database.
type Comment struct {
	ID        uuid.UUID    `json:"id" sql:"type:uuid;primary key"`
	UserID    uuid.UUID    `json:"user_id" sql:"type:uuid"`
	PostID    uuid.UUID    `json:"post_id" sql:"type:uuid"`
	Content   string       `json:"content" sql:"type:text"`
	ImageURL  string       `json:"image_url" sql:"type:text"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

// Comments is a slice of Comment structs.
type Comments []Comment

// IsValid checks if the comment is valid.
func (c *Comment) IsValid() bool {
	// Check if the content is not empty and does not consist only of whitespace characters.
	if (c.Content == "" || strings.TrimSpace(c.Content) == "") && c.ImageURL == "" {
		return false
	}

	// Check if the PostID is not empty.
	if c.PostID == uuid.Nil {
		return false
	}

	// If all checks pass, the comment is valid.
	return true
}

// Create inserts a new comment into the database.
func (c *Comment) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	// Validate the comment before proceeding.
	if !c.IsValid() {
		return errors.New("comment is not valid")
	}

	c.ID = uuid.New()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	query := `INSERT INTO comments (id, user_id, post_id, content, image_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		c.ID,
		c.UserID,
		c.PostID,
		html.EscapeString(c.Content),
		html.EscapeString(c.ImageURL),
		c.CreatedAt,
		c.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Get retrieves a comment from the database.
func (c *Comment) Get(db *sql.DB, id uuid.UUID) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, user_id, post_id, content, image_url, created_at, updated_at, deleted_at FROM comments WHERE id = $1 AND deleted_at IS NULL`

	err := db.QueryRow(query, id).Scan(
		&c.ID,
		&c.UserID,
		&c.PostID,
		&c.Content,
		&c.ImageURL,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no comment found with id %v", id)
		}
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Update modifies a comment in the database.
func (c *Comment) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	query := `UPDATE comments SET content = $1, image_url = $2, updated_at = $3 WHERE id = $4`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		html.EscapeString(c.Content),
		html.EscapeString(c.ImageURL),
		time.Now(),
		c.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Delete removes a comment from the database.
func (c *Comment) Delete(db *sql.DB) error {
	query := `UPDATE comments SET deleted_at = $1 WHERE id = $2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		time.Now(),
		c.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	return nil
}

// GetCommentsForPost retrieves all comments for a specific post.
func (c *Comments) GetCommentsForPost(db *sql.DB, postID uuid.UUID) error {
	query := `SELECT id, user_id, post_id, content, image_url, created_at, updated_at, deleted_at FROM comments WHERE post_id = $1 AND deleted_at IS NULL`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
			&comment.ImageURL,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}
		*c = append(*c, comment)
	}

	return nil
}
func (c *Comment) ExploitForRendering(db *sql.DB, postPrivacy string, postGroupId uuid.UUID) map[string]interface{} {
	user := User{}
	user.Get(db, c.UserID)
	return map[string]interface{}{
		"group_id":           postGroupId,
		"postPrivacy":        postPrivacy,
		"id":                 c.ID,
		"post_id":            c.PostID,
		"userCompleteName":   user.FirstName + " " + user.LastName,
		"imageUrl":           c.ImageURL,
		"content":            c.Content,
		"userAvatarImageUrl": user.AvatarImage,
		"createdAt":          timeAgo(c.CreatedAt),
		"userOwnerNickname":  user.Nickname,
	}
}
func (comments *Comments) ExploitForRendering(db *sql.DB, postPrivacy string, postGroupId uuid.UUID) []map[string]interface{} {
	valueToReturn := []map[string]interface{}{}
	for _, v := range *comments {
		user := User{}
		user.Get(db, v.UserID)
		valueToReturn = append(valueToReturn, v.ExploitForRendering(db, postPrivacy, postGroupId))
	}
	return valueToReturn
}
