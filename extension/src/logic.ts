import { getElementScreenPos } from "./dom";

const promptElSelector =
    "#app-mount > div.appAsidePanelWrapper__714a6 > div.notAppAsidePanel__9d124 > div.app_b1f720 > div > div.layers__1c917.layers_a23c37 > div > div > div > div > div.chat__52833 > div.content__1a4fe > div > div.chatContainer__23434 > main > form > div > div > div > div.textArea__74543.textAreaSlate_e0e383.slateContainer_b692b3 > div > div.markup_a7e664.editor__66464.slateTextArea__0661c.fontSize16Padding__48818 > div";
let promptEl: HTMLElement | null = null;

export function getPromptEl(): HTMLElement {
    if (!promptEl) {
        promptEl = document.querySelector(promptElSelector);
    }
    if (!promptEl) {
        const errMsg = "Failed to find prompt element";
        console.warn(errMsg);
        throw new Error(errMsg);
    }
    return promptEl;
}

export function getPromptPos(): { x: number; y: number } | false {
    try {
        return getElementScreenPos(getPromptEl());
    } catch (e) {
        return { x: -1, y: -1 };
    }
}
