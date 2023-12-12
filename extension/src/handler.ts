import { Command, Message } from "./schema/message";
import { getPromptPos, newMessage } from "./logic";
import { getElementScreenPos } from "./dom";
import * as log from "./log";

let active = false;

export function onWsOpen(ws: WebSocket) {
    log.info("Client connected");
    active = true;
}

export function onWsClose(ws: WebSocket) {
    log.info("Client disconnected");
    active = false;
}

export async function onWsMessage(ws: WebSocket, message: unknown) {
    log.info("Message received", "message", message);
    if (!active) {
        return;
    }

    try {
        const msg = Message.decode(await transformEventData(message));
        log.info("Decoded message", "message", msg);
        switch (msg.command) {
            case Command.GET_INPUT_POSITION:
                const answer = newMessage(
                    { command: Command.RETURN_INPUT_POSITION },
                    getPromptPos(),
                );
                send(ws, answer);
                break;
            case Command.REGISTER_PROMPT:
                log.warn("TODO: implement register prompt");
                break;
            default:
                log.error("Unrecognized command", "command", msg.command);
                break;
        }
    } catch (e) {
        log.error("Failed to decode message", "error", e);
    }
}

function send(ws: WebSocket, message: Message) {
    if (!active) {
        return;
    }

    const bytes = Message.encode(message).finish();

    try {
        ws.send(bytes);
    } catch (e) {
        console.error("Failed to send message", e);
        active = false;
    }
}

async function transformEventData(data: unknown): Promise<Uint8Array> {
    if (typeof data === "string") {
        console.log("transforming string data", data);
        return Promise.resolve(Uint8Array.from(data, (c) => c.charCodeAt(0)));
    }

    if (data instanceof ArrayBuffer) {
        console.log("transforming array buffer data", data);
        return Promise.resolve(new Uint8Array(data));
    }

    if (data instanceof Uint8Array) {
        console.log("transforming uint8array data", data);
        return Promise.resolve(data);
    }

    if (data instanceof Blob) {
        console.log("transforming blob data", data);
        const buffer = await readBlobAsArrayBuffer(data);
        return new Uint8Array(buffer);
    }

    throw new Error(
        `Unsupported data type: ${typeof data} ${(data as object).constructor
            ?.name}`,
    );
}

function readBlobAsArrayBuffer(blob: Blob): Promise<ArrayBuffer> {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => {
            if (reader.result instanceof ArrayBuffer) {
                resolve(reader.result);
            } else {
                reject(new Error("Failed to read Blob as ArrayBuffer"));
            }
        };
        reader.onerror = () => {
            reject(new Error("Error reading Blob as ArrayBuffer"));
        };
        reader.readAsArrayBuffer(blob);
    });
}
