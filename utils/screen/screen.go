package screen

import (
	"bufio"
	"bytes"
	"image/jpeg"

  "github.com/kbinani/screenshot"
)

type Frame []byte

func Capture(quality uint8) Frame {
  img, err := screenshot.CaptureDisplay(0)
  if err != nil {
    panic(err)
  }

  var buf bytes.Buffer
  file := bufio.NewWriter(&buf)
  jpeg.Encode(file, img, &jpeg.Options{ Quality: int(quality) })

  return buf.Bytes()
}

func CaptureSeq(num uint8, quality uint8) []Frame {
  var frames []Frame
  numInt := int(num)
  for i := 0; i < numInt; i++ {
    frames = append(frames, Capture(quality))
  }
  return frames
}
