package ws

import (
	"log"
	"github.com/gofiber/contrib/websocket"
)

var connections []*websocket.Conn
var conn_password = ""

func GetPassword() string {
  return conn_password
}

func SetPassword(password string) {
  conn_password = password
}

func HandleConn(c *websocket.Conn) error {
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

func AddGuest(c *websocket.Conn, password string) bool {
  if password != conn_password {
    return false
  }
  connections = append(connections, c)
  return true
}
