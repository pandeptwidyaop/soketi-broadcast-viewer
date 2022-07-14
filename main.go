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
		var body map[string]string

		if err := c.BodyParser(body); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		fmt.Println(body)
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":8000"))
}
