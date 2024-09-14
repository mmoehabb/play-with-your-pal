package handlers

import (
	"context"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"goweb/ui/components"
	"goweb/ws"
)

type Credentials struct {
  Password string `json:"password" xml:"password" form:"password"`
}

func Join(c *fiber.Ctx) error {
  creds := new(Credentials)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  if err:= c.BodyParser(creds); err != nil {
    return c.SendStatus(fiber.StatusBadRequest)
  }
  if creds.Password != ws.GetPassword() {
    return c.SendStatus(fiber.StatusUnauthorized)
  } 
  components.Stream().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}

func WSJoin() fiber.Handler {
  return websocket.New(func(c *websocket.Conn) {
    password := c.Params("password")
    ok := ws.AddGuest(c, password)
    if ok == false {
      log.Println("couldn't join the session!")
      return
    }
    log.Println("new guest has joined the session!")
    for {
      if err := ws.HandleConn(c); err != nil {
        log.Println("guest left: ", err)
        return
      }
    }
  })
}
