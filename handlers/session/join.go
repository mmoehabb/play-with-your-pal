package session

import (
	"log"
  "strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"goweb/ws"
)

func Join() fiber.Handler {
  return websocket.New(func(c *websocket.Conn) {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
      log.Println("invalid id parameter!")
      return
    }
    ok := ws.JoinSession(c, ws.SessionId(id))
    if ok == false {
      log.Println("couldn't join the session!")
      return
    }
    for {
      if err := ws.HandleConn(ws.SessionId(id), c); err != nil {
        log.Println("guest leaved: ", err)
        return
      }
    }
  })
}
