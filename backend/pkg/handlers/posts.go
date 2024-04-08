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

func insertPostHandler(ctx *octopus.Context) {
	newPost := models.Post{}
	if err := ctx.BodyParser(&newPost); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while  creating new post",
		})
		return
	}
	userPostOwnerId := ctx.Values["userId"].(uuid.UUID)
	newPost.UserID = userPostOwnerId
	if err := newPost.Create(ctx.Db.Conn); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while creating new post",
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"status": http.StatusOK,
		"data":   newPost.ExploitForRendering(ctx.Db.Conn),
	})
}

func insertCommentHandler(ctx *octopus.Context) {
	newComment := models.Comment{}
	if err := ctx.BodyParser(&newComment); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while  creating new comment",
		})
		return
	}
	newComment.UserID = ctx.Values["userId"].(uuid.UUID)
	post := models.Post{}
	post.Get(ctx.Db.Conn, newComment.PostID)
	fmt.Println(post)
	if err := newComment.Create(ctx.Db.Conn); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while  creating new comment",
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"status": http.StatusOK,
		"data":   newComment.ExploitForRendering(ctx.Db.Conn, string(post.Privacy), post.GroupID),
	})
}

func feedHandler(ctx *octopus.Context) {

	feedPosts := models.Posts{}
	user := ctx.Values["userId"].(uuid.UUID)

	if err := feedPosts.GetAvailablePostForUser(ctx.Db.Conn, user); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while getting posts",
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"status": http.StatusOK,
		"data":   feedPosts.ExploitForRendering(ctx.Db.Conn),
	})
}
func handleGetGroupPost(ctx *octopus.Context) {
	id := ctx.Request.URL.Query().Get("id")
	post := models.Posts{}
	post.GetPostByGroupId(ctx.Db.Conn, id)

	ctx.JSON(map[string]interface{}{
		"status": http.StatusOK,
		"data":   post.ExploitForRendering(ctx.Db.Conn),
	})

}

var getGroupsPostRoute = route{
	path:   "/post/groups",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handleGetGroupPost, // Handler function to process the authentication request.
	},
}

// AuthenticationHandler defines the structure for handling a	uthentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var insertPostRoute = route{
	path:   "/post/insert",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		insertPostHandler, // Handler function to process the authentication request.
	},
}

var getFeedPostsRoute = route{
	path:   "/post/getFeed",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		feedHandler, // Handler function to process the authentication request.
	},
}
var insertCommentRoot = route{
	path:   "/post/insertComment",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		insertCommentHandler, // Handler function to process the authentication request.
	},
}

func init() {
	// Register the authentication route with the global AllHandler map.
	AllHandler[insertCommentRoot.path] = insertCommentRoot
	AllHandler[getFeedPostsRoute.path] = getFeedPostsRoute
	AllHandler[insertPostRoute.path] = insertPostRoute
	AllHandler[getGroupsPostRoute.path] = getGroupsPostRoute
}
