package middlewares

import (
	"davet.link/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.uber.org/zap"
)

func AuthRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := utils.GetLogger()

		isAuthenticated, err := utils.IsAuthenticated(c, store)
		if err != nil {
			logger.Warn("Failed authentication check", zap.Error(err))
			return c.Redirect("/auth/login")
		}

		if !isAuthenticated {
			logger.Warn("Unauthenticated access attempt", zap.String("path", c.Path()))
			return c.Redirect("/auth/login")
		}

		userID, err := utils.GetUserID(c, store)
		if err != nil {
			logger.Warn("Failed to get user ID", zap.Error(err))
			return c.Redirect("/auth/login")
		}

		logger.Info("Authenticated access", zap.String("path", c.Path()), zap.Uint("user_id", userID))
		return c.Next()
	}
}
