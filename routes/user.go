package routes

import (
	"AngelicaRG/encuestasGo/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(g fiber.Router) {
	g.Get("/all", controllers.GetAllUser).Name("api.user.index")
}
