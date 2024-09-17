package ws

import (
	"goweb/utils/encoder"
	"goweb/utils/keyboard"
	"goweb/utils/screen"
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

var framesChannel = make(chan screen.Frame, 24)
var videoBufChannel = make(chan []byte)

func RunServer() {
  if config.Noscreen == true {
    return
  }
  go screenLoop()
  go videoBufLoop()
  for {
    session_mu.Lock()
    buf := <- videoBufChannel
    for _, container := range session {
      go sendMsgTo(buf, container)
    }
    session_mu.Unlock()
  }
}

func screenLoop() {
  for {
    frame, err := screen.Capture(uint8(config.Quality))
    if err != nil {
      panic(err)
    }
    framesChannel <- frame
  }
}

var collectedFrames []screen.Frame
func videoBufLoop() {
  t := time.Now()
  for {
    collectedFrames = append(collectedFrames, <-framesChannel)
    if time.Since(t) < 20 * time.Millisecond {
      continue
    }
    videoBufChannel <- encoder.Encode(collectedFrames)
    collectedFrames = []screen.Frame{}
    t = time.Now()
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
      keyboard.ExecKey(method, key)
    }
  }()
  return nil
}
