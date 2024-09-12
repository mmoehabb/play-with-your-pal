package session

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"goweb/ws"
)

func Host() fiber.Handler {
  return websocket.New(func(c *websocket.Conn) {
    sessionId := ws.CreateSession(c)
    data, err := json.Marshal(map[string]int{"session_id": int(sessionId)})
    if err != nil {
      panic("json.Marshal error: " + err.Error())
    }
    c.WriteMessage(websocket.TextMessage, data)
    for {
      if err := ws.HandleConn(sessionId, c); err != nil {
        log.Println("session ended: ", err)
        return
      }
    }
  })
}

