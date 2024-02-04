const classOk = [
    "textArea__74543",
    "textAreaSlate_e0e383",
    "slateContainer_b692b3",
    "editor__66464",
];

document.addEventListener(
    "click",
    (e) => {
        const el = e.target as HTMLElement;

        for (const c of classOk) {
            if (el.classList.contains(c)) {
                return;
            }
        }

        if (el.getAttribute("data-slate-node") === "element") {
            return;
        }

        console.log("blocking click", el);

        e.preventDefault();
        e.stopPropagation();
        return false;
    },
    {
        capture: true,
        passive: false,
    },
);
