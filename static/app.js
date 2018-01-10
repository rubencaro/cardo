// SSE example

let socket = new WebSocket("ws://" + document.location.host + "/ws")

socket.onmessage = (e) => {
  let txt = document.createTextNode(e.data)
  let div = document.createElement("div")
  div.appendChild(txt)

  let first = document.getElementById("messages").firstChild
  document.getElementById("messages").insertBefore(div, first)
}

// handle reconnection on our own
let lastState = WebSocket.CONNECTING
window.setInterval(() => {
  let state = socket.readyState
  if (state == WebSocket.CLOSED && lastState == WebSocket.CLOSED) {
    console.log("Resetting WebSocket")
    socket.close()
    socket = new WebSocket("ws://" + document.location.host + "/ws")
  }
  else {
    lastState = state
  }
}, 5000)

// let add = function () {
//   let body = { "msg": document.getElementById("text").value }
//   fetch("/add", {
//     "method": "POST",
//     "headers": { "content-type": "application/json" },
//     "body": JSON.stringify(body)
//   })
// }