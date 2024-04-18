// index.js
document.addEventListener('DOMContentLoaded', function() {
    const gameBoard = createGameBoard(11, 11); // Create an 11x11 board
    document.body.appendChild(gameBoard); // Append the game board to the body

    const player = createPlayer();
    placePlayer(gameBoard, player, 0); // Start the player in the first cell

    // Setup event listeners for player movement
    document.addEventListener('keydown', function(event) {
        handlePlayerMovement(event, player);
    });
});

function handlePlayerMovement(event, player) {
    // Logic to move player based on key press (left, right, up, down)
    switch (event.keyCode) {
        case 37: // Arrow Left
            movePlayerLeft(player);
            break;
        case 38: // Arrow Up
            movePlayerUp(player);
            break;
        case 39: // Arrow Right
            movePlayerRight(player);
            break;
        case 40: // Arrow Down
            movePlayerDown(player);
            break;
    }
}
