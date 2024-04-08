package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func sendInvitationHandler(ctx *octopus.Context) {
	newMember := models.GroupMember{
		Status: models.MemberStatusInvited,
		Role:   models.MemberRoleUser,
	}

	groupId := ctx.Values["group_id"].(uuid.UUID)
	userId := ctx.Values["userId"].(uuid.UUID)
	invitedUserId := ctx.Values["invited_user_id"].(uuid.UUID)
	err := newMember.CreateMember(ctx.Db.Conn, invitedUserId, groupId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	var invitation models.GroupInvitation

	if err := invitation.SaveInvitation(ctx.Db.Conn, newMember, userId, invitedUserId); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	group := ctx.Values["group"].(*models.Group)
	notification := models.Notification{
		ConcernID: invitedUserId,
		UserID:    userId,
		MemberId:  newMember.ID,
		GroupId:   groupId,
		Is_invite: true,
		Type:      models.TypeGroupInvitation,
		Message:   "initied you to join the group " + group.Title,
	}
	err = notification.Create(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(map[string]interface{}{
		"message": "Invitation sent successfully",
		"data":    newMember,
	})
}

var sendInvitationRoute = route{
	path:   "/send-invitation",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.HaveGroupAccess,
		middleware.IsGroupExist,
		middleware.IsInvitedUserExist,
		sendInvitationHandler,
	},
}

func acceptIntegrationHandler(ctx *octopus.Context) {
	member := ctx.Values["member"].(*models.GroupMember)
	member.Status = models.MemberStatusAccepted
	err := member.UpdateMember(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "Invitation accepted successfully",
		"data":    member,
	})
}

var acceptIntegrationRoute = route{
	path:   "/accept-invitation",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsInvitationExist,
		acceptIntegrationHandler,
	},
}

func declineIntegrationHandler(ctx *octopus.Context) {
	member := ctx.Values["member"].(*models.GroupMember)
	member.Status = models.MemberStatusDeclined
	err := member.UpdateMember(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "Invitation declined successfully",
		"data":    member,
	})
}

var declineIntegrationRoute = route{
	path:   "/decline-invitation",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsInvitationExist,
		declineIntegrationHandler,
	},
}

func demandAccessHandler(ctx *octopus.Context) {
	newMember := models.GroupMember{
		Status: models.MemberStatusRequesting,
		Role:   models.MemberRoleUser,
	}

	group := ctx.Values["group"].(*models.Group)
	requestingUserId := ctx.Values["userId"].(uuid.UUID)
	err := newMember.CreateMember(ctx.Db.Conn, requestingUserId, group.ID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	notification := models.Notification{
		ConcernID: group.CreatorID,
		GroupId:   group.ID,
		MemberId:  newMember.ID,
		Is_invite: false,
		UserID:    requestingUserId,
		Type:      models.TypeGroupInvitation,
		Message:   "A user has requested to join your group",
	}
	err = notification.Create(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "Access demand sent successfully",
		"data":    newMember,
	})
}

var demandAccessRoute = route{
	path:   "/demand-access",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.NoGroupAccess,
		demandAccessHandler,
	},
}

func getAllAccessDemand(ctx *octopus.Context) {
	group := models.Group{
		ID: ctx.Values["group_id"].(uuid.UUID),
	}
	err := group.GetMembers(ctx.Db.Conn, models.MemberStatusRequesting, true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	requestingUsers := group.GroupMembers

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "All access demand",
		"data":    requestingUsers,
	})
}

var getAllAccessDemandRoute = route{
	path:   "/get-all-access-demand",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllAccessDemand,
	},
}

var acceptAccessDemandRoute = route{
	path:   "/accept-access-demand",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupAdmin,
		middleware.IsAccessDemandExist,
		acceptIntegrationHandler,
	},
}

var declineAccessDemandRoute = route{
	path:   "/decline-access-demand",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupAdmin,
		middleware.IsAccessDemandExist,
		declineIntegrationHandler,
	},
}

func getAllInvitations(ctx *octopus.Context) {
	var inv models.Invitations
	err := inv.GetInvitations(ctx.Db.Conn, ctx.Values["userId"].(uuid.UUID))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "All access demand",
		"data":    inv,
	})
}

var getAllInvitationsRoute = route{
	path:   "/get-all-invitations",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		getAllInvitations,
	},
}

func init() {
	AllHandler[sendInvitationRoute.path] = sendInvitationRoute
	AllHandler[acceptIntegrationRoute.path] = acceptIntegrationRoute
	AllHandler[declineIntegrationRoute.path] = declineIntegrationRoute
	AllHandler[demandAccessRoute.path] = demandAccessRoute
	AllHandler[getAllAccessDemandRoute.path] = getAllAccessDemandRoute
	AllHandler[acceptAccessDemandRoute.path] = acceptAccessDemandRoute
	AllHandler[declineAccessDemandRoute.path] = declineAccessDemandRoute
	AllHandler[getAllInvitationsRoute.path] = getAllInvitationsRoute
}
