import {lobbyView} from "/views/lobbyView.js";
import {gameView} from "/views/gameView.js";
import {updatePlayerPosition} from "../gameLogic/movement.js";
import { removePlayerFromGame, handlePlayerLoseLife } from "../gameLogic/player.js";
import { setupChat, handleChatMessage, broadcastPlayerDisconnect } from "../gameLogic/gameChat.js";
import { normalizeField } from "../gameLogic/mapEdit.js";
import {activateBomb, activateFlames, disableImmunity} from "../gameLogic/player.js";
import { updatePlayerPowerupsDisplay } from "../gameLogic/gameInfo.js";

function setupWebSocket() {
    const ws = new WebSocket('ws://localhost:8080/ws'); // Adjust this URL to your server
    ws.onmessage = function(event) {
        const msg = JSON.parse(event.data);
        switch(msg.type) {
            case 'updateCounter':
                document.getElementById('playerCount').textContent = msg.payload.toString();
                break;
            case 'updateTime':
                document.getElementById('countdown').textContent = msg.payload.toString();
                break;
            case 'gameStart':
                gameView(msg.payload, ws);
                setupChat(ws);
                break;
            case 'playerMovement':
                console.log(msg.payload)
                if (msg.payload.newPosition != undefined) {
                updatePlayerPosition(msg.payload.playerID, msg.payload.newPosition);
                }
                break;
            case 'bomb':
                console.log(msg.payload)
                activateBomb(msg.payload);
                break;
            case 'flames':
                console.log(msg.payload)
                activateFlames(msg.payload);
                break;
            case 'fieldUpdate':
                normalizeField(msg.payload);
                break;
            case 'playerPowerup':
                // update player powerups
                updatePlayerPowerupsDisplay(msg.payload);
                break;
            case "invalidUsername":
                alert("Username already taken")
                window.reload()
                break;
            // Handle other messages
            case "chatMessage":
                // handle incoming chat messages
                handleChatMessage(msg.payload);
                break;
            case "playerLeft":
                // handle player leaving
                removePlayerFromGame(msg.payload.playerID);
                // alert that player has left
                broadcastPlayerDisconnect(msg.payload.name);
                break;
            case "playerLoseLife":
                // handle player losing life
                handlePlayerLoseLife(msg.payload);
                break;
            case "gameOver":
                // handle game over
                alert("Game Over! Winner: " + msg.payload.winner);
                window.location.reload();
                break;
            case 'immunity':
                disableImmunity(msg.payload.playerID);
                break;
        }
    };
    return ws;
}

export function submitUsername() {
    const username = document.getElementById('username').value.trim();
    if (!username) {
        alert('Please enter a username.');
        return;
    }
    const ws = setupWebSocket();
    ws.onopen = function() {
        ws.send(JSON.stringify({ type: 'setUsername', payload: username }));
        sessionStorage.setItem("username", username)
        lobbyView();
    };
}


