import { Pos } from "./schema/pos";

const url = "ws://localhost:8080/set_input_pos";
let socket: WebSocket | null = null;
function connect() {
    if (socket && socket.readyState === WebSocket.OPEN) {
        return;
    }
    socket = new WebSocket(url);
    socket.addEventListener("close", function () {
        socket = null;
        console.warn("Socket closed", url);
    });
    socket.addEventListener("open", function () {
        console.debug("Socket opened", url);
    });
}

connect();

export function sendInputPos(x: number, y: number) {
    if (!socket) {
        connect();
        return;
    }
    if (socket.readyState !== WebSocket.OPEN) {
        return;
    }
    const pos = Pos.create({ x, y });
    const bytes = Pos.encode(pos).finish();
    socket.send(bytes);
}
