package handlers

import (
	"strconv"
	"time"

	"davet.link/dtos"
	"davet.link/repositories"
	"davet.link/services"

	"github.com/gofiber/fiber/v2"
)

func userService() *services.UserService {
	repository := repositories.NewUserRepository()
	return services.NewUserService(repository)
}

func GetAllUsers(c *fiber.Ctx) error {
	service := userService()

	users, err := service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get users: " + err.Error())
	}

	return c.Render("users/list", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Listesi",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Users",
		"activePage":      "Users-List",
		"Users":           users,
	})
}

func CreateUserView(c *fiber.Ctx) error {
	csrfToken := c.Locals("csrf").(string)
	return c.Render("users/create", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Ekle",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Users",
		"activePage":      "Users-Create",
		"csrfToken":       csrfToken,
	})
}

func CreateUser(c *fiber.Ctx) error {
	service := userService()

	var dto dtos.UserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	if err := services.ValidateUserCreate(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Validation failed: " + err.Error())
	}

	if _, err := service.CreateUser(dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user: " + err.Error())
	}

	return c.Redirect("/dashboard/users")
}

func EditUserView(c *fiber.Ctx) error {
	service := userService()

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID: " + err.Error())
	}

	user, err := service.GetUserByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user: " + err.Error())
	}

	csrfToken := c.Locals("csrf").(string)
	return c.Render("users/edit", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Düzenle",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Users",
		"activePage":      "Users-List",
		"User":            user,
		"csrfToken":       csrfToken,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	service := userService()

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID: " + err.Error())
	}

	var dto dtos.UserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	if err := services.ValidateUserUpdate(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Validation failed: " + err.Error())
	}

	user, err := service.UpdateUser(uint(userID), dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update user: " + err.Error())
	}

	return c.Render("users/edit", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Düzenle",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Users",
		"activePage":      "Users-List",
		"User":            user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	service := userService()

	userIDParam := c.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID: " + err.Error())
	}

	if err := service.DeleteUser(uint(userID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete user: " + err.Error())
	}

	return c.Redirect("/dashboard/users")
}
