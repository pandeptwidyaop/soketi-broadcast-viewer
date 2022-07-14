package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type WebhookEvent struct {
	Name    string `json:"name"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
	Data    string `json:"data"`
}

type PusherWebhook struct {
	Time   int            `json:"time_ms"`
	Events []WebhookEvent `json:"events"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/webhook", func(c *fiber.Ctx) error {
		body := new(PusherWebhook)

		if err := c.BodyParser(body); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		fmt.Println(body)
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":8000"))
}
