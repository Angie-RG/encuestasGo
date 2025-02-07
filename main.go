package main

import (
	"fmt"
	"log/slog"
	"os"

	aplication "AngelicaRG/encuestasGo/app"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		slog.Error(fmt.Sprintf("Couldn't read the variables environment %v", err))
		os.Exit(1)
	}

	app.Get("/testApi", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"menssage": "Go fiber primer api creada",
		})
	})

	aplication.Seeders()

	if err := app.Listen(os.Getenv("APP_ADDRESS")); err != nil {
		slog.Error(fmt.Sprintf("Could'nt connection server %v", err))
		os.Exit(1)
	}
}
