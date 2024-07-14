package middlewares

import (
	"davet.link/configs"
	"davet.link/dtos"
	"davet.link/models"
	"github.com/gofiber/fiber/v2"
)

func ContextMiddleware(c *fiber.Ctx) error {
	var globalVars []models.Global
	if err := configs.DB.Find(&globalVars).Error; err != nil {
		return err
	}

	for _, globalVar := range globalVars {
		globalDTO := dtos.GlobalDTO{
			Key:   globalVar.Key,
			Value: globalVar.Value,
		}
		c.Locals(globalDTO.Key, globalDTO.Value)
	}

	return c.Next()
}
