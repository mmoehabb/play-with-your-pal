package screen

import (
	"bufio"
	"bytes"
	"image/jpeg"

	"github.com/kbinani/screenshot"
)

type Frame []byte

func Capture(quality uint8) (Frame, error) {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    return nil, err
  }

  var buf bytes.Buffer
  file := bufio.NewWriter(&buf)
  jpeg.Encode(file, img, &jpeg.Options{ Quality: int(quality) })

  return buf.Bytes(), nil
}

func CaptureSeq(num uint8, quality uint8) ([]Frame, error) {
  var frames []Frame
  numInt := int(num)
  for i := 0; i < numInt; i++ {
    newframe, err := Capture(quality)
    if err != nil {
      return nil, err
    }
    frames = append(frames, newframe)
  }
  return frames, nil
}

