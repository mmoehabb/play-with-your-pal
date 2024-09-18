package screen

import (
	"image"
	"github.com/kbinani/screenshot"
)

func Capture(quality int) (*image.RGBA, error) {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    return nil, err
  }
  return img, nil
}

