package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		slog.Error(fmt.Sprintf("Error al leer las varibles de entorno %v", err))
		os.Exit(1)
	}

	app.Get("/testApi", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"menssage": "Go fiber primer api creada",
		})
	})

	if err := app.Listen(os.Getenv("APP_ADDRESS")); err != nil {
		slog.Error(fmt.Sprintf("No se pudo conectar al server %v", err))
		os.Exit(1)
	}
}
