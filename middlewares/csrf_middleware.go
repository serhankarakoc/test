package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func CsrfMiddleware() fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      "form:_csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		CookieSecure:   false,
	})
}
