package models

import (
	"database/sql"
	"fmt"
	"html"
	"time"

	"github.com/google/uuid"
)

type Groups []Group
type GroupMembers []GroupMember

type Group struct {
	ID           uuid.UUID `sql:"type:uuid;primary key"`
	Title        string    `sql:"type:varchar(255)"`
	Description  string    `sql:"type:text"`
	BannerURL    string    `sql:"type:varchar(255)"`
	CreatorID    uuid.UUID `sql:"type:uuid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
	GroupMembers GroupMembers
}

type GroupInvitation struct {
	ID             uuid.UUID
	InvitingUserId uuid.UUID
	InvitedUserId  uuid.UUID
	GroupMemberId  uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
}

type GroupMemberStatus string

const (
	MemberStatusInvited    GroupMemberStatus = "invited"
	MemberStatusRequesting GroupMemberStatus = "requesting"
	MemberStatusAccepted   GroupMemberStatus = "accepted"
	MemberStatusDeclined   GroupMemberStatus = "declined"
)

type GroupMemberRole string

const (
	MemberRoleAdmin GroupMemberRole = "admin"
	MemberRoleUser  GroupMemberRole = "user"
)

type GroupMember struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	MemberID  uuid.UUID `sql:"type:uuid"`
	Status    GroupMemberStatus
	Role      GroupMemberRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	User      User
}

// Create inserts a new group into the database
func (g *Group) Create(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	// Define the group default properties
	g.ID = uuid.New()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	query := `INSERT INTO groups (id, title, description, banner_url, creator_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(g.ID, html.EscapeString(g.Title), html.EscapeString(g.Description), html.EscapeString(g.BannerURL), g.CreatorID, g.CreatedAt, g.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	// Create the group creator as a admin of the group
	gm := GroupMember{
		Status: MemberStatusAccepted,
		Role:   MemberRoleAdmin,
	}
	err = gm.CreateMember(db, g.CreatorID, g.ID)
	if err != nil {
		return fmt.Errorf("unable to create group member. %v", err)
	}

	g.GroupMembers = append(g.GroupMembers, gm)

	return nil
}

// Get retrieves a group from the database
func (g *Group) Get(db *sql.DB, id uuid.UUID, getmembers, getuser bool) error {
	// Mux.RLock()
	// defer Mux.RUnlock()
	query := `SELECT id, title, description, banner_url, creator_id, created_at, updated_at, deleted_at FROM groups WHERE id=$1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	row := stm.QueryRow(id)
	err = row.Scan(
		&g.ID,
		&g.Title,
		&g.Description,
		&g.BannerURL,
		&g.CreatorID,
		&g.CreatedAt,
		&g.UpdatedAt,
		&g.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("unable to scan the row. %v", err)
	}

	if getmembers {
		err = g.GetMembers(db, GroupMemberStatus(StatusAccepted), getuser)
		if err != nil {
			return fmt.Errorf("unable to get group members. %v", err)
		}
	}

	return nil
}

// Update updates the group in the database
func (g *Group) Update(db *sql.DB) error {
	// Mux.Lock()
	// defer Mux.Unlock()
	g.UpdatedAt = time.Now()
	query := `UPDATE groups SET title=$1, description=$2, banner_url=$3, updated_at=$4 WHERE id=$5`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(html.EscapeString(g.Title), html.EscapeString(g.Description), html.EscapeString(g.BannerURL), g.UpdatedAt, g.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Delete removes the group from the database
func (g *Group) Delete(db *sql.DB) error {
	query := `UPDATE groups SET deleted_at=$1 WHERE id=$2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), g.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetAllGroups retrieves all groups from the database
func (gs *Groups) GetAllGroups(db *sql.DB, getmembers, getuser bool) error {
	query := `SELECT id, title, description, banner_url, creator_id, created_at, updated_at, deleted_at FROM groups WHERE deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var g Group
		err = rows.Scan(
			&g.ID,
			&g.Title,
			&g.Description,
			&g.BannerURL,
			&g.CreatorID,
			&g.CreatedAt,
			&g.UpdatedAt,
			&g.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		if getmembers {
			err = g.GetMembers(db, GroupMemberStatus(StatusAccepted), getuser)
			if err != nil {
				return fmt.Errorf("unable to get group members. %v", err)
			}
		}

		*gs = append(*gs, g)
	}

	return nil
}

// CreateMember inserts a new member into the group in the database
func (gm *GroupMember) CreateMember(db *sql.DB, memberID, groupID uuid.UUID) error {
	gm.ID = uuid.New()
	gm.GroupID = groupID
	gm.MemberID = memberID
	gm.CreatedAt = time.Now()
	gm.UpdatedAt = time.Now()
	query := `INSERT INTO group_members (id, group_id, member_id, status, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(gm.ID, gm.GroupID, gm.MemberID, gm.Status, gm.Role, gm.CreatedAt, gm.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetMember retrieves a member from the group in the database
func (gm *GroupMember) GetMember(db *sql.DB, memberID, groupID uuid.UUID, getuser bool) error {
	query := `SELECT id, group_id, member_id, status, role, created_at, updated_at, deleted_at FROM group_members WHERE group_id=$1 AND member_id=$2 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	row := stm.QueryRow(groupID, memberID)
	err = row.Scan(
		&gm.ID,
		&gm.GroupID,
		&gm.MemberID,
		&gm.Status,
		&gm.Role,
		&gm.CreatedAt,
		&gm.UpdatedAt,
		&gm.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("unable to scan the row. %v", err)
	}

	if getuser {
		var user = new(User)
		err = user.Get(db, memberID)
		if err != nil {
			return fmt.Errorf("unable to get user. %v", err)
		}
		gm.User = *user
	}

	return nil
}

func (gm *GroupMember) GetMemberById(db *sql.DB, id uuid.UUID, getuser bool) error {
	query := `SELECT id, group_id, member_id, status, role, created_at, updated_at, deleted_at FROM group_members WHERE id=$1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	row := stm.QueryRow(id)
	err = row.Scan(
		&gm.ID,
		&gm.GroupID,
		&gm.MemberID,
		&gm.Status,
		&gm.Role,
		&gm.CreatedAt,
		&gm.UpdatedAt,
		&gm.DeletedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("unable to scan the row. %v", err)
	}

	if getuser {
		var user = new(User)
		err = user.Get(db, gm.MemberID)
		if err != nil {
			return fmt.Errorf("unable to get user. %v", err)
		}
		gm.User = *user
	}

	return nil
}

// UpdateMember updates the member in the group in the database
func (gm *GroupMember) UpdateMember(db *sql.DB) error {
	gm.UpdatedAt = time.Now()
	query := `UPDATE group_members SET status=$1, updated_at=$2 WHERE id=$3`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(gm.Status, gm.UpdatedAt, gm.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// DeleteMember removes the member from the group in the database
func (gm *GroupMember) DeleteMember(db *sql.DB) error {
	query := `UPDATE group_members SET deleted_at=$1 WHERE id=$2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), gm.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetMembers retrieves all members of the group from the database
func (g *Group) GetMembers(db *sql.DB, status GroupMemberStatus, getusers bool) error {
	query := `SELECT id, group_id, member_id, status, role, created_at, updated_at, deleted_at FROM group_members WHERE group_id=$1 AND status=$2 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(g.ID, status)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var gm GroupMember
		err = rows.Scan(
			&gm.ID,
			&gm.GroupID,
			&gm.MemberID,
			&gm.Status,
			&gm.Role,
			&gm.CreatedAt,
			&gm.UpdatedAt,
			&gm.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		if getusers {
			var user = new(User)

			err = user.Get(db, gm.MemberID)
			if err != nil {
				return fmt.Errorf("unable to get user. %v", err)
			}
			gm.User = *user
		}

		g.GroupMembers = append(g.GroupMembers, gm)
	}

	return nil
}

func (mg *GroupMembers) Get(db *sql.DB, id uuid.UUID, status GroupMemberStatus) error {
	query := `SELECT id, group_id, member_id, status, role, created_at, updated_at, deleted_at FROM group_members WHERE group_id=$1 AND status=$2 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(id, status)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var gm GroupMember
		err = rows.Scan(
			&gm.ID,
			&gm.GroupID,
			&gm.MemberID,
			&gm.Status,
			&gm.Role,
			&gm.CreatedAt,
			&gm.UpdatedAt,
			&gm.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		*mg = append(*mg, gm)
	}

	return nil
}

// GetPosts retrieves all posts of the group from the database
func (p *Posts) GetGroupPosts(db *sql.DB, groupID uuid.UUID) error {
	query := `SELECT id, group_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE group_id=$1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(groupID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.GroupID,
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

func (inv *GroupInvitation) SaveInvitation(db *sql.DB, gm GroupMember, invitingUserId uuid.UUID, invitedUserId uuid.UUID) error {
	inv.ID = uuid.New()
	inv.InvitingUserId = invitingUserId
	inv.InvitedUserId = invitedUserId
	inv.GroupMemberId = gm.ID
	inv.CreatedAt = time.Now()
	inv.UpdatedAt = time.Now()
	query := `INSERT INTO invitations (id, inviting_user_id, invited_user_id, group_member_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(inv.ID, inv.InvitingUserId, inv.InvitedUserId, inv.GroupMemberId, inv.CreatedAt, inv.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

type Invitations []struct {
	Group    Group
	MemberId uuid.UUID
}

func (inv *Invitations) GetInvitations(db *sql.DB, userId uuid.UUID) error {
	query := `SELECT groups.*, group_members.id FROM groups JOIN group_members ON groups.id=group_members.group_id JOIN invitations ON group_members.id = invitations.group_member_id where group_members.status="invited" AND invitations.invited_user_id=$1;`
	var memberId uuid.UUID
	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(userId)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var g Group
		err = rows.Scan(
			&g.ID,
			&g.Title,
			&g.Description,
			&g.BannerURL,
			&g.CreatorID,
			&g.CreatedAt,
			&g.UpdatedAt,
			&g.DeletedAt,
			&memberId,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		*inv = append(*inv, struct {
			Group    Group
			MemberId uuid.UUID
		}{Group: g, MemberId: memberId})
	}

	return nil
}
