package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/webhook", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":8000"))
}
