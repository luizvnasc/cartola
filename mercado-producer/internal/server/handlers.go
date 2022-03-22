package server

import "github.com/gofiber/fiber/v2"

func Versao(c *fiber.Ctx) error {
	return c.SendString("v0.0.1")
}
