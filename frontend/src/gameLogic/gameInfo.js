import { createStructure } from "/framework.js";
import { players } from "./gameSetup.js";

export function createTimerDisplay() {
    // Create the timer display structure
    const timerDisplay = createStructure({
        tag: 'div',
        attr: ['class', 'timer-display'],
        children: [
            { tag: 'h3' },
            createStructure({
                tag: 'div',
                attr: ['class', 'timer-flex-container'], // Added class for styling
                children: [
                    { tag: 'p', attr: ['id', 'minutes', 'class', 'time-element'], children: '00' },
                    { tag: 'p', attr: ['class', 'time-element'], children: ':' },
                    { tag: 'p', attr: ['id', 'seconds', 'class', 'time-element'], children: '00' }
                ]
            })
        ]
    });

    // Append the timer display to the body or any specific element you prefer
    document.body.appendChild(timerDisplay);

    // Start the timer
    startTimer();

    return timerDisplay;  // Optionally return this if you need to manipulate it later
}

// Define the startTimer function to update the timer display
function startTimer() {
    let totalSeconds = 0;
    const minutesDisplay = document.getElementById('minutes');
    const secondsDisplay = document.getElementById('seconds');

    setInterval(() => {
        totalSeconds++;
        const minutes = Math.floor(totalSeconds / 60);
        const seconds = totalSeconds % 60;
        minutesDisplay.textContent = pad(minutes);
        secondsDisplay.textContent = pad(seconds);
    }, 1000);

    function pad(num) {
        return num.toString().padStart(2, '0');
    }
}

export function createPlayersDisplay() {
    return createStructure({
        tag: 'div',
        attr: ['class', 'players-display'],
        children: players.map(player => createPlayerDisplay(player))
    });
}

function createPlayerDisplay(player) {
    const playerIndex = player.id;

    return createStructure({
        tag: 'div',
        attr: ['class', 'player-info'],
        children: [
            createStructure({
                tag: 'img',
                attr: [
                    'src', `../public/images/player${playerIndex}.png`, // Use a default image if none is provided
                    'alt', 'Player Avatar',
                    'class', 'player-avatar' // Optional: add a class for styling
                ]
            }),
            createStructure({
                tag: 'div',
                attr: ['class', 'player-details'],
                children: [
                    { tag: 'h4', children: player.username },
                    createStructure({
                        tag: 'div',
                        attr: ['class', 'player-lives'],
                        children: [
                            { tag: 'img', attr: ['src', `../public/images/pixelheart.png`, 'alt', 'Lives Icon'] },
                            { tag: 'span',attr: ['id', 'lives'+player.id], children: `x${player.lives}` }
                        ]
                    })
                ]
            }),
            createPowerupsDisplay(player.powerups)
        ]
    });
}

export function updateLivesDisplay(playerID, lives) {
    const livesDisplay = document.getElementById('lives' + playerID);
    if (livesDisplay) {
        livesDisplay.textContent = `x${lives}`;
    }

}

export function createPowerupsDisplay(powerups) {
    // Check if the powerups object has any properties
    if (Object.keys(powerups).length === 0) {
        // Return an empty structure or a message if no powerups are present
        return createStructure({
            tag: 'div',
            children: 'No powerups available'
        });
    }

    // If powerups are present, create a list to display them
    return createStructure({
        tag: 'ul',
        attr: ['class', 'powerups-list'], // Optionally add a class for styling
        children: Object.entries(powerups).filter(([key, value]) => value === 2).map(([key, value]) => {
            let childrenDescription = `${key.charAt(0).toUpperCase() + key.slice(1)}: ${value}`;
            let iconPath = '';

            // Determine the correct icon path based on the type of powerup
            switch (key) {
                case 'bomb':
                    iconPath = '../public/images/bomb_powerup.png';
                    break;
                case 'flamerange':
                    iconPath = '../public/images/flamerange_powerup.png';
                    break;
                case 'speed':
                    iconPath = '../public/images/speed_powerup.png';
                    break;
                default:
                    iconPath = ''; // No icon for other types
                    break;
            }

            if (iconPath) {
                childrenDescription = [
                    { tag: 'img', attr: ['src', iconPath, 'alt', `${key} Powerup`] }
                ];
            }

            return createStructure({
                tag: 'li',
                children: childrenDescription
            });
        })
    });
}



