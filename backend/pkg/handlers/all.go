package handlers

import (
	octopus "backend/app"
	"net/http"
)

// HandlerConstructor is a type alias for a function that takes a path and a variadic number of HandlerFuncs.
// It's used to define the constructor for creating new routes with associated middleware and handlers.
type HandlerConstructor func(path string, middlewareAndHandler ...octopus.HandlerFunc)

// Handler represents a route with its associated path, constructor, and middleware/handler functions.
type route struct {
	path, method         string
	middlewareAndHandler []octopus.HandlerFunc
}

// AllHandler is a slice of Handler structures that defines all the routes for the application.
// Each Handler in the slice includes the path, the constructor for creating the route, and the middleware/handler functions to be executed.
var AllHandler = map[string]route{}

// HandleAll is a function that iterates over the AllHandler slice and applies each Handler's constructor to register the routes.
// This function should be called during the initialization phase of the application to set up all the routes.
var HandleAll = func(app *octopus.App) {
	var mapConstructors = map[string]HandlerConstructor{
		http.MethodGet:  app.GET,
		http.MethodDelete:  app.DELETE,
		http.MethodPost: app.POST,
		http.MethodPut:  app.PUT,
	}
	for _, v := range AllHandler {
		mapConstructors[v.method](v.path, v.middlewareAndHandler...)
	}
}
