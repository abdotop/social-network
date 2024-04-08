package middleware

import (
	octopus "backend/app"
	"backend/pkg/config"
	"net/http"
	"os"
	"strings"
)

// AuthRequired verify if the user is connected
func AuthRequired(ctx *octopus.Context) {
	var token string
	headerBearer := ctx.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	if !config.Sess.Start(ctx).Valid(token) {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	userId, err := config.Sess.Start(ctx).Get(token)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}

	ctx.Values["userId"] = userId
	ctx.Values["token"] = token
	ctx.Next()
}

// NoAuthRequired verify if the user is not connected
func NoAuthRequired(ctx *octopus.Context) {
	var token string
	headerBearer := ctx.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	if config.Sess.Start(ctx).Valid(token) {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "You already have session",
		})
		return
	}
	ctx.Next()
}

func AllowedSever(ctx *octopus.Context) {
	key := ctx.Request.URL.Query().Get("key")
	if key != os.Getenv("SERVER_KEY") {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "You are not allowed to access this server socket.",
		})
		return
	}
	ctx.Next()
}
