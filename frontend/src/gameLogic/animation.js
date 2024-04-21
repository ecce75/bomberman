const canvas = document.getElementById(`cell player${index + 1}`);
const ctx = canvas.getContext('2d');
const CANVAS_WIDTH = canvas.width = 60;
const CANVAS_HEIGHT = canvas.height = 80;

const playerImage = new Image();
playerImage.src = '../img/sprites.png';
const spriteWidth = 52;
const spriteHeight = 73;
let frameX = 0;
let frameY = 0;
let gameFrame = 0;
const staggerFrames = 5;

export function animate() {
    ctx.clearRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);
    ctx.fillRect(100,50,100,100);
    ctx.drawImage(playerImage, frameX * spriteWidth, frameY * spriteHeight,
        spriteWidth, spriteHeight, 0, 0, spriteWidth, spriteHeight);
    if  (gameFrame % staggerFrames == 0) {
        if (frameX < 6) frameX++;
        else frameX = 0;
    }

    gameFrame++;
    requestAnimationFrame(animate);
};
animate();