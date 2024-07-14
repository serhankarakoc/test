package routers

import (
	"davet.link/middlewares"
	authRoutes "davet.link/routers/auth"
	dashboardRoutes "davet.link/routers/dashboard"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Router interface {
	SetupRouters(router fiber.Router)
}

func SetupRouters(app *fiber.App, store *session.Store) {
	// Auth middleware uygulanacak route'ları için grup oluşturuyoruz
	authRequiredGroup := app.Group("/", middlewares.AuthRequired(store))

	// Auth middleware uygulanmayacak route'lar için normal bir grup oluşturuyoruz
	publicGroup := app.Group("/")

	// Router'ları tanımlıyoruz
	routersWithAuth := []Router{
		&dashboardRoutes.DashboardHomeRouter{Store: store},
	}

	routersWithoutAuth := []Router{
		&authRoutes.AuthRouter{Store: store},
		&dashboardRoutes.DashboardCategoryRouter{Store: store},
	}

	// Auth middleware uygulanacak router'ları ekliyoruz
	for _, router := range routersWithAuth {
		router.SetupRouters(authRequiredGroup)
	}

	// Auth middleware uygulanmayacak router'ları ekliyoruz
	for _, router := range routersWithoutAuth {
		router.SetupRouters(publicGroup)
	}
}
