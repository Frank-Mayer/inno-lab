import { Res } from "./schema/res";

const socket = new WebSocket("ws://localhost:8080/send_image");
socket.addEventListener("close", function () {
    alert("Connection to server lost. Please refresh the page.");
});

export function sendImage(prompt: string, src: string, date: Date) {
    if (socket.readyState !== WebSocket.OPEN) {
        return;
    }
    const pos = Res.create({ prompt, src, time: date.getTime() });
    const bytes = Res.encode(pos).finish();
    socket.send(bytes);
}
