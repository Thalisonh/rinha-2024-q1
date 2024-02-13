package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func Init() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("✋")
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":8080"))
}
