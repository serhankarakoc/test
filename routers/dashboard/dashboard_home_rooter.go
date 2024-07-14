package routers

import (
	handlers "davet.link/handlers/dashboard"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type DashboardHomeRouter struct {
	Store *session.Store
}

func (r *DashboardHomeRouter) SetupRouters(router fiber.Router) {
	router.Get("/dashboard", handlers.DashboardHomeView)
}
