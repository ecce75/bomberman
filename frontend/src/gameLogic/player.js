// Define a function to create and position player elements on the map
function createPlayers() {
    // Define player coordinates and indices
    const players = [
        { x: 100, y: 100, index: 0 }, // Player 1
        { x: 200, y: 100, index: 1 }, // Player 2
        { x: 100, y: 200, index: 2 }, // Player 3
        { x: 200, y: 200, index: 3 }  // Player 4
    ];

    // Create and position player elements
    players.forEach(player => {
        createPlayer(player.x, player.y, player.index);
    });
}

// Define a function to create and position a player element on the map
function createPlayer(x, y, index) {
    // Create a new player element
    var player = document.createElement('div');
    player.className = 'player';

    // Position the player based on coordinates
    player.style.left = x + 'px';
    player.style.top = y + 'px';

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

    // Add player element to the game map container
    document.getElementById('game-map').appendChild(player);
}

// Call the function to create and position all players
createPlayers();

function placePlayer(board, player, startPosition) {
    const corners = ['top-left', 'top-right', 'bottom-left', 'bottom-right'];
    const cornerCellIndexes = {
        'top-left': 0,
        'top-right': 10,
        'bottom-left': 110,
        'bottom-right': 120
    };

    const startCellIndex = cornerCellIndexes[corners[startPosition]];
    const startCell = board.children[startCellIndex];
    createChild(startCell, player);
}