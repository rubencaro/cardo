// SSE example

let source = new EventSource('/events')

source.onmessage = (e) => {
  let txt = document.createTextNode(e.data)
  let div = document.createElement("div")
  div.appendChild(txt)

  let first = document.getElementById("messages").firstChild
  document.getElementById("messages").insertBefore(div, first)
}

// handle reconnection on our own
let lastState = EventSource.CONNECTING
window.setInterval(() => {
  let state = source.readyState
  if (state == EventSource.CLOSED && lastState == EventSource.CLOSED) {
    console.log("Resetting EventSource")
    source.close()
    source = new EventSource('/events')
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