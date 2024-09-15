package ws

import (
	"goweb/utils/encoder"
	"goweb/utils/keyboard"
	"goweb/utils/screen"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type Config struct{
  Password string
  Quality int
  Noscreen bool
}
var config Config
func SetConfig(c Config) {
  config = c
}
func GetPassword() string {
  return config.Password
}

type ConnContainer struct {
  Conn *websocket.Conn
  Mu sync.Mutex
}

var session []*ConnContainer
var session_mu sync.Mutex

func RunServer() {
  if config.Noscreen == true {
    return
  }
  for {
    session_mu.Lock()
    frames := screen.CaptureSeq(24, uint8(config.Quality))
    video_buf := encoder.Encode(frames)
    for _, container := range session {
      go sendMsgTo(video_buf, container)
    }
    session_mu.Unlock()
    time.Sleep(1 * time.Second)
  }
}

func sendMsgTo(msg []byte, c *ConnContainer) {
  c.Mu.Lock()
  c.Conn.WriteMessage(websocket.BinaryMessage, msg)
  c.Mu.Unlock()
}

func AddGuest(c *websocket.Conn, password string) bool {
  if password != config.Password {
    return false
  }

  session_mu.Lock()
  session = append(session, &ConnContainer{ Conn: c })
  session_mu.Unlock()

  c.SetCloseHandler(func(code int, text string) error {
    session_mu.Lock()
    var i int
    for ci, container := range session {
      if container.Conn == c {
        i = ci
        break
      }
    }
    session = append(session[:i], session[i+1:]...)
    session_mu.Unlock()
    return nil
  })
  return true
}

func HandleConn(c *websocket.Conn) error {
  msgType, msg, err := c.ReadMessage()
  if err != nil {
    return err
  }
  cmd := strings.Split(string(msg), "_")
  method, key := cmd[0], cmd[1]
  go func() {
    if msgType == websocket.TextMessage {
      err := keyboard.ExecKey(method, key)
      if err != nil {
        log.Println(err)
      }
    }
  }()
  return nil
}
