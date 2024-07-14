package routers

import (
	"davet.link/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthRouter struct {
	Store *session.Store
}

func (r *AuthRouter) SetupRouters(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Get("/login", csrf.New(), handlers.LoginPage)
	auth.Post("/login", csrf.New(), handlers.Login(r.Store))
	auth.Get("/register", csrf.New(), handlers.RegisterPage)
	auth.Post("/register", csrf.New(), handlers.Register(r.Store))
	auth.Get("/logout", handlers.Logout(r.Store))
	auth.Get("/forgot-password", csrf.New(), handlers.ForgotPasswordPage)
	auth.Post("/forgot-password", csrf.New(), handlers.ForgotPassword)
	auth.Get("/reset-password/:token", csrf.New(), handlers.ResetPasswordPage)
	auth.Post("/reset-password", csrf.New(), handlers.ResetPassword)
	auth.Get("/verify-email", csrf.New(), handlers.VerifyEmailPage)
	auth.Post("/verify-email", csrf.New(), handlers.VerifyEmail)
	auth.Get("/confirm-account", csrf.New(), handlers.ConfirmAccountPage)
	auth.Post("/confirm-account", csrf.New(), handlers.ConfirmAccount)
}
