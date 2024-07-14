package handlers

import (
	"strconv"
	"time"

	"davet.link/dtos"
	"davet.link/repositories"
	"davet.link/services"

	"github.com/gofiber/fiber/v2"
)

func categoryService() *services.CategoryService {
	repository := repositories.NewCategoryRepository()
	return services.NewCategoryService(repository)
}

func GetAllCategories(c *fiber.Ctx) error {
	service := categoryService()

	categories, err := service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get categories: " + err.Error())
	}

	return c.Render("categories/list", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Listesi",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Categories",
		"activePage":      "Categories-List",
		"Categories":      categories,
	})
}

func CreateCategoryView(c *fiber.Ctx) error {
	csrfToken := c.Locals("csrf").(string)
	return c.Render("categories/create", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Ekle",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Categories",
		"activePage":      "Categories-Create",
		"csrfToken":       csrfToken,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	service := categoryService()

	var dto dtos.CategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	if err := services.ValidateCategoryCreate(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Validation failed: " + err.Error())
	}

	if _, err := service.CreateCategory(dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create category: " + err.Error())
	}

	return c.Redirect("/dashboard/categories")
}

func EditCategoryView(c *fiber.Ctx) error {
	service := categoryService()

	categoryIDParam := c.Params("id")
	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid category ID: " + err.Error())
	}

	category, err := service.GetCategoryByID(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get category: " + err.Error())
	}

	csrfToken := c.Locals("csrf").(string)
	return c.Render("categories/edit", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Düzenle",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Categories",
		"activePage":      "Categories-List",
		"Category":        category,
		"csrfToken":       csrfToken,
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	service := categoryService()

	categoryIDParam := c.Params("id")
	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid category ID: " + err.Error())
	}

	var dto dtos.CategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	if err := services.ValidateCategoryUpdate(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Validation failed: " + err.Error())
	}

	category, err := service.UpdateCategory(uint(categoryID), dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update category: " + err.Error())
	}

	return c.Render("categories/edit", fiber.Map{
		"Year":            time.Now().Year(),
		"root":            "dashboard",
		"pageTitle":       "Kategori Düzenle",
		"breadcrumbTitle": "Kategoriler",
		"activeGroup":     "Categories",
		"activePage":      "Categories-List",
		"Category":        category,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	service := categoryService()

	categoryIDParam := c.Params("id")
	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid category ID: " + err.Error())
	}

	if err := service.DeleteCategory(uint(categoryID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete category: " + err.Error())
	}

	return c.Redirect("/dashboard/categories")
}
