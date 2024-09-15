package keyboard

import (
	"errors"
	"log"

	"github.com/micmonay/keybd_event"
)

var JSKeysMap = map[string]int {
  "a": keybd_event.VK_A,
  "d": keybd_event.VK_D,
  "w": keybd_event.VK_W,
  "s": keybd_event.VK_S,
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

