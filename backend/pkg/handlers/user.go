package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type updateValues struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

type avatarUpdate struct {
	AvatarImage string `json:"avatarImage"`
	Email       string `json:"email"`
}

func handleUpdateUser(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}
	user.ID = userId
	if err := user.Validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	newHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while hashing the password.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	user.Password = string(newHash)
	if err := user.Update(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "User updated successfully",
		"status":  http.StatusOK,
	})
}

func handleGetUser(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	type request struct {
		Action   string `json:"action"`
		Nickname string `json:"nickname"`
	}
	req := new(request)
	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}
	switch req.Action {
	case "get":
		user := new(models.User)
		if req.Nickname == "" {
			if err := user.Get(ctx.Db.Conn, userId); err != nil {
				ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
					"message": err.Error(),
					"status":  http.StatusInternalServerError,
				})
				return
			}
		} else {
			if err := user.Get(ctx.Db.Conn, req.Nickname); err != nil {
				ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
					"message": err.Error(),
					"status":  http.StatusInternalServerError,
				})
				return
			}
		}

		follower := new(models.Follower)
		follower.FollowerID = userId
		follower.FolloweeID = user.ID
		follower.Get(ctx.Db.Conn)
		if follower.Status == "" {
			follower.Status = "none"
		}
		if userId == user.ID {
			follower.Status = "self"
		}

		follow := new(models.Followers).CountAllByFollowerID(ctx.Db.Conn, userId)
		following := new(models.Followers).CountAllByFolloweeID(ctx.Db.Conn, userId)
		numpost, _ := models.CountPostsByUser(ctx.Db.Conn, user.ID)
		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "User fetched successfully",
			"status":  http.StatusOK,
			"data": map[string]interface{}{
				"id":            user.ID,
				"firstname":     user.FirstName,
				"lastname":      user.LastName,
				"email":         user.Email,
				"nickname":      user.Nickname,
				"birthday":      user.DateOfBirth,
				"about":         user.AboutMe,
				"avatar":        user.AvatarImage,
				"created":       user.CreatedAt,
				"updated":       user.UpdatedAt,
				"follow":        follow,
				"following":     following,
				"followStatus":  follower.Status,
				"NumberOfPosts": numpost,
			},
		})
	case "posts":
		posts := new(models.Posts)
		if err := posts.GetUserPosts(ctx.Db.Conn, userId); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "User posts fetched successfully",
			"status":  http.StatusOK,
			"posts":   posts,
		})
	default:
		ctx.JSON(map[string]interface{}{
			"message": "Invalid action.",
			"status":  http.StatusBadRequest,
		})
	}
}

func handleUpdateUserInfos(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}
	user.ID = userId
	if err := user.Validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	if err := user.Update(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	// newUser := models.User{}

	// err := newUser.Get(ctx.Db.Conn, user.Email, true)
	// if err != nil {
	// 	ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
	// 		"session": "",
	// 		"message": "User not found.",
	// 		"status":  http.StatusUnauthorized,
	// 		"data":    nil,
	// 	})
	// 	return
	// }
	// idSession, err := config.Sess.Start(ctx).Set(user.ID)
	// // Start a new session for the user and set the user's ID as the session key.
	// if err != nil {
	// 	// If starting the session fails, log the error and return an HTTP status  500.
	// 	ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
	// 		"session": "",
	// 		"message": "Error while starting the session.",
	// 		"status":  http.StatusInternalServerError,
	// 	})
	// 	return
	// }
	data := map[string]interface{}{
		"message": "User informations updated successfully.",
		"session": "", //idSession,
		"data":    user,
		"status":  http.StatusOK,
	}
	ctx.JSON(data)
}

func handleUpdateUserPassword(ctx *octopus.Context) {
	// Log the client's IP address that reached the login route.
	var newCredentials = updateValues{}
	// Try to deserialize the form data into the User instance.
	if err := ctx.BodyParser(&newCredentials); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while parsing the form data.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	var credentials = credentials{
		Email:    newCredentials.Email,
		Password: newCredentials.Password,
	}
	if err := credentials.Validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}
	newUser := models.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}
	err := newUser.Get(ctx.Db.Conn, credentials.Email, true)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "invalid email.",
			"status":  http.StatusUnauthorized,
		})
		return
	}
	// Check if the user's credentials are valid.
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(credentials.Password)); err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "Invalid credentials. Please try again.",
			"status":  http.StatusUnauthorized,
		})
		return
	}
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newCredentials.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while hashing the password.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	newUser.Password = string(newPasswordHash)
	// Attempts to update the user in the database with the provided data.
	if newUser.Update(ctx.Db.Conn) != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while creating the user.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	// Starts a new session for the user and sets the user's ID as the session key.
	// idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	// if err != nil {
	// 	// If starting the session fails, logs the error and returns an HTTP status  500.
	// 	ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
	// 		"session": "",
	// 		"message": "Error while starting the session.",
	// 		"status":  http.StatusInternalServerError,
	// 	})
	// 	return
	// }
	ctx.Status(http.StatusAccepted).JSON(map[string]interface{}{
		"session": "", //idSession,
		"data":    newUser,
		"message": "User password successfully updated.",
		"status":  "200",
	})
}

func handleUpdateAvatar(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)
	user := new(avatarUpdate)
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	currentUser := new(models.User)
	currentUser.ID = userId
	err := currentUser.Get(ctx.Db.Conn, user.Email, true)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "invalid email.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	currentUser.AvatarImage = user.AvatarImage
	if err := currentUser.Update(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "User Avatar updated successfully",
		"status":  http.StatusOK,
	})

}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var updateUserRoute = route{
	path:   "/updateuser",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleUpdateUser,        // Handler function to process the authentication request.
	},
}

var getUserRoute = route{
	path:   "/getuser",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleGetUser,           // Handler function to process the authentication request.
	},
}

var updateUserInfosRoute = route{
	path:   "/edituser",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleUpdateUserInfos,   // Handler function to process the user informations update request.
	},
}

var updateUserPasswordRoute = route{
	path:   "/updatepassword",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,  // Middleware to check if the request is authenticated.
		handleUpdateUserPassword, // Handler function to process the update the user password request.
	},
}

var updateAvatarRoute = route{
	path:   "/updateavatar",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handleUpdateAvatar,
	},
}

func init() {
	AllHandler[updateUserRoute.path] = updateUserRoute
	AllHandler[getUserRoute.path] = getUserRoute
	AllHandler[updateUserInfosRoute.path] = updateUserInfosRoute
	AllHandler[updateUserPasswordRoute.path] = updateUserPasswordRoute
	AllHandler[updateAvatarRoute.path] = updateAvatarRoute
}
