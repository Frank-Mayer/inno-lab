import { onWsClose, onWsMessage, onWsOpen } from "./handler";

const socket = new WebSocket("ws://localhost:8080/ws");
socket.addEventListener("open", function () {
    onWsOpen(socket);
});

socket.addEventListener("close", function () {
    onWsClose(socket);
});

socket.addEventListener("message", function (event) {
    console.log("Message received", event);
    onWsMessage(socket, event.data);
});
