package handlers

import (
	"time"

	"davet.link/dtos"
	"davet.link/repositories"
	"davet.link/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func authService() *services.AuthService {
	repo := repositories.NewAuthRepository()
	return services.NewAuthService(repo)
}

func LoginPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/login", fiber.Map{
		"Title": "Giriş Yap",
	})
}

func Login(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto dtos.LoginDTO
		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
		}

		service := authService()

		var rememberToken string
		if dto.RememberToken != nil {
			rememberToken = *dto.RememberToken
		}

		user, err := service.Authenticate(dto.Email, dto.Password, rememberToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Email veya şifre yanlış: " + err.Error())
		}

		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		sess.Set("userID", user.ID)
		if user.RememberToken != nil && *user.RememberToken != "" {
			c.Cookie(&fiber.Cookie{
				Name:     "remember_token",
				Value:    *user.RememberToken,
				Expires:  time.Now().Add(30 * 24 * time.Hour),
				HTTPOnly: true,
			})
		}

		if err := sess.Save(); err != nil {
			return err
		}

		return c.Redirect("/")
	}
}

func Logout(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		sess.Destroy()
		c.ClearCookie("remember_token")
		return c.Redirect("/auth/login")
	}
}

func RegisterPage(c *fiber.Ctx) error {
	csrfToken, ok := c.Locals("csrf").(string)
	if !ok || csrfToken == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("CSRF token not found")
	}
	return c.Render("pages/auth/register", fiber.Map{
		"Title": "Kayıt Ol",
		"CSRF":  csrfToken,
	})
}

func Register(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto dtos.RegisterDTO
		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
		}

		service := authService()
		if err := service.Register(dto); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Kayıt başarısız: " + err.Error())
		}

		return c.Redirect("/auth/login")
	}
}

func ForgotPasswordPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/forgot_password", fiber.Map{
		"Title": "Şifremi Unuttum",
	})
}

func ForgotPassword(c *fiber.Ctx) error {
	var dto dtos.ForgotPasswordDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	service := authService()
	if err := service.SendPasswordResetEmail(dto.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send password reset email: " + err.Error())
	}

	return c.SendString("Şifre sıfırlama talimatları email adresinize gönderildi.")
}

func ResetPasswordPage(c *fiber.Ctx) error {
	token := c.Params("token")
	return c.Render("pages/auth/reset_password", fiber.Map{
		"Title": "Şifreyi Sıfırla",
		"Token": token,
	})
}

func ResetPassword(c *fiber.Ctx) error {
	var dto dtos.ResetPasswordDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	service := authService()
	if err := service.ResetPassword(dto.Token, dto.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to reset password: " + err.Error())
	}

	return c.Redirect("/auth/login")
}

func VerifyEmailPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/verify_email", fiber.Map{
		"Title": "Email Doğrulama",
	})
}

func VerifyEmail(c *fiber.Ctx) error {
	var dto dtos.VerifyEmailDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	service := authService()
	if err := service.SendVerificationEmail(dto.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send verification email: " + err.Error())
	}

	return c.SendString("Doğrulama emaili gönderildi.")
}

func ConfirmAccountPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/confirm_account", fiber.Map{
		"Title": "Hesap Doğrulama",
	})
}

func ConfirmAccount(c *fiber.Ctx) error {
	var dto dtos.ConfirmAccountDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	service := authService()
	if err := service.ConfirmAccount(dto.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to confirm account: " + err.Error())
	}

	return c.Redirect("/auth/login")
}
