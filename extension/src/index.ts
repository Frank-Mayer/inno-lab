import { getPromptPos } from "./logic";
import { sendImage } from "./sendImage";
import { sendInputPos } from "./sendInputPos";
import "./dontClick";

const intervals = new Set<number>();

function cleanup() {
    for (const interval of intervals) {
        window.clearInterval(interval);
    }
    alert("Connection to server lost. Please refresh the page.");
}

window.setTimeout(() => {
    console.debug("Starting prompt position interval");
    intervals.add(
        window.setInterval(() => {
            try {
                const pos = getPromptPos();
                console.debug("Prompt pos", pos);
                if (!pos) {
                    return;
                }

                sendInputPos(pos.x, pos.y);
            } catch (e) {
                console.error(e);
                cleanup();
            }
        }, 5000),
    );
}, 2500);

const reportedMessages = new Set<string>();
console.debug("Starting message interval");
intervals.add(
    window.setInterval(() => {
        try {
            watchForNewMessages();
        } catch (e) {
            console.error(e);
            cleanup();
        }
    }, 1000),
);

const startDateTime = Date.now();
function watchForNewMessages() {
    const messages = Array.from(
        document.querySelectorAll("*[aria-roledescription='Message']"),
    ) as HTMLElement[];
    for (const message of messages) {
        // Ignore messages older than 24 hours
        const timeEl = message.querySelector("time");
        if (!timeEl) {
            continue;
        }
        const time = timeEl.getAttribute("datetime");
        if (!time) {
            continue;
        }
        const date = new Date(time);
        if (date.getTime() < startDateTime) {
            continue;
        }

        // Ignore messages that are not 100% complete (regex in message.innerHTML)
        if (
            message.innerHTML.includes("Waiting to start") ||
            message.innerHTML.includes("<span>%</span>")
        ) {
            continue;
        }

        // Get the image
        const img = message.querySelector("a[data-role='img']");
        if (!img) {
            continue;
        }
        const src = img.getAttribute("href")!;
        if (!src) {
            continue;
        }

        // Get the prompt
        const promptEl = message.querySelector("strong");
        if (!promptEl) {
            continue;
        }
        const prompt = promptEl.textContent;
        if (!prompt) {
            continue;
        }

        if (reportedMessages.has(prompt)) {
            continue;
        }

        // Send the image to the server
        sendImage(prompt, src, date);

        reportedMessages.add(prompt);
    }
}
