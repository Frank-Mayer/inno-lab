import { getElementScreenPos } from "./dom";
import { Command, Message } from "./schema/message";

const promptElSelector =
    "#app-mount > div.appAsidePanelWrapper__714a6 > div.notAppAsidePanel__9d124 > div.app_b1f720 > div > div.layers__1c917.layers_a23c37 > div > div > div > div > div.chat__52833 > div.content__1a4fe > div > div.chatContainer__23434 > main > form > div > div > div > div.textArea__74543.textAreaSlate_e0e383.slateContainer_b692b3 > div > div.markup_a7e664.editor__66464.slateTextArea__0661c.fontSize16Padding__48818 > div";
let promptEl: HTMLElement | null = null;

export function getPromptEl(): HTMLElement {
    if (!promptEl) {
        promptEl = document.querySelector(promptElSelector);
    }
    if (!promptEl) {
        const errMsg = "Failed to find prompt element";
        alert(errMsg);
        throw new Error(errMsg);
    }
    return promptEl;
}

export function newMessage(...combine: Partial<Message>[]): Message {
    const msg: Message = {
        command: Command.UNRECOGNIZED,
        x: 0,
        y: 0,
        prompt: "",
        url: "",
    };

    for (const c of combine) {
        Object.assign(msg, c);
    }
    return msg;
}

export function getPromptPos(): { x: number; y: number } {
    return getElementScreenPos(getPromptEl());
}
