import {createGameBoard, createChatbox, createScoreboard} from "../gameLogic/gameSetup.js"

export default function gameView() {
    const root = document.getElementById('root'); // Ensure you have a div with id="app-root" in your HTML

    // Create the game board and chatbox
    const gameBoard = createGameBoard(11, 11);
    const chatbox = createChatbox();
    const score = createScoreboard();

    // Append game board and chatbox to the root
    root.appendChild(chatbox);
    root.appendChild(gameBoard);
    root.appendChild(score);
}

