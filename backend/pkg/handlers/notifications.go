// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func handlernotifications(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	notifications := new(models.Notifications)
	if err := notifications.GetByUser(ctx.Db.Conn, userId); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err,
		})
		return
	}
	AllNotification := []map[string]interface{}{}
	for _, notification := range *notifications {

		// member := new(models.GroupMember)
		// is_invite := false

		// if notification.Type == models.TypeGroupInvitation {
		// 	if notification.GroupId != uuid.Nil {
		// 		member.GetMember(ctx.Db.Conn, notification.UserID, notification.GroupId, false)
		// 		notification.UserID = member.MemberID
		// 	} else {
		// 		is_invite = true
		// 	}
		// }
		user := new(models.User)
		user.Get(ctx.Db.Conn, notification.UserID)
		user.Password = ""

		AllNotification = append(AllNotification, map[string]interface{}{
			"id":         notification.ID,
			"type":       notification.Type,
			"concernID":  notification.ConcernID,
			"user":       user,
			"message":    notification.Message,
			"created_at": notification.CreatedAt,
			"group_id":   notification.GroupId,
			"member_id":  notification.MemberId,
			"is_invite":  notification.Is_invite,
		})
	}
	ctx.JSON(AllNotification)
}

func handlerclearnotifications(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	type request struct {
		Type   string `json:"type"`
		Id     string `json:"id"`
		Action string `json:"action"`
	}

	req := new(request)
	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid request",
		})
		return
	}

	if req.Type == "clear" {
		notification := new(models.Notification)
		if err := notification.Get(ctx.Db.Conn, uuid.MustParse(req.Id)); err != nil {
			ctx.Status(http.StatusNotFound).JSON(map[string]interface{}{
				"error": "Notification not found",
			})
			return
		}
		if notification.ConcernID != userId {
			ctx.Status(http.StatusForbidden).JSON(map[string]interface{}{
				"error": "You are not authorized to clear this notification",
			})
			return
		}
		if err := notification.Delete(ctx.Db.Conn); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "something went wrong",
			})
			return
		}
		ctx.JSON(map[string]interface{}{
			"message": "Notification cleared",
		})
		return
	} else if req.Type == "clear_all" {
		notifications := new(models.Notifications)
		if err := notifications.GetByUser(ctx.Db.Conn, userId); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "something went wrong",
			})
			return
		}
		for _, notification := range *notifications {
			if notification.Type == models.TypeNewMessage && req.Action != "new_message" {
				continue
			} else if strings.HasPrefix(string(notification.Type), "follow") && req.Action != "follow" {
				continue
			}
			if err := notification.Delete(ctx.Db.Conn); err != nil {
				ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
					"error": "something went wrong",
				})
				return
			}
		}
		ctx.JSON(map[string]interface{}{
			"message": "Notifications cleared",
		})
		return

	} else {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid request",
		})
		return
	}
}

var notificationsRoute = route{
	path:   "/getnotifications",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handlernotifications,
	},
}

var clearnotificationsRoute = route{
	path:   "/clearnotifications",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handlerclearnotifications,
	},
}

// clearnotifications
func init() {
	AllHandler[notificationsRoute.path] = notificationsRoute
	AllHandler[clearnotificationsRoute.path] = clearnotificationsRoute
}
