package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"net/http"
)

func handleUpload(ctx *octopus.Context) {
	ctx.JSON(map[string]interface{}{
		"imageurl": ctx.Values["file"],
	})
}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var UploadRoute = route{
	path:   "/upload",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		middleware.ImageUploadMiddleware,
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handleUpload, // Handler function to process the authentication request.
	},
}

func init() {
	AllHandler[UploadRoute.path] = UploadRoute
}
