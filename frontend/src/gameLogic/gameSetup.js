import {createChild, createStructure} from "/framework.js";

export let players = [];

export function createChatbox() {
    return createStructure({
        tag: 'div',
        attr: ['class', 'chatbox'],
        children: [
            {
                tag: 'h3',
                children: 'Messages'
            },
            {
                tag: 'div',
                attr: ['id', 'chatMessages']
            },
            {
                tag: 'input',
                attr: ['type', 'text', 'id', 'messageInput', 'placeholder', 'Type your message...']
            },
            {
                tag: 'button',
                attr: ['id', 'sendButton'],
                children: 'Send'
            }
        ]
    });
}


export function createGameBoard(game) {
    const map = game.map; // Get the map from the game object
    const gamePlayers = game.players; // Get the players from the game object
    console.log(gamePlayers);
    for (const player of gamePlayers) {
        addPlayer(player); // Add each player to the players array
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
            if (map[i][j] === 1) {
                cellClass += ' indestructible'; // Add indestructible obstacles at every other row and column
            }
            if (map[i][j] === 2) {
                cellClass += ' destructible'; // Add destructible obstacles at every other row and column
            }
            if (map[i][j] === 3) {
                cellClass += ' player1'; // Add player1 at the first cell
                addPlayer(i, j, 1); // Add player1 to the players array
            }
            if (map[i][j] === 4) {
                cellClass += ' player2'; // Add player2 at the second cell
                addPlayer(i, j, 2); // Add player1 to the players array
            }
            if (map[i][j] === 5) {
                cellClass += ' player3'; // Add player3 at the third cell
                addPlayer(i, j, 3); // Add player1 to the players array
            }
            if (map[i][j] === 6) {
                cellClass += ' player4'; // Add player4 at the fourth cell
                addPlayer(i, j, 1); // Add player1 to the players array
            }
            
            const cell = createStructure({
                tag: 'div',
                attr: ['class', cellClass],
                style: ['grid-column', j + 1, 'grid-row', i + 1],
            });
            createChild(board, cell);

            cellCounter++; // Increment the cell counter

            // Check if it's the 13th cell
            if (cellCounter % 13 === 0 && cellCounter !== 13 * 13) {
                const lineBreak = createStructure({ tag: 'br' }); // Create <br> element
                createChild(board, lineBreak); // Append <br> element
            }
        }
    }
    console.log(players);
    return board;
}


export function createScoreboard() {
    const score = createStructure({
        tag: 'div',
        attr: ['class', 'scoreboard'],
        children: [
            { tag: 'h3', children: 'Timer' },
            { tag: 'p', attr:'minutes' ,children: '00'},
            { tag: 'p', attr:'colon' ,children: ':'},
            { tag: 'p', attr:'seconds' ,children: '00'},
        ]
    });
    return score;
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
            flames: 1,
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