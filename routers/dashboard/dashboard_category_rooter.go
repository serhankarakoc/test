package routers

import (
	dashboard "davet.link/handlers/dashboard"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type DashboardCategoryRouter struct {
	Store *session.Store
}

func (r *DashboardCategoryRouter) SetupRouters(router fiber.Router) {
	categories := router.Group("/dashboard/categories")
	categories.Get("/", dashboard.GetAllCategories)
	categories.Get("/create", dashboard.CreateCategoryView)
	categories.Post("/create", dashboard.CreateCategory)
	categories.Get("/edit/:id", dashboard.EditCategoryView)
	categories.Put("/edit/:id", dashboard.UpdateCategory)
	categories.Delete("/delete/:id", dashboard.DeleteCategory)
}
