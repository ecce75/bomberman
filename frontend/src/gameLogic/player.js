
export function activateBomb(pos) {
    const bombElement = document.querySelector(`.cell[style*="grid-area: ${pos.Y + 1} / ${pos.X + 1}"]`);

    if (bombElement) {
        bombElement.classList.add('bomb');
    }
    setTimeout(() => {
        bombElement.classList.remove('bomb');
    }, 3000);
}

export function activateFlames(pos) {
    for (let i = 0; i < pos.length; i++) {
        const flameElement = document.querySelector(`.cell[style*="grid-area: ${pos[i].Y + 1} / ${pos[i].X + 1}"]`);
        if (flameElement) {
            console.log('flame activated', pos[i].X, pos[i].Y, pos[0].X, pos[0].Y)
            flameElement.classList.remove('destructible');
            if ((pos[i].X === pos[0].X) && (pos[i].Y === pos[0].Y)) {
                flameElement.classList.add('flame-center');
                console.log('flame center')
            }
            if (pos[i].Y > pos[0].Y) {
                flameElement.classList.add('flame-down');
                console.log('flame up')
            }
            if (pos[i].Y < pos[0].Y) {
                flameElement.classList.add('flame-up');
                console.log('flame down')
            }
            if (pos[i].X > pos[0].X) {
                flameElement.classList.add('flame-right');
                console.log('flame right')
            }
            if (pos[i].X < pos[0].X) {
                flameElement.classList.add('flame-left');
                console.log('flame left')
            }


            setTimeout(() => {
                flameElement.classList.remove('flame-up', 'flame-down', 'flame-right', 'flame-left', 'flame-center');
            }, 1500);
        }
    }
}

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