function switchMode() {
  const html = document.getElementsByTagName("html")[0]
  const theme = html.getAttribute("theme")
  html.setAttribute("theme", theme !== "dark" ? "dark" : "light")
}

let img_stream = null
let conn = null
let password = ""

function join() {
  img_stream = document.querySelector("#stream")
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.host + "/ws/join/" + password);
      conn.onclose = function (evt) {
          console.log("session closed.")
      };
      conn.onmessage = function (evt) {
          //img_stream.src = "data:image/jpeg;base64," + evt.data
          console.log(evt.data)
          conn.close()
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
  img_stream.style.position = "absolute"
  img_stream.style.width = "100vw"

  document.querySelector("main").append(closeBtn)
}

function normalscreen() {
  document.querySelector("#closebtn").style.display = "none"
  document.querySelector("header").style.display = "flex"
  document.querySelector("footer").style.display = "flex"
  img_stream.style.position = ""
  img_stream.style.width = ""
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
