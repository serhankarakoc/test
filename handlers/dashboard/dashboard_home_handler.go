package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func DashboardHomeView(c *fiber.Ctx) error {
	data := fiber.Map{
		"home":            true,
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Ana Sayfa",
		"breadcrumbTitle": nil,
		"activeGroup":     nil,
		"activePage":      "Home",
	}
	return c.Render("pages/dashboard", data)
}
