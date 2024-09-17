function switchMode() {
  const html = document.getElementsByTagName("html")[0]
  const theme = html.getAttribute("theme")
  html.setAttribute("theme", theme !== "dark" ? "dark" : "light")
}

let conn = null
let password = ""

// preserve URL.createObjectURL result to be revoked and swaped on each ws message.
let burl = null

let video_stream = null
let canvas_stream = null
let ctx = null

function drawVideo() {
  ctx.drawImage(video_stream, 0, 0, canvas_stream.width, canvas_stream.height)
  requestAnimationFrame(drawVideo)
}

function join() {
  video_stream = document.getElementById("video_stream")
  canvas_stream = document.getElementById("canvas_stream")
  canvas_stream.width = window.innerWidth - 100
  canvas_stream.height = canvas_stream.width * 720 / 1280
  ctx = canvas_stream.getContext("2d")
  requestAnimationFrame(drawVideo)
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.host + "/ws/join/" + password);
      conn.onclose = function (evt) {
          console.log("session closed.")
      };
      conn.onmessage = function (evt) {
        const b = new Blob([evt.data], { type: "video/mp4" })
        URL.revokeObjectURL(burl)
        burl = URL.createObjectURL(b)
        video_stream.src = burl
      };
  } else {
      console.log("Your browser does not support WebSockets.")
  }
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
  canvas_stream.width = window.innerWidth

  document.querySelector("main").append(closeBtn)
}

function normalscreen() {
  document.querySelector("#closebtn").remove()
  document.querySelector("header").style.display = "flex"
  document.querySelector("footer").style.display = "flex"
  canvas_stream.style.position = ""
  canvas_stream.width = window.innerWidth - 100
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
