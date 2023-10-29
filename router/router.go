package router

import (
	"github.com/gofiber/fiber/v2"

	"Login-api/controllers"
)

func Register(app *fiber.App) {

	dhcp := app.Group("/login")
	dhcp.Post("/token", controllers.Logar)
	dhcp.Post("/create", controllers.Create)
	dhcp.Delete("/delete", controllers.Delete)
	dhcp.Post("/validador", controllers.VerificaJwt)
}