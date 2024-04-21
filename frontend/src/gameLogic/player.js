export function removePlayerFromGame(playerID) {
    // Find the player in the game
    const oldPlayerElement = document.getElementById('player' + playerID)
    // If found, remove the class from the old position element
    console.log('Player ' + playerID + ' has left the game, removing element')
    if (oldPlayerElement) {
        // Assuming there should only be one element with this class at a time
        oldPlayerElement.removeAttribute('id');
    }
}