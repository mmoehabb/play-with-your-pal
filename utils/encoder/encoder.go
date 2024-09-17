// cat *.png | ffmpeg -framerate 1 -i pipe: ./stream.mp4
package encoder

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"goweb/utils/screen"
)

var STREAM_FILE_PATH = "./tmp/pwyp-stream.mp4"

func checkErr(err error) {
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
}

func Init() {
  os.Mkdir("./tmp", os.ModePerm)
}

func Encode(frames []screen.Frame) []byte {
  cmd := exec.Command("ffmpeg",
    "-framerate", strconv.Itoa(len(frames)),
    "-y", "-i", "pipe:",
    "-preset", "veryfast",
    "-s", "1280x720",
    STREAM_FILE_PATH,
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

  buf, err := os.ReadFile(STREAM_FILE_PATH)
  checkErr(err)
  return buf
}

