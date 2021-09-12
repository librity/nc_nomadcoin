const span = document.querySelector("span");

const socket = new WebSocket("ws://localhost:7000/chat/ping");
var isConnected = true;

socket.addEventListener("open", () => {
  span.innerHTML =
    span.innerHTML + `<br/>Websocket connection initialized with ${socket.url}`;
});

socket.addEventListener("close", () => {
  span.innerHTML =
    span.innerHTML + `<br/>Websocket connection terminated from ${socket.url}`;
  isConnected = false;
});

setTimeout(() => {
  sendMessage("Message from the browser through web sockets!");
}, 3000);

setTimeout(() => {
  sendMessage("Another message from the browser through web sockets!");
}, 6000);

setTimeout(() => {
  sendMessage("Yet another message from the browser through web sockets!");
}, 9000);

ping();

function ping() {
  setTimeout(() => {
    if (!isConnected) {
      return;
    }

    sendMessage("Ping");

    ping();
  }, 1000);
}

function sendMessage(message) {
  span.innerHTML =
    span.innerHTML + `<br/>Sending message "${message}" to ${socket.url}`;
  socket.send(message);
}
