package handler

import "github.com/gofiber/fiber/v3"

func Teste(c fiber.Ctx) {
	c.JSON("user")

}
