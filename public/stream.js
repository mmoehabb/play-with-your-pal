/*
let canvas;
let ctx;
let init;
let config;
let decoder;

function startStream() {
  canvas = document.getElementById("stream");
  ctx = canvas.getContext("2d");

  function handleFrame(frame) {
    ctx.drawImage(frame, 0, 0)
  }

  init = {
    output: handleFrame,
    error: (e) => {
      console.log(e.message);
    },
  };

  config = {
    codec: "avc1.4d002a",
  };

  decoder = null

  VideoDecoder.isConfigSupported(config).then(({supported}) => {
    if (supported) {
      decoder = new VideoDecoder(init);
      decoder.configure(config);
    } else {
      console.error("config error!");
    }
  })
}

function drawChunks(chunks) {
  console.log(chunks)
  for (let i = 0; i < chunks.length; i++) {
    const chunk = new EncodedVideoChunk({
      timestamp: 0,
      type: "key",
      data: new Uint8Array(chunks[i]),
    });
    console.log(chunk)
    decoder.decode(chunk);
  }
  decoder.flush();
}
*/
