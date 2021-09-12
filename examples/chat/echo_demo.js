const form = document.querySelector("form");
const input = document.querySelector("input");
const span = document.querySelector("span");

const socket = new WebSocket("ws://localhost:7000/chat/echo");

socket.addEventListener("open", () => {
  span.innerHTML =
    span.innerHTML + `<br/>Websocket connection initialized with ${socket.url}`;
});

socket.addEventListener("close", () => {
  span.innerHTML =
    span.innerHTML + `<br/>Websocket connection terminated from ${socket.url}`;
});

socket.addEventListener("message", (e) => {
  span.innerHTML = span.innerHTML + `<br/>${e.data}`;
});

form.addEventListener("submit", (e) => {
  e.preventDefault();
  sendMessage(input.value);
  input.value = "";
});

function sendMessage(message) {
  span.innerHTML =
    span.innerHTML + `<br/>Sending message "${message}" to ${socket.url}`;
  socket.send(message);
}
