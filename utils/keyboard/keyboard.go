package keyboard

import (
	"bufio"
	"bytes"
	"errors"
	"image/jpeg"
	"log"

	"github.com/kbinani/screenshot"
	"github.com/micmonay/keybd_event"
)

func CapScreenBase64(q int) []byte {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    panic(err)
  }

  var buf bytes.Buffer
  file := bufio.NewWriter(&buf)
  jpeg.Encode(file, img, &jpeg.Options{ Quality: q })

  return buf.Bytes()
}

var JSKeysMap = map[string]int {
  "a": keybd_event.VK_A,
  "b": keybd_event.VK_B,
  "c": keybd_event.VK_C,
  "d": keybd_event.VK_D,
  "e": keybd_event.VK_E,
  "f": keybd_event.VK_F,
  "g": keybd_event.VK_G,
  "h": keybd_event.VK_H,
  "i": keybd_event.VK_I,
  "j": keybd_event.VK_J,
  "k": keybd_event.VK_K,
  "l": keybd_event.VK_L,
  "m": keybd_event.VK_M,
  "n": keybd_event.VK_N,
  "o": keybd_event.VK_O,
  "p": keybd_event.VK_P,
  "q": keybd_event.VK_Q,
  "r": keybd_event.VK_R,
  "s": keybd_event.VK_S,
  "t": keybd_event.VK_T,
  "u": keybd_event.VK_U,
  "v": keybd_event.VK_V,
  "w": keybd_event.VK_W,
  "x": keybd_event.VK_X,
  "y": keybd_event.VK_Y,
  "z": keybd_event.VK_Z,
  "ArrowUp": keybd_event.VK_UP,
  "ArrowDown": keybd_event.VK_DOWN,
  "ArrowRight": keybd_event.VK_RIGHT,
  "ArrowLeft": keybd_event.VK_LEFT,
}

var kb keybd_event.KeyBonding

func InitKB() {
  var err error
  kb, err = keybd_event.NewKeyBonding()
  if err != nil {
    panic(err)
  }
}

func ExecKey(method string, key string) error {
  var err error
  if method == "press" {
    err = PressKey(key)
  } else if method == "release" {
    err = ReleaseKey(key)
  } else {
    err = errors.New("Invalid method!")
  }
  return err 
}

func PressKey(key string) error {
  if JSKeysMap[key] == 0 {
    return errors.New(key + " key not found!")
  }
  // Select keys to be pressed
  kb.SetKeys(JSKeysMap[key]) 
  err := kb.Press()
  if err != nil {
    return err
  }
  log.Println(key, " pressed.")
  return nil
}

func ReleaseKey(key string) error {
  if JSKeysMap[key] == 0 {
    return errors.New(key + " key not found!")
  }
  // Select keys to be released
  kb.SetKeys(JSKeysMap[key]) 
  err := kb.Release()
  if err != nil {
    return err
  }
  log.Println(key, " released.")
  return nil
}

