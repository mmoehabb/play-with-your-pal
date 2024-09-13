function switchMode() {
  const html = document.getElementsByTagName("html")[0]
  const theme = html.getAttribute("theme")
  html.setAttribute("theme", theme !== "dark" ? "dark" : "light")
}

let conn = null
let password = ""

function join() {
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.host + "/ws/join/" + password);
      conn.onclose = function (evt) {
          console.log("session closed.")
      };
      conn.onmessage = function (evt) {
          console.log("ws message: ", evt.data)
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
