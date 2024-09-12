package main

import (
	"context"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"goweb/handlers/session"
	"goweb/pages"
)

func main() {
  app := fiber.New()
  app.Static("/public", "./public/")

  app.Get("/", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    pages.Index().Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(200)
  })
  
  app.Use(func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

  app.Get("/ws/host", session.Host())
  app.Get("/ws/join/:id<int>", session.Join())

  app.Listen(":3000")
}
