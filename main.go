package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
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

func InitFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

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

	var client = pusher.Client{
		AppID:  "app",
		Key:    "app",
		Secret: "app",
		Host:   "ws-rnd-btw.iziedev.com",
		Secure: true,
	}

	app.Post("/auth", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		res, err := client.AuthenticatePrivateChannel(c.Body())
		if err != nil {
			return err
		}
		return c.Send(res)
	})

	return app
}

func main() {
	app := InitFiber()
	log.Fatal(app.Listen(":8000"))
}
