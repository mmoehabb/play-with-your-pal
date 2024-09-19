/*
decode function takes the dynamic image code string and returns an array of chunks
whereas a chunk is defined as follows:
{
 data: ImageData, // duplicated pixels info stored in ImageData obj
 start: number, // the start of the chunk as represented in the code
 end: number, // the end of the chunck as represented in the code
}
code example: 0,10#FF0000|11,15#00FF00
the above example represents two chunks: one starts from position 0 and ends at 10 displaying eleven blue pixels, whereas the other starts at 11 and ends at 15 displaying five green pixels.

Chunk code template: "START,END#BBGGRR|"
*/
function decode(code) {
  const chunks = []
  const ccs = code.split("|");
  for (let code_chunk of ccs) {
    if (code_chunk == "") continue;
    const ch = {}
    const ch_parts = code_chunk.split("#")
    // get the start and end numbers
    const header = ch_parts[0].split(",")
    ch.start = parseInt("0x" + header[0])
    ch.end = parseInt("0x" + header[1])
    // get the pixel array [opacity, blue, green, red]
    const bgr_str = ch_parts[1]
    const abgr = []
    if (bgr_str) {
      abgr[0] = parseInt("0x" + bgr_str.slice(0, 2))
      abgr[1] = parseInt("0x" + bgr_str.slice(2, 4))
      abgr[2] = parseInt("0x" + bgr_str.slice(4, 6))
      abgr[3] = 255
    }
    // convert the array to ImageData and store it in ch
    ch.data = abgr
    // push the created chunk into chunks array
    chunks.push(ch)
  }
  return chunks
}
