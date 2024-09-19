package screen

import (
	"image"

	"github.com/kbinani/screenshot"
	"github.com/nfnt/resize"
)

func Capture(quality int) (*image.RGBA, error) {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    return nil, err
  }
  width, height := uint(1280 * quality / 100), uint(720 * quality / 100)
  newimg := resize.Resize(width, height, img, resize.Lanczos2)
  return newimg.(*image.RGBA), nil
}

