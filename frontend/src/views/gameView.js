import {createGameBoard} from "../gameLogic/gameSetup.js"

export default function gameView() {
    const root = document.getElementById('root'); // Ensure you have a div with id="app-root" in your HTML
    const content = createGameBoard(11, 11)
    root.appendChild(content);
}