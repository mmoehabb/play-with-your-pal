package encoder

import (
	"fmt"
	"image"
	"slices"
	"strings"
	"sync"
)

type EncodedChunk struct {
  order uint8
  code string
}

const (
  threshold = 500000 // threshold mod 4 should equal 0
)

var lastcode string
var lastMap = make(map[int]string)
var curMap = make(map[int]string)

var chunksChan = make(chan EncodedChunk)

var mu sync.Mutex
var mapMu sync.Mutex

func Encode(img *image.RGBA) string {
  // Fill chuncks channel with EncodedChunks
  var pixels = img.Pix
  var pln = len(pixels)
  var nu_chunks = pln / threshold
  for cn := 0; cn < int(nu_chunks); cn++ {
    go encodeChunk(cn, pixels[cn*threshold:(cn+1)*threshold])
  }
  // Collect chunks from the channel and join into one big chunk 
  var chunks = make([]string, nu_chunks)
  var l = 0
  for l < nu_chunks {
    echunk := <-chunksChan
    chunks[echunk.order] = echunk.code
    l += 1
  }
  var newcode = strings.Join(chunks, "")
  // Detect the difference from the lastcode var, and return it only
  if lastcode == "" {
    lastcode = newcode
    return newcode
  }
  var diff strings.Builder
  sp := strings.Split(lastcode, "|")
  sp = sp[:len(sp)-1]
  for i, elm := range curMap {
    if lastMap[i] != elm {
      diff.WriteString(elm)
      diff.WriteString("|")
    }
  }
  lastcode = newcode
  lastMap = curMap
  curMap = make(map[int]string)
  return diff.String()
}

func encodeChunk(order int, chunck []uint8) {
  var res EncodedChunk
  var sb strings.Builder
  res.order = uint8(order)

  var lasthex = chunck[0:4]
  start := (order * threshold)/4
  end := start
  for i := 0; i < len(chunck); i+=4 {
    if slices.Compare(lasthex, chunck[i:i+4]) == 0 {
      end += 1
      continue
    }

    hex := toHex(lasthex, start, end-1)
    sb.Write([]byte(hex))

    mapMu.Lock()
    curMap[start] = hex
    curMap[end] = hex
    mapMu.Unlock()

    lasthex = chunck[i:i+4]
    start = (order * threshold)/4 + i/4
    end = start + 1
  }
  res.code = sb.String()
  mu.Lock()
  chunksChan <- res
  mu.Unlock()
}

func toHex(pix []uint8, start int, end int) string {
  return fmt.Sprintf("%X,%X#%02X%02X%02X|", start, end, pix[0], pix[1], pix[2])
}
