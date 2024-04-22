import { players } from './gameSetup.js';
import { updateLivesDisplay } from './gameInfo.js';
import {updateField} from "./movement.js";

export function activateBomb(pos) {
    const bombElement = document.querySelector(`.cell[style*="grid-area: ${pos.Y + 1} / ${pos.X + 1}"]`);

    if (bombElement) {
        bombElement.classList.add('bomb');
    }
    setTimeout(() => {
        bombElement.classList.remove('bomb');
    }, 3000);
}

export function activateFlames(pos) {
    for (let i = 0; i < pos.length; i++) {
        const flamePos = pos[i].Position;
        const centerFlamePos = pos[0].Position;
        const flameElement = document.querySelector(`.cell[style*="grid-area: ${flamePos.Y + 1} / ${flamePos.X + 1}"]`);
        if (flameElement) {
            flameElement.classList.remove('destructible');
            if ((flamePos.X === centerFlamePos.X) && (flamePos.Y === centerFlamePos.Y)) {
                flameElement.classList.add('flame-center');
            }
            if (flamePos.Y > centerFlamePos.Y) {
                flameElement.classList.add('flame-down');
            }
            if (flamePos.Y < centerFlamePos.Y) {
                flameElement.classList.add('flame-up');
            }
            if (flamePos.X > centerFlamePos.X) {
                flameElement.classList.add('flame-right');
            }
            if (flamePos.X < centerFlamePos.X) {
                flameElement.classList.add('flame-left');
            }

            setTimeout(() => {
                flameElement.classList.remove('flame-up', 'flame-down', 'flame-right', 'flame-left', 'flame-center');
                updateField(flamePos, pos[i].FieldCode);
            }, 1000);
        }
    }
}

export function removePlayerFromGame(playerID) {
    // Find the player in the game
    const oldPlayerElement = document.getElementById('player' + playerID)
    // If found, remove the class from the old position element
    console.log('Player ' + playerID + ' has left the game, removing element')
    if (oldPlayerElement) {
        // Assuming there should only be one element with this class at a time
        oldPlayerElement.removeAttribute('id');
    }
}

export function handlePlayerLoseLife(payload) {
    for (let i = 0; i < players.length; i++) {
        if (players[i].id === payload.playerID) {
            const playerID = 'player' + players[i].id;
            const player = document.getElementById(playerID);
            if (player) {
                addBlinkingEffect(playerID);
            }
            players[i].lives = payload.lives;
            updateLivesDisplay(players[i].id, players[i].lives);

        }
    }
}

export function disableImmunity(playerID) {
    removeBlinkingEffect('player' + playerID);
}

function addBlinkingEffect(playerID) {
    const element = document.getElementById(playerID);
    if (element) {
        element.classList.add('blinking');
    }
}

function removeBlinkingEffect(playerID) {
    const element = document.getElementById(playerID);
    if (element) {
        element.classList.remove('blinking');
    }
}
