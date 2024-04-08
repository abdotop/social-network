package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
)

func handleMessages(ctx *octopus.Context) {
}

func GetUsers(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	AllFollows := new(models.Followers)
	AllFollows.GetAllByFolloweeID(ctx.Db.Conn, userId)
	AllFollows.GetAllByFollowerID(ctx.Db.Conn, userId)

	exist := map[uuid.UUID]bool{}
	AllUser := []map[string]interface{}{}
	for _, follow := range *AllFollows {
		user := new(models.User)
		id := follow.FolloweeID
		if follow.FolloweeID == userId {
			id = follow.FollowerID
		}
		if _, ok := exist[id]; ok {
			continue
		}
		exist[id] = true
		if user.Get(ctx.Db.Conn, id) != nil {
			ctx.JSON(map[string]interface{}{
				"error": "error geting users.",
			})
		}
		lastMessage := new(models.PrivateMessage)
		lastMessage.GetLastMessage(ctx.Db.Conn, userId, id)
		user.Password = ""
		AllUser = append(AllUser, map[string]interface{}{
			"user":        user,
			"lastMessage": lastMessage,
		})
	}

	ctx.JSON(map[string]interface{}{
		"data": AllUser,
	})
}
func handlerGetMessages(ctx *octopus.Context) {
	var senderId = ctx.Values["userId"].(uuid.UUID)
	var messages = new(models.PrivateMessages)
	var receiverId map[string]string
	if err := ctx.BodyParser(&receiverId); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{"message": "bad request"})
		return
	}
	err1 := messages.GetPrivateMessages(ctx.Db.Conn, uuid.MustParse(receiverId["receiver_id"]), senderId)
	if err1 != nil {
		// HandleError(ctx.ResponseWriter, http.StatusInternalServerError, "Error getting users : "+err1.Error())
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{"message": "bad request"})
		return
	}
	data := map[string]interface{}{
		"status": http.StatusOK,
		"data":   messages,
	}
	ctx.JSON(data)
	// HandleError(ctx.ResponseWriter, http.StatusUnauthorized, "No active session")
}

var messagesRoutes = route{
	path:   "/groups/messages",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handleMessages,
	},
}
var getUsers = route{
	path:   "/usersByFollow",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		GetUsers,                // Handler function to process the messages request.
	},
}
var getMessages = route{
	path:   "/getMessages",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handlerGetMessages,      // Handler function to process the messages request.
	},
}

func init() {
	// Register the events route with the global AllHandler map.
	AllHandler[getUsers.path] = getUsers
	AllHandler[getMessages.path] = getMessages
}
