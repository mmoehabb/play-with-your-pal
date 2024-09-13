package ws

import (
	"log"
	"github.com/gofiber/contrib/websocket"
)

var connections []*websocket.Conn
var conn_password = ""
var quality = 15

func RunServer() {
  for {
    for _, conn := range connections {
      conn.WriteMessage(websocket.TextMessage, []byte(CapScreenBase64(quality)))
    }
  }
}

func GetPassword() string {
  return conn_password
}

func SetPassword(password string) {
  conn_password = password
}

func SetQuality(q int) {
  if q > 100 {
    quality = 100
    return
  } 
  if q < 1 {
    quality = 1
    return
  }
  quality = q
}

func AddGuest(c *websocket.Conn, password string) bool {
  if password != conn_password {
    return false
  }
  connections = append(connections, c)
  return true
}

func HandleConn(c *websocket.Conn) error {
  msgType, msg, err := c.ReadMessage()
  if err != nil {
    return err
  }
  go func() {
    if msgType == websocket.TextMessage {
      err := ExecKeyEvent(string(msg))
      if err != nil {
        log.Println(string(msg), err)
      }
    }
  }()
  return nil
}
