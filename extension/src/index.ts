console.log("content.js loaded");

let interval = 0;

let i = 0;

const verticalCorrection = 85;

const socket = new WebSocket("ws://localhost:8080/ws");
socket.addEventListener("open", function (event) {
  console.log("socket opened", event);
  interval = window.setInterval(() => {
    ++i;
    if (i > 3) {
      window.clearInterval(interval);
      return;
    }
    main()
      .then(() => {
        console.log("main done");
      })
      .catch((err) => {
        console.error(err);
      });
  }, 2000);
});

socket.addEventListener("close", function (event) {
  if (interval) {
    interval = 0;
    window.clearInterval(interval);
    console.log("socket closed", event);
  }
});

socket.addEventListener("message", function (event) {
  alert("Message from server " + event.data);
});

async function main() {
  console.log("main");
  clickOn(
    "#app-mount > div.appAsidePanelWrapper__714a6 > div.notAppAsidePanel__9d124 > div.app_b1f720 > div > div.layers__1c917.layers_a23c37 > div > div > div > div > div.chat__52833 > div.content__1a4fe > div > div.chatContainer__23434 > main > form > div > div > div > div.textArea__74543.textAreaSlate_e0e383.slateContainer_b692b3 > div > div.markup_a7e664.editor__66464.slateTextArea__0661c.fontSize16Padding__48818 > div",
  );
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * @param {HTMLElement} element
 * @returns {Object} {x: number, y: number}
 */
function getElementScreenPos(element) {
  const rect = element.getBoundingClientRect();
  const scrollLeft = window.pageXOffset || document.documentElement.scrollLeft;
  const scrollTop = window.pageYOffset || document.documentElement.scrollTop;

  return {
    x: rect.left + scrollLeft + window.screenX,
    y: rect.top + scrollTop + window.screenY + verticalCorrection,
  };
}

function clickOn(selector) {
  const element = document.querySelector(selector);
  if (!element) {
    console.error("no element found for selector", selector);
    return;
  }

  console.log("clicking", element);

  const { x, y } = getElementScreenPos(element);

  socket.send(JSON.stringify({ command: "click", x, y }));

  window.setTimeout(() => {
    socket.send(JSON.stringify({ command: "prompt" }));
  }, 1000);
}
