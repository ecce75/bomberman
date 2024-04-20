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

export function createGameBoard(rows, columns) {
    const board = createStructure({
        tag: 'div',
        attr: ['class', 'game-board'],
        children: []
    });

    let cellCounter = 0; // Counter to track the number of cells


    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < columns; j++) {
            let cellClass = 'cell';
            if ((i % 2 !== 0) && (j % 2 !== 0)) {
                cellClass += ' indestructible'; // Add indestructible obstacles at every other row and column
            }
            const cell = createStructure({
                tag: 'div',
                attr: ['class', cellClass],
                style: ['grid-column', j + 1, 'grid-row', i + 1],
            });
            createChild(board, cell);

            cellCounter++; // Increment the cell counter

            // Check if it's the 11th cell
            if (cellCounter % 11 === 0 && cellCounter !== rows * columns) {
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

