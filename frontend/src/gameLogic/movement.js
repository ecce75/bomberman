
export function updatePlayerPosition(playerID, newPosition) {
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

    if (oldPlayerElement.classList.contains('blinking')) {
        oldPlayerElement.classList.remove('blinking');
        newPlayerElement.classList.add('blinking');
    }
}

export function updateField(pos, fieldCode) {
        const cell = document.querySelector(`.cell[style*="grid-area: ${pos.Y + 1} / ${pos.X + 1}"]`);
        if (cell) {
            switch (fieldCode) {
                case 9:
                    cell.classList.add('speedPowerup');
                    break;
                case 10:
                    cell.classList.add('rangePowerup');
                    break;
                case 11:
                    cell.classList.add('bombPowerup');
                    break;
            }
        }
}


