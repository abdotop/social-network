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

type PostPrivacy string
type Posts []Post

const (
	PrivacyGroup         PostPrivacy = "group"
	PrivacyPublic        PostPrivacy = "public"
	PrivacyPrivate       PostPrivacy = "private"
	PrivacyAlmostPrivate PostPrivacy = "almost private"
	PrivacyUnlisted      PostPrivacy = "unlisted"
)

type Post struct {
	ID                uuid.UUID    `json:"id" sql:"type:uuid;primary key"`
	GroupID           uuid.UUID    `sql:"type:uuid" json:"group_id"`
	UserID            uuid.UUID    `json:"user_id" sql:"type:uuid"`
	Title             string       `json:"title" sql:"type:varchar(255)"`
	Content           string       `json:"content" sql:"type:text"`
	ImageURL          string       `json:"image_url" sql:"type:varchar(255)"`
	Privacy           PostPrivacy  `json:"privacy"`
	SelectedFollowers []uuid.UUID  `json:"followersSelectedID"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	DeletedAt         sql.NullTime `json:"deleted_at"`
}

// IsPublic returns true if the post is public
func (p *Post) IsPublic() bool {
	return p.Privacy == PrivacyPublic
}

// IsPrivate returns true if the post is private
func (p *Post) IsPrivate() bool {
	return p.Privacy == PrivacyPrivate
}

// IsAlmostPrivate returns true if the post is almost private
func (p *Post) IsAlmostPrivate() bool {
	return p.Privacy == PrivacyAlmostPrivate
}

// IsUnlisted returns true if the post is unlisted
func (p *Post) IsUnlisted() bool {
	return p.Privacy == PrivacyUnlisted
}

// IsDeleted returns true if the post is deleted
func (p *Post) IsDeleted() bool {
	return p.DeletedAt.Valid
}

// PostPrivacyFromString returns the post privacy from a string
func PostPrivacyFromString(s string) (PostPrivacy, error) {
	switch s {
	case "public":
		return PrivacyPublic, nil
	case "private":
		return PrivacyPrivate, nil
	case "almost private":
		return PrivacyAlmostPrivate, nil
	case "unlisted":
		return PrivacyUnlisted, nil
	default:
		return "", fmt.Errorf("invalid post privacy")
	}
}

// Create inserts a new post into the database
func (p *Post) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	// Define the post default properties
	if !p.IsValid() {
		return errors.New("something wrong with the comment")
	}
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	query := `INSERT INTO posts (id, user_id, group_id, title, content, image_url, privacy, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		p.ID,
		p.UserID.String(),
		p.GroupID.String(),
		html.EscapeString(p.Title),
		html.EscapeString(p.Content),
		html.EscapeString(p.ImageURL),
		p.Privacy,
		p.CreatedAt,
		p.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v, privacy %v", err, p.Privacy)
	}
	if p.Privacy != PrivacyAlmostPrivate {
		return nil
	}

	return p.saveFolowersSelection(db)
}

func (p *Post) IsValid() bool {
	if p.ID.String() != uuid.Nil.String() {
		return false
	}
	if p.Content == "" || strings.TrimSpace(p.Content) == "" {
		return false
	}

	return true
}

// Get retrieves a post from the database
func (p *Post) Get(db *sql.DB, id uuid.UUID) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at, group_id  FROM posts WHERE id = $1 AND deleted_at IS NULL`

	err := db.QueryRow(query, id).Scan(
		&p.ID,
		&p.UserID,
		&p.Title,
		&p.Content,
		&p.ImageURL,
		&p.Privacy,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
		&p.GroupID,
	)
	fmt.Println(p.Content, p.GroupID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no post found with id %v", id)
		}
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Update modifies a post in the database
func (p *Post) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	query := `UPDATE posts SET title = $1, content = $2, image_url = $3, privacy = $4, updated_at = $5 WHERE id = $6`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		html.EscapeString(p.Title),
		html.EscapeString(p.Content),
		html.EscapeString(p.ImageURL),
		p.Privacy,
		time.Now(),
		p.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Delete removes a post from the database
func (p *Post) Delete(db *sql.DB) error {
	query := `UPDATE posts SET deleted_at = $1 WHERE id = $2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		time.Now(),
		p.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetUserPosts retrieves all the posts from a user
func (p *Posts) GetUserPosts(db *sql.DB, userID uuid.UUID) error {
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE user_id = $1 AND deleted_at IS NULL`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImageURL,
			&post.Privacy,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}
		*p = append(*p, post)
	}

	return nil
}

// GetAll retrieves all the posts from the database
func (p *Posts) GetAll(db *sql.DB) error {
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE deleted_at IS NULL`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImageURL,
			&post.Privacy,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}
		*p = append(*p, post)
	}

	return nil
}

func (p *Post) saveFolowersSelection(db *sql.DB) error {
	// Préparez la requête SQL pour insérer les utilisateurs sélectionnés dans la table userSelected
	query := `INSERT INTO selected_users (id, post_id, user_id) VALUES (? ,?, ?)`

	// Préparez la requête SQL
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	// Exécutez la requête pour chaque utilisateur sélectionné
	for _, userID := range p.SelectedFollowers {
		_, err = stmt.Exec(uuid.New(), p.ID, userID)
		if err != nil {
			return fmt.Errorf("unable to execute the query. %v", err)
		}
	}
	return nil
}

func (p *Posts) GetAvailablePostForUser(db *sql.DB, userID uuid.UUID) error {
	query := `SELECT * FROM posts WHERE 
    (privacy = 'public' OR 
    (privacy = 'private' AND deleted_at IS NULL AND EXISTS (SELECT 1 FROM followers f WHERE posts.user_id = f.followee_id AND f.follower_id = ? AND f.status = 'accepted')) OR 
    (privacy = 'almost private' AND deleted_at IS NULL AND EXISTS (SELECT 1 FROM selected_users us WHERE posts.id = us.post_id AND us.user_id = ?)) OR 
    user_id = ?) AND 
    deleted_at IS NULL AND (group_id  IS NULL OR group_id = "00000000-0000-0000-0000-000000000000") AND privacy != 'group' 
    ORDER BY created_at DESC`
	if err := p.getPostsFromQuery(db, query, userID, userID, userID); err != nil {
		return err
	}
	return nil
}

// Une méthode d'aide pour exécuter une requête et remplir les posts
func (p *Posts) getPostsFromQuery(db *sql.DB, query string, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.GroupID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImageURL,
			&post.Privacy,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}
		*p = append(*p, post)
	}
	return nil
}
func (Posts *Posts) ExploitForRendering(db *sql.DB) []map[string]interface{} {
	valueToReturn := []map[string]interface{}{}
	for _, v := range *Posts {
		user := User{}
		user.Get(db, v.UserID)
		valueToReturn = append(valueToReturn, v.ExploitForRendering(db))
	}
	return valueToReturn
}
func (p *Post) ExploitForRendering(db *sql.DB) map[string]interface{} {
	user := User{}
	user.Get(db, p.UserID)
	postComments := Comments{}
	postComments.GetCommentsForPost(db, p.ID)

	return map[string]interface{}{
		"group_id":           p.GroupID,
		"id":                 p.ID,
		"userCompleteName":   user.FirstName + " " + user.LastName,
		"imageUrl":           p.ImageURL,
		"content":            p.Content,
		"userAvatarImageUrl": user.AvatarImage,
		"createdAt":          timeAgo(p.CreatedAt),
		"comments":           postComments.ExploitForRendering(db, string(p.Privacy), p.GroupID),
		"userOwnerNickname":  user.Nickname,
	}
}

func (p *Posts) GetPostByGroupId(db *sql.DB, groupID string) error {
	// Préparez la requête SQL pour récupérer les posts par l'ID du groupe
	query := `SELECT * FROM posts WHERE group_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC`

	// Préparez la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	// Exécutez la requête
	rows, err := stmt.Query(groupID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	// Parcourez les résultats et ajoutez chaque post au slice
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.GroupID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImageURL,
			&post.Privacy,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.DeletedAt,
		)
		if err != nil {
			fmt.Errorf("unable to scan the row. %v", err)
		}
		*p = append(*p, post)
	}
	return nil
}

func timeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	switch {
	case diff.Hours() > 24:
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	case diff.Hours() > 1:
		hours := int(diff.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	case diff.Minutes() > 1:
		minutes := int(diff.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	case diff.Seconds() < 1:
		return "now"
	default:
		seconds := int(diff.Seconds())
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}
func CountPostsByUser(db *sql.DB, userID uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM posts WHERE user_id = $1 AND deleted_at IS NULL`

	var count int
	err := db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("unable to execute the query. %v", err)
	}

	return count, nil
}
