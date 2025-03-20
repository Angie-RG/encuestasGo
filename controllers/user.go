package controllers

import (
	"AngelicaRG/encuestasGo/app"
	"AngelicaRG/encuestasGo/models"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	users := []models.User{}
	if err := app.DB().Model(&models.User{}).Select("id, first_name, last_name, email, active").
		Where("deleted_at is null").
		Find(&users).Error; err != nil {
		slog.Error(fmt.Sprintf("Error al obtener todos los usuarios %v.", err))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": users})
}
