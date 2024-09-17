// cat *.png | ffmpeg -framerate 1 -i pipe: ./stream.mp4
package encoder

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"goweb/utils/screen"
)

func checkErr(err error) {
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
}

func Encode(frames []screen.Frame) []byte {
  cmd := exec.Command("ffmpeg",
    "-framerate", strconv.Itoa(len(frames)),
    "-y", "-i", "pipe:",
    "-preset", "veryfast",
    "-s", "1280x720",
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

  _, err = cmd.CombinedOutput()
  checkErr(err)

  buf, err := os.ReadFile("./stream.mp4")
  checkErr(err)
  return buf
}

