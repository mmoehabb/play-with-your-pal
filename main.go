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
	"goweb/utils/keyboard"
	"goweb/ws"
)

var port = flag.Int("port", 8080, "the port on which the server is listening.")
var password = flag.String("password", "empty", "the password of your session.")
var quality = flag.Int("quality", 15, "the quality of the video stream.")
var noscreen = flag.Bool("noscreen", false, "use this flag to disable sharing your screen.")

func main() {
  flag.Parse()
  if *quality > 100 {
    *quality = 100
  } else if *quality < 1 {
    *quality = 1
  }
  ws.SetConfig(ws.Config{
    Password: *password,
    Quality: *quality,
    Noscreen: *noscreen,
  })
  keyboard.InitKB()

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
