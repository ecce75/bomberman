import {createChild, createStructure} from "/framework.js";

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

// Documentation for map field codes: {
//     0: "free",
//     1: "indestructible",
//     2: "destructible",
//     3: "player1",
//     4: "player2",
//     5: "player3",
//     6: "player4",
//     7: "bomb",
//     8: "booked" // for development purposes
//     9: "powerup: speed"
//     10: "powerup: explosion length"
//     11: "powerup: bombCount"
//     9: "flame"
// }

export function createGameBoard(map) {
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
            }
            if (map[i][j] === 4) {
                cellClass += ' player2'; // Add player2 at the second cell
            }
            if (map[i][j] === 5) {
                cellClass += ' player3'; // Add player3 at the third cell
            }
            if (map[i][j] === 6) {
                cellClass += ' player4'; // Add player4 at the fourth cell
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


/*
function createPlayer() {
    return createStructure({
        tag: 'div',
        attr: ['class', 'player'],
    });
}

function placePlayer(board, player, startPosition) {
    const startCell = board.children[startPosition];
    createChild(startCell, player);
}*/