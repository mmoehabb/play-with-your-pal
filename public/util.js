function switchMode() {
  const html = document.getElementsByTagName("html")[0]
  const theme = html.getAttribute("theme")
  html.setAttribute("theme", theme !== "dark" ? "dark" : "light")
}

let conn = null
let password = ""

let canvas_stream = null
let ctx = null

let firstmsg = true
let w = 0
let h = 0
let ctx_img_arr = []
let img_arr_len = 0

function join() {
  canvas_stream = document.getElementById("canvas_stream")
  ctx = canvas_stream.getContext("2d")
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.host + "/ws/join/" + password);
      conn.onclose = function () {
          console.log("session closed.")
      }
      conn.onmessage = function (evt) {
        if (firstmsg) {
          console.log(evt.data)
          setCanvasDim(evt.data)
          firstmsg = false
          return
        }
        chunks = decode(evt.data)
        for (let ch of chunks) {
          for (let i = ch.start; i <= ch.end; i++) {
            ctx_img_arr[i*4] = ch.data[0] 
            ctx_img_arr[i*4 + 1] = ch.data[1] 
            ctx_img_arr[i*4 + 2] = ch.data[2] 
            ctx_img_arr[i*4 + 3] = ch.data[3] 
          }
        }
        ctx_img_arr.length = img_arr_len // workaround... shall change it later 
        try {
          ctx.putImageData(new ImageData(new Uint8ClampedArray(ctx_img_arr), w), 0, 0)
        }
        catch(err) {
          console.error(err)
          conn.close()
        }
      }
  } else {
      console.log("Your browser does not support WebSockets.")
  }
}

function setCanvasDim(firstMsg) {
  const dim = firstMsg.split(",")
  w = parseInt(dim[0])
  h = parseInt(dim[1])
  canvas_stream.width = w
  canvas_stream.height = h
  img_arr_len = w * h * 4
  ctx_img_arr.length = img_arr_len
  ctx_img_arr.fill(0)
}

function fullscreen() {
  const closeBtn = document.createElement("button")
  closeBtn.id = "closebtn"
  closeBtn.innerText = "ⓧ"
  closeBtn.onclick = () => normalscreen()
  closeBtn.className = "absolute top-4 right-4 color-accent text-2xl"

  document.querySelector("header").style.display = "none"
  document.querySelector("footer").style.display = "none"
  canvas_stream.style.position = "absolute"

  document.querySelector("main").append(closeBtn)
}

function normalscreen() {
  document.querySelector("#closebtn").remove()
  document.querySelector("header").style.display = "flex"
  document.querySelector("footer").style.display = "flex"
  canvas_stream.style.position = ""
}

const fired_keys = {}
window.onkeydown = (e) => {
  if (!conn) return;
  if (!fired_keys[e.key]) {
    conn.send("press_" + e.key);
    fired_keys[e.key] = true
  }
  if (e.key == "Escape") {
    normalscreen()
  }
}
window.onkeyup = (e) => {
  if (!conn) return;
  if (fired_keys[e.key] === true) {
    conn.send("release_" + e.key);
    fired_keys[e.key] = false
  }
}

window.addEventListener("load", (event) => {
  document.querySelector("input").onchange = (e) => password = e.target.value
});
