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
          img_stream.src = "data:image/jpeg;base64," + evt.data
      };
  } else {
      console.log("Your browser does not support WebSockets.")
  }
}

window.onkeydown = (e) => {
  if (!conn) return;
  conn.send(e.key);
}

window.addEventListener("load", (event) => {
  document.querySelector("input").onchange = (e) => password = e.target.value
});
