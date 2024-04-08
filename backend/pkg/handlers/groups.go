// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func createGroup(ctx *octopus.Context) {
	newGroup := models.Group{}

	if err := ctx.BodyParser(&newGroup); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newGroup.CreatorID = ctx.Values["userId"].(uuid.UUID)
	if err := newGroup.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(map[string]interface{}{
		"message": "Group created successfully",
		"data":    newGroup,
	})
}

var createGroupRoute = route{
	path:   "/create-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupValid,
		createGroup,
	},
}

func getAllGroups(ctx *octopus.Context) {
	groups := models.Groups{}
	isMemberNeeded := ctx.Request.URL.Query().Get("isMemberNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := groups.GetAllGroups(ctx.Db.Conn, isMemberNeeded, isUserNeeded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(map[string]interface{}{
		"message": "All groups",
		"data":    groups,
	})
}

var getAllGroupsRoute = route{
	path:   "/get-all-groups",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		getAllGroups,
	},
}

func getGroupById(ctx *octopus.Context) {
	group := models.Group{}
	groupID := ctx.Values["group_id"].(uuid.UUID)
	isMemberNeeded := ctx.Request.URL.Query().Get("isMemberNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := group.Get(ctx.Db.Conn, groupID, isMemberNeeded, isUserNeeded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(map[string]interface{}{
		"message": "Group",
		"data":    group,
	})
}

var getGroupByIdRoute = route{
	path:   "/get-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		getGroupById,
	},
}

func createPostGroup(ctx *octopus.Context) {
	post := models.Post{}
	post.GroupID = ctx.Values["group_id"].(uuid.UUID)
	post.UserID = ctx.Values["userId"].(uuid.UUID)

	if err := ctx.BodyParser(&post); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(post, err.Error())
		return
	}

	post.Privacy = "group"

	if err := post.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(map[string]interface{}{
		"message": "Post created successfully",
		"data": post.ExploitForRendering(ctx.Db.Conn),
	})
}

var createPostGroupRoute = route{
	path:   "/create-post-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupPostValid,
		createPostGroup,
	},
}

func getAllGroupPosts(ctx *octopus.Context) {
	posts := models.Posts{}
	groupID := ctx.Values["group_id"].(uuid.UUID)
	err := posts.GetGroupPosts(ctx.Db.Conn, groupID)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(map[string]interface{}{
		"data": posts.ExploitForRendering(ctx.Db.Conn),
	})
}

var getAllGroupPostsRoute = route{
	path:   "/get-all-post-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllGroupPosts,
	},
}

func getAllGroupMessages(ctx *octopus.Context) {
	groupID := ctx.Values["group_id"].(uuid.UUID)
	messages := models.GroupMessages{}
	err := messages.GetGroupMessages(ctx.Db.Conn, groupID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(map[string]interface{}{
		"message": "All messages",
		"data":    messages,
	})
}

var getAllGroupMessagesRoute = route{
	path:   "/group/messages",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllGroupMessages,
	},
}

func addNewGroupMessage(ctx *octopus.Context) {
	newMessage := models.GroupMessage{}
	if err := ctx.BodyParser(&newMessage); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newMessage.GroupID = ctx.Values["group_id"].(uuid.UUID)
	newMessage.SenderID = ctx.Values["userId"].(uuid.UUID)
	if err := newMessage.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(map[string]interface{}{
		"message": "Message sent",
		"data":    newMessage,
	})
}

var addNewGroupMessageRoute = route{
	path:   "/group/messages/new",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		addNewGroupMessage,
	},
}

func init() {
	AllHandler[createGroupRoute.path] = createGroupRoute
	AllHandler[getAllGroupsRoute.path] = getAllGroupsRoute
	AllHandler[getGroupByIdRoute.path] = getGroupByIdRoute
	AllHandler[createPostGroupRoute.path] = createPostGroupRoute
	AllHandler[getAllGroupPostsRoute.path] = getAllGroupPostsRoute
	AllHandler[getAllGroupMessagesRoute.path] = getAllGroupMessagesRoute
	AllHandler[addNewGroupMessageRoute.path] = addNewGroupMessageRoute
}
