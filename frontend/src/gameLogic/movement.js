
function movePlayerX(player, direction) {
    // Move player horizontally based on direction
    if (direction === 'left') {
        player.x -= 1;
    } else if (direction === 'right') {
        player.x += 1;
    }
}

function movePlayerY(player, direction) {
    // Move player vertically based on direction
    if (direction === 'up') {
        player.y -= 1;
    } else if (direction === 'down') {
        player.y += 1;
    }
}

export function updatePlayerPosition(playerID, newPosition) {
    // i want to remove the 'player' + playerID class from the old position
    // and add it to the new position
    const  oldPlayerElement = document.getElementById('player' + playerID)
    // If found, remove the class from the old position element
    console.log('player' + playerID)
    if (oldPlayerElement) {
        // Assuming there should only be one element with this class at a time
        oldPlayerElement.removeAttribute('id');
    }

    console.log(`.cell[style*="grid-area: ${newPosition.Y} / ${newPosition.X}"]`)

    const newSelector = `.cell[style*="grid-area: ${newPosition.Y + 1} / ${newPosition.X + 1}"]`; // Plus 1 to match 1-based indexing
    const newPlayerElement = document.querySelector(newSelector);
    if (newPlayerElement) {
        newPlayerElement.id = 'player' + playerID;
    }else {
        console.log('no element found')
    }
}


