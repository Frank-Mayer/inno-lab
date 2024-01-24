import { Res } from "./schema/res";

const url = "ws://localhost:8080/send_image";
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

export function sendImage(prompt: string, src: string, date: Date) {
    if (!socket) {
        connect();
        setTimeout(() => {
            sendImage(prompt, src, date);
        }, 1000);
        console.warn("Socket not connected", url, "retry in one second");
        return;
    }

    if (socket.readyState !== WebSocket.OPEN) {
        console.warn("Socket not open", socket.readyState);
        return;
    }
    const pos = Res.create({ prompt, src, time: date.getTime() });
    const bytes = Res.encode(pos).finish();
    socket.send(bytes);
}
