function switchMode() {
  const html = document.getElementsByTagName("html")[0]
  const theme = html.getAttribute("theme")
  html.setAttribute("theme", theme !== "dark" ? "dark" : "light")
}

let session_id = null
let conn = null

function host() {
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.host + "/ws/host");
      conn.onclose = function (evt) {
          console.log("session closed.")
      };
      conn.onmessage = function (evt) {
          console.log("ws message: ", evt.data)
          session_id = JSON.parse(evt.data)['session_id']
      };
  } else {
      console.log("Your browser does not support WebSockets.")
  }
}

function join(id) {
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.host + "/ws/join/" + id);
      conn.onclose = function (evt) {
          console.log("join session closed.")
      };
      conn.onmessage = function (evt) {
          console.log("join ws message: ", evt.data)
      };
  } else {
      console.log("Your browser does not support WebSockets.")
  }
}

window.onkeydown = (e) => {
  if (!conn) return;
  conn.send(e.key);
}
