import {speed} from "../gameLogic/gameInfo.js";

export function setupEventListeners(ws) {
    const throttle = (func, getSpeed) => {
        let lastFunc;
        let lastRan;
        return function() {
            const context = this;
            const args = arguments;
            const limit = getSpeed();
            if (!lastRan) {
                func.apply(context, args);
                lastRan = Date.now();
            } else {
                clearTimeout(lastFunc);
                lastFunc = setTimeout(function() {
                    if ((Date.now() - lastRan) >= limit) {
                        func.apply(context, args);
                        lastRan = Date.now();
                    }
                }, limit - (Date.now() - lastRan));
            }
        }
    };

    const handleKeyDown = (event) => {
        let move = "";
        switch (event.key) {
            case 'ArrowUp': move = 'up'; break;
            case 'ArrowDown': move = 'down'; break;
            case 'ArrowLeft': move = 'left'; break;
            case 'ArrowRight': move = 'right'; break;
            case ' ':
                console.log("bomb");
                ws.send(JSON.stringify({ type: 'bomb', payload: "bomb" }));
                return;
            default: return; // Ignore other keys
        }
        ws.send(JSON.stringify({ type: 'move', payload: move }));
    };

    const throttledHandleKeyDown = throttle(handleKeyDown, getSpeed);

    document.addEventListener('keydown', event => {
        if (event.key === ' ') {
            handleKeyDown(event);
        }
    });

    document.addEventListener('keydown', throttledHandleKeyDown);
}

function getSpeed() {
    console.log(speed)
    return speed;
}