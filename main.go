package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"goweb/handlers"
	"goweb/pages"
	"goweb/ws"
)

var port = flag.Int("port", 8080, "the port on which the server is listening.")
var password = flag.String("password", "empty", "the password of your session.")
var quality = flag.Int("quality", 5, "the quality of the video stream.")

func main() {
  flag.Parse()
  ws.SetPassword(*password)
  ws.SetQuality(*quality)
  app := fiber.New()
  app.Static("/public", "./public/")

  app.Get("/", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    pages.Index().Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(200)
  })
  app.Post("/join", handlers.Join)
  
  app.Use(func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
  app.Get("/ws/join/:password", handlers.WSJoin())

  go ws.RunServer()
  log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
