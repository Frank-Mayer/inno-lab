import { Pos } from "./schema/pos";

const socket = new WebSocket("ws://localhost:8080/set_input_pos");
socket.addEventListener("close", function () {
    alert("Connection to server lost. Please refresh the page.");
});

export function sendInputPos(x: number, y: number) {
    if (socket.readyState !== WebSocket.OPEN) {
        return;
    }
    const pos = Pos.create({ x, y });
    const bytes = Pos.encode(pos).finish();
    socket.send(bytes);
}
