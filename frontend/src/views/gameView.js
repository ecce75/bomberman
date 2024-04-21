import { createGameBoard, createScoreboard } from "../gameLogic/gameSetup.js"
import { createChatbox } from "../gameLogic/chat.js";
import { createStructure } from "../framework.js";

export function gameView(map) {
    const root = document.getElementById('root'); // Ensure you have a div with id="app-root" in your HTML

    const gameView = createStructure({
        tag: 'div',
        attr: ['id', 'gameView', 'class', 'bg-default'],
        children: [
        ]
    })


    // Create the game board and chatbox
    const gameBoard = createGameBoard(map);
    const chatbox = createChatbox();
    const score = createScoreboard();

    // Append game board and chatbox to the root
    root.innerHTML = '';
    root.appendChild(gameView);
    gameView.appendChild(chatbox);
    gameView.appendChild(gameBoard);
    gameView.appendChild(score);

}

