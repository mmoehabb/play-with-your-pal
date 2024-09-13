package ws

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"image/jpeg"
	"runtime"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/micmonay/keybd_event"
)

func CapScreenBase64(q int) string {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    panic(err)
  }

  var buf bytes.Buffer
  file := bufio.NewWriter(&buf)
  jpeg.Encode(file, img, &jpeg.Options{ Quality: q })

  return base64.StdEncoding.EncodeToString(buf.Bytes())
}

var JSKeysMap = map[string]int {
  "a": keybd_event.VK_A,
  "d": keybd_event.VK_D,
  "w": keybd_event.VK_W,
  "s": keybd_event.VK_S,
}

func ExecKeyEvent(key string) error {
  kb, err := keybd_event.NewKeyBonding()
  if err != nil {
    return err
  }

  // For linux, it is very important to wait 2 seconds
  if runtime.GOOS == "linux" {
    time.Sleep(2 * time.Second)
  }

  if JSKeysMap[key] == 0 {
    return errors.New("key not found!")
  }

  // Select keys to be pressed
  kb.SetKeys(JSKeysMap[key]) 

	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()

  return nil
}

