
export function updatePlayerPosition(playerID, newPosition) {
    console.log('updating player position', newPosition)
    const  oldPlayerElement = document.getElementById('player' + playerID)

    if (oldPlayerElement) {
        oldPlayerElement.removeAttribute('id');
    }

    const newSelector = `.cell[style*="grid-area: ${newPosition.Y + 1} / ${newPosition.X + 1}"]`; // Plus 1 to match 1-based indexing
    const newPlayerElement = document.querySelector(newSelector);
    if (newPlayerElement) {
        newPlayerElement.id = 'player' + playerID;
    }else {
        console.log('no element found')
    }
}
