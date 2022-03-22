package server

import "github.com/gofiber/fiber/v2"

func Start() {
	app := fiber.New()

	app.Get("/", Versao)

	app.Listen(":3000")
}
