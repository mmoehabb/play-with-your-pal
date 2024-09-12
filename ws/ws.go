package ws

import (
	"log"
	"math/rand"

	"github.com/gofiber/contrib/websocket"
)

type Session struct{
  host *websocket.Conn
  guest *websocket.Conn
}
type SessionId int32

var sessions = make(map[SessionId]*Session)

func HandleConn(id SessionId, c *websocket.Conn) error {
  msgType, msg, err := c.ReadMessage()
  if err != nil {
    return err
  }
  go func() {
    if msgType == websocket.TextMessage {
      log.Println("key clicked: ", string(msg))
    }
  }()
  return nil
}

func CreateSession(c *websocket.Conn) SessionId {
  id := SessionId(rand.Intn(9999))
  for sessions[id] != nil {
    id = SessionId(rand.Intn(9999))
  }
  sessions[id] = &Session{
    host: c,
  }
  c.SetCloseHandler(func(code int, text string) error {
    CloseSession(id)
    return nil
  })
  return id
}

func JoinSession(c *websocket.Conn, id SessionId) bool {
  if sessions[id] == nil || sessions[id].guest != nil {
    return false
  }
  sessions[id].guest = c
  return true
}

func KickGuest(id SessionId) {
  sessions[id].guest.Close()
}

func CloseSession(id SessionId) {
  sessions[id].host.Close()
  if sessions[id].guest != nil {
    sessions[id].guest.Close()
  }
  delete(sessions, id)
}
