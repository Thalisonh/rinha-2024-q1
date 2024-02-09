package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func Init() {
	app := fiber.New()

	app.Get("/api/*", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":3000"))
}
