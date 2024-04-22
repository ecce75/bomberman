import {createGameBoard, createScoreboard, players} from "../gameLogic/gameSetup.js"
import {createStructure} from "../framework.js";
import {createChatbox} from "../gameLogic/gameChat.js";
import {setupEventListeners} from "./utils.js";

export function gameView(game, ws) {
    const root = document.getElementById('root');

    const gameViewStructure = createStructure({
        tag: 'div',
        attr: ['id', 'gameView', 'class', 'bg-default'],
        children: [{
            tag: 'div',
            attr: ['id', 'countdownDisplay', 'class', 'countdown-style'],
            children: ['Game a go start in 10 sekond. WALK GOOD']
        }]
    });

    const gameBoard = createGameBoard(game);
    const chatbox = createChatbox();
    const scoreboard = createScoreboard();

    const { startGame, stopGame } = setupEventListeners(ws);

    root.innerHTML = '';
    root.appendChild(gameViewStructure);
    gameViewStructure.appendChild(chatbox);
    gameViewStructure.appendChild(gameBoard);
    gameViewStructure.appendChild(scoreboard);

    startCountdown(startGame);
}

function startCountdown(startGame) {
    let countdown = 10;
    const countdownDisplay = document.getElementById('countdownDisplay');

    const intervalId = setInterval(() => {
        countdown -= 1;
        countdownDisplay.textContent = `Game a go start in ${countdown} sekond. WALK GOOD`;

        if (countdown <= 0) {
            clearInterval(intervalId);
            countdownDisplay.textContent = 'Sho\' dem, mon!';
            startGame();  // Start the game interactions
        }
    }, 1000);
}
