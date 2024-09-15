// cat *.png | ffmpeg -framerate 1 -i pipe: ./stream.mp4
package encoder

import (
	"os/exec"
	"strconv"

	"goweb/utils/screen"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func Encode(frames []screen.Frame) []byte {
  cmd := exec.Command("ffmpeg",
    "-framerate", strconv.Itoa(len(frames)),
    "-y", "-i", "pipe:",
    "./stream.mp4",
  )
  
  inpipe, err := cmd.StdinPipe()
  checkErr(err)
  go func() {
    defer inpipe.Close()
    for _, f := range frames {
      inpipe.Write(f)
    }
  }()

  out, err := cmd.CombinedOutput()
  checkErr(err)
  return out
}

