package configs

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func SetupSession() {
	Store = session.New(session.Config{
		KeyLookup:      "cookie:session_id",
		CookieHTTPOnly: true,
		CookieSecure:   false,
		CookieSameSite: "Strict",
	})
}
