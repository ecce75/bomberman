
import {createGameBoard, createScoreboard} from "../gameLogic/gameSetup.js"
import {createStructure, addEvent} from "../framework.js";
import {players} from "../gameLogic/gameSetup.js";
import { createChatbox } from "../gameLogic/gameChat.js";
import {setupEventListeners} from "./utils.js";

export function gameView(game, ws) {
    const root = document.getElementById('root'); // Ensure you have a div with id="app-root" in your HTML

    const gameView = createStructure({
        tag: 'div',
        attr: ['id', 'gameView', 'class', 'bg-default'],
        children: [
        ]
    })


    // Create the game board and chatbox
    const gameBoard = createGameBoard(game);
    const chatbox = createChatbox();
    const score = createScoreboard();
    //const player = createPlayersAndPlace();

    setupEventListeners(ws);



    // Append game board and chatbox to the root
    root.innerHTML = '';
    root.appendChild(gameView);
    gameView.appendChild(chatbox);
    gameView.appendChild(gameBoard);
    gameView.appendChild(score);
    //gameView.appendChild(player);

}

