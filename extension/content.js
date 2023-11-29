console.log("content.js loaded");

const socket = new WebSocket("ws://localhost:8080/ws");
socket.addEventListener("open", function (event) {
  main();
});

socket.addEventListener("message", function (event) {
  alert("Message from server " + event.data);
});

async function main() {
  sleep(10000);
  selectChannelIfNotSelected("Midjourney");
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

  const x = rect.x + window.screenX;
  const y = rect.y + window.screenY;

  return { x, y };
}

/*
selected:
interactive__776ee interactive_a868bc interactiveSelected_ec846b selected_d94cf9

not selected:
interactive__776ee interactive_a868bc
*/

function selectChannelIfNotSelected(name) {
  const els = Array.of(
    document.querySelectorAll(
      "interactive__776ee.interactive_a868bc:not(.ineractiveSelected_ec846b):not(.selected_d94cf9)",
    ),
  );

  for (const el of els) {
    const channelName = el.textContent;
    if (!channelName.includes(name)) {
      continue;
    }

    console.log("clicking", el);

    const { x, y } = getElementScreenPos(el);

    socket.send(JSON.stringify({ command: "click", x, y }));

    window.setTimeout(() => {
      socket.send(JSON.stringify({ command: "prompt" }));
    }, 1000);
  }
}
