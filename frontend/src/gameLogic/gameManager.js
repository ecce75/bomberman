function gameLoop() {
    updateGameState();  // Update game logic, positions, etc.
    render();           // Draw the current game state to the screen

    requestAnimationFrame(gameLoop); // Loop this function
}



function updateGameState() {
    // Handle player movement
    // Check for collisions
    // Update object positions and check game logic
    // Possibly update the score or game status
}

function render() {
    // Clear the screen
    // Draw all game objects
    // Update the DOM elements if necessary (like player positions or animations)
}
