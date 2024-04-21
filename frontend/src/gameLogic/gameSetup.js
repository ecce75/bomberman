import { createChild, createStructure } from "/framework.js";
import { createPlayersDisplay, createTimerDisplay } from "./gameInfo.js";

// Map field codes: {
//     0: "free",
//     1: "indestructible",
//     2: "destructible",
//     3: "player1",
//     4: "player2",
//     5: "player3",
//     6: "player4",
//     7: "bomb",
//     8: "flame",
//     9: "powerup: speed"
//     10: "powerup: flameRange"
//     11: "powerup: bombCount"
//
// }

export let players = [];

export function createGameBoard(game) {
    const map = game.map; // Get the map from the game object
    const gamePlayers = game.players; // Get the players from the game object
    console.log(gamePlayers);
    
    for (const player of gamePlayers) {
        addPlayer(player); // Add the player to the players array

    }
    const board = createStructure({
        tag: 'div',
        attr: ['class', 'game-board'],
        children: []
    });
    let cellCounter = 0; // Counter to track the number of cells

    for (let i = 0; i < 13; i++) {
        for (let j = 0; j < 13; j++) {
            let cellClass = 'cell';
            let cellID = '';

            switch (map[i][j]) {
                case 1:
                    cellClass += ' indestructible';
                    break;
                case 2:
                    cellClass += ' destructible';
                    break;
                case 3:
                    cellID = 'player1';
                    break;
                case 4:
                    cellID = 'player2';
                    break;
                case 5:
                    cellID = 'player3';
                    break;
                case 6:
                    cellID = 'player4';
                    break;
            }

            const cellAttributes = ['class', cellClass];
            if (cellID) {
                cellAttributes.push('id', cellID);
            }

            const cell = createStructure({
                tag: 'div',
                attr: cellAttributes,
                style: ['grid-area', `${i+1 } / ${j + 1}`], // Using 1-based index for grid-area
            });
            createChild(board, cell);
            cellCounter++; // Increment the cell counter

            // Check if it's the 13th cell
            if (cellCounter % 13 === 0 && cellCounter !== 13 * 13) {
                const lineBreak = createStructure({tag: 'br'}); // Create <br> element
                createChild(board, lineBreak); // Append <br> element
            }
        }
    }
    return board;
}

export function createScoreboard() {
    const scoreboard = createStructure({
        tag: 'div',
        attr: ['class', 'scoreboard'],
        children: [

            createTimerDisplay(),
            createPlayersDisplay()
        ]
    });
    return scoreboard;
}


function addPlayer(player) {
    if (player.Position === undefined) {
        return;
    }
    players.push({
        id: player.ID,
        username: player.Username,
        lives: player.Lives,
        powerups: {
            bomb: 1,
            flamerange: 1,
            speed: 1,
        },
        position: {
            x: player.Position.X,
            y: player.Position.Y,
        },
        immunityTimer: null,
        activeBombsPlaced: 0,
    });
}


function placePlayer(board, player, startPosition) {
    const startCell = board.children[startPosition];
    createChild(startCell, player);
}