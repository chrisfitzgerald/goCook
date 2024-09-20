const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

let gameStarted = false;
let gameOver = false;
let score = 0;

let block = {
    x: 50,
    y: canvas.height - 60,
    width: 30,
    height: 30,
    dy: 0,
    gravity: 0.3,
    jumpPower: -10,
    isJumping: false
};

let platforms = [];

// Create initial platforms
function createPlatforms() {
    platforms = [];
    for (let i = 0; i < 5; i++) {
        platforms.push({
            x: canvas.width + i * 200,
            y: Math.random() * (canvas.height - 100) + 50,
            width: 100,
            height: 10
        });
    }
}

// Title screen
function drawTitleScreen() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.fillStyle = 'black';
    ctx.font = '48px Arial';
    ctx.fillText('Super Fun Time', canvas.width / 2 - 150, canvas.height / 2 - 50);
    ctx.font = '24px Arial';
    ctx.fillText('Press Space Bar to Start', canvas.width / 2 - 130, canvas.height / 2);
    ctx.fillText('Controls:', canvas.width / 2 - 50, canvas.height / 2 + 50);
    ctx.fillText('Space Bar - Jump', canvas.width / 2 - 80, canvas.height / 2 + 80);
}

// Game over screen
function drawGameOverScreen() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.fillStyle = 'black';
    ctx.font = '48px Arial';
    ctx.fillText('Game Over', canvas.width / 2 - 120, canvas.height / 2 - 50);
    ctx.font = '24px Arial';
    ctx.fillText('Your Score: ' + score, canvas.width / 2 - 70, canvas.height / 2);
    ctx.fillText('Press Space Bar to Restart', canvas.width / 2 - 140, canvas.height / 2 + 50);
}

// Game loop
function gameLoop() {
    if (!gameStarted) {
        drawTitleScreen();
        return;
    }

    if (gameOver) {
        drawGameOverScreen();
        return;
    }

    update();
    draw();
    requestAnimationFrame(gameLoop);
}

// Update game state
function update() {
    // Apply gravity
    block.dy += block.gravity;
    block.y += block.dy;

    // Prevent block from going off the top
    if (block.y < 0) {
        block.y = 0;
        block.dy = 0;
    }

    // Check for collision with platforms
    for (let platform of platforms) {
        if (
            block.x < platform.x + platform.width &&
            block.x + block.width > platform.x &&
            block.y + block.height > platform.y &&
            block.y < platform.y + platform.height &&
            block.dy >= 0
        ) {
            block.y = platform.y - block.height;
            block.dy = 0;
            block.isJumping = false;
        }
    }

    // Check for ground collision
    if (block.y + block.height >= canvas.height) {
        block.y = canvas.height - block.height;
        block.dy = 0;
        block.isJumping = false;
    }

    // Move platforms to the left
    for (let platform of platforms) {
        platform.x -= 2; // Scroll speed
        if (platform.x + platform.width < 0) {
            platform.x = canvas.width;
            platform.y = Math.random() * (canvas.height - 100) + 50;
            score++;
        }
    }

    // Check for game over (block falls below canvas)
    if (block.y > canvas.height) {
        gameOver = true;
    }
}

// Draw everything
function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // Draw block
    ctx.fillStyle = 'red';
    ctx.fillRect(block.x, block.y, block.width, block.height);

    // Draw platforms
    ctx.fillStyle = 'green';
    for (let platform of platforms) {
        ctx.fillRect(platform.x, platform.y, platform.width, platform.height);
    }

    // Draw score
    ctx.fillStyle = 'black';
    ctx.font = '20px Arial';
    ctx.fillText('Score: ' + score, 10, 30);
}

// Handle key presses
window.addEventListener('keydown', function (e) {
    if (e.code === 'Space') {
        if (!gameStarted) {
            gameStarted = true;
            gameOver = false;
            score = 0;
            block.y = canvas.height - 60;
            block.dy = 0;
            createPlatforms();
            gameLoop();
        } else if (gameOver) {
            // Restart the game
            gameOver = false;
            score = 0;
            block.y = canvas.height - 60;
            block.dy = 0;
            createPlatforms();
            gameLoop();
        } else if (!block.isJumping) {
            block.dy = block.jumpPower;
            block.isJumping = true;
        }
    }
});

// Start the game
createPlatforms();
gameLoop();
