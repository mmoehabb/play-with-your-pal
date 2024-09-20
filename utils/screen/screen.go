package screen

import (
	"image"

	"github.com/kbinani/screenshot"
	"github.com/nfnt/resize"
)

var Bounds = screenshot.GetDisplayBounds(0)

func Capture(quality int) (*image.RGBA, error) {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    return nil, err
  }
  width := uint(Bounds.Max.X * quality / 100)
  height := uint(Bounds.Max.Y * quality / 100)
  newimg := resize.Resize(width, height, img, resize.Lanczos2)
  return newimg.(*image.RGBA), nil
}

