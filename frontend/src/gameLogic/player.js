import { createChild, createStructure } from '../framework.js';

export function createPlayersAndPlace() {
    // Define player indices
    const players = [
        { index: 0 }, // Player 1
        { index: 1 }, // Player 2
        { index: 2 }, // Player 3
        { index: 3 }  // Player 4
    ];

    // Define board and corners
    const board = document.getElementById('game-map');
    const corners = ['top-left', 'top-right', 'bottom-left', 'bottom-right'];
    const cornerCellIndexes = {
        'top-left': 0,
        'top-right': 10,
        'bottom-left': 110,
        'bottom-right': 120
    };

    // Create, position, and place player elements
    players.forEach(playerData => {
        const { index } = playerData;
        const player = createPlayerElement(index);

        // Determine the corner to place the player
        const startPosition = corners[index];
        const startCellIndex = cornerCellIndexes[startPosition];
        const startCell = board.children[startCellIndex];
        createChild(startCell, player);
    });
}

function createPlayerElement(index) {
    const player = createStructure({
        tag: 'div',
        attr: ['class', 'player'],
        style: ['backgroundPosition', `${-52 * index}px 0px`] // Assuming 52px width per sprite
    });

    return player;
}
