const verticalCorrection = 85;

export function getElementScreenPos(element: HTMLElement) {
    const rect = element.getBoundingClientRect();
    const scrollLeft =
        window.pageXOffset || document.documentElement.scrollLeft;
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;

    return {
        x: rect.left + scrollLeft + window.screenX,
        y: rect.top + scrollTop + window.screenY + verticalCorrection,
    };
}
