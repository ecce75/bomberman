import { speed } from "../gameLogic/gameInfo.js";

export function setupEventListeners(ws) {
    let gameStarted = false;

    const throttle = (func, getSpeed) => {
        let lastFunc;
        let lastRan;
        return function() {
            const context = this;
            const args = arguments;
            if (!gameStarted) return;  // Stop processing if game hasn't started
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
        if (!gameStarted) return;  // Stop key processing if game hasn't started
        let move = "";
        switch (event.key) {
            case 'ArrowUp': move = 'up'; break;
            case 'ArrowDown': move = 'down'; break;
            case 'ArrowLeft': move = 'left'; break;
            case 'ArrowRight': move = 'right'; break;
            case ' ':
                console.log("Bomb action taken");
                ws.send(JSON.stringify({ type: 'bomb', payload: "bomb" }));
                return;
            default: return;
        }
        ws.send(JSON.stringify({ type: 'move', payload: move }));
    };

    const throttledHandleKeyDown = throttle(handleKeyDown, getSpeed);

    document.addEventListener('keydown', throttledHandleKeyDown);

    return {
        startGame: () => gameStarted = true,
        stopGame: () => gameStarted = false
    };
}

function getSpeed() {
    return speed;  // Assume speed is a constant or dynamic value fetched from somewhere
}
