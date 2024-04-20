// Define a function to create and position player elements on the map
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
        const player = document.createElement('div');
        player.className = 'player';

        // Set player sprite based on index
        switch (index) {
            case 0:
                player.style.backgroundPosition = '0px 0px'; // Player 1, first sprite
                break;
            case 1:
                player.style.backgroundPosition = '-52px 0px'; // Player 2, first sprite
                break;
            case 2:
                player.style.backgroundPosition = '-104px 0px'; // Player 3, first sprite
                break;
            case 3:
                player.style.backgroundPosition = '-156px 0px'; // Player 4, first sprite
                break;
        }

        // Determine the corner to place the player
        const startPosition = corners[index];
        const startCellIndex = cornerCellIndexes[startPosition];
        const startCell = board.children[startCellIndex];
        createChild(startCell, player);
    });
}
