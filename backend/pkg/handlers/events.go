// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func createEventHandler(ctx *octopus.Context) {
	newEvent := models.Event{}

	if err := ctx.BodyParser(&newEvent); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newEvent.CreatorID = ctx.Values["userId"].(uuid.UUID)
	newEvent.GroupID = ctx.Values["group_id"].(uuid.UUID)
	if err := newEvent.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	members := new(models.GroupMembers)

	if err := members.Get(ctx.Db.Conn, newEvent.GroupID, models.MemberStatusAccepted); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	group := ctx.Values["group"].(*models.Group)

	for _, member := range *members {
		log.Println(member.ID)
		notif := &models.Notification{
			UserID:    newEvent.CreatorID,
			ConcernID: member.MemberID,
			MemberId:  member.ID,
			GroupId:   newEvent.GroupID,
			Message:   fmt.Sprintf("add new event in the group %s", group.Title),
			Type:      models.TypeNewEvent,
		}

		if err := notif.Create(ctx.Db.Conn); err != nil {
			log.Println(err)
			continue
		}
		log.Println(notif.ID)
	}

	ctx.Status(http.StatusCreated).JSON(map[string]interface{}{
		"message": "Event created successfully",
		"data":    newEvent,
	})
}

var createEventRoute = route{
	path:   "/create-event",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		createEventHandler,
	},
}

func getAllEventByGroup(ctx *octopus.Context) {
	events := models.Events{}
	groupId := ctx.Values["group_id"].(uuid.UUID)
	isParticipantNeeded := ctx.Request.URL.Query().Get("isParticipantNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := events.GetGroupEvents(ctx.Db.Conn, groupId, isParticipantNeeded, isUserNeeded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "All events",
		"data":    events,
	})
}

var getAllEventRoute = route{
	path:   "/get-all-event-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllEventByGroup,
	},
}

func respondEventHandler(ctx *octopus.Context) {
	event := ctx.Values["event"].(*models.Event)
	member := ctx.Values["member"].(*models.GroupMember)

	participant := models.EventParticipant{}
	_participant := models.EventParticipant{}
	if err := ctx.BodyParser(&_participant); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if _participant.Response != "not_going" && _participant.Response != "going" {
		ctx.Status(http.StatusBadRequest).JSON("Invalid response")
		return
	}
	err := participant.GetParticipant(ctx.Db.Conn, event.ID, member.ID, member.MemberID, false)
	participant.Response = _participant.Response
	if err != nil {
		err := participant.CreateParticipant(ctx.Db.Conn, event.ID, member.ID)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	} else {
		err := participant.UpdateParticipant(ctx.Db.Conn)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "Response updated",
		"data":    participant,
	})
}

var respondEventRoute = route{
	path:   "/response-event",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsEventExist,
		respondEventHandler,
	},
}

func init() {
	AllHandler[createEventRoute.path] = createEventRoute
	AllHandler[getAllEventRoute.path] = getAllEventRoute
	AllHandler[respondEventRoute.path] = respondEventRoute
}
