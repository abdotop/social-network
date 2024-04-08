package config

import "backend/app/session"

var (
	conf = session.Config{CookieName: "sessions"}
	Sess = session.New(&conf)
)
