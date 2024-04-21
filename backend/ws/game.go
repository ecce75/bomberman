package ws

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func startGame(lobby *Lobby) {
	gameMap := NewGameMap(lobby.Players)
	newGame := &Game{
		ID:      lobby.ID,
		Players: lobby.Players,
		Timer:   nil, // No timer for now
		Map:     gameMap,
	}
	games[lobby.ID] = newGame

	newGame.BroadcastStartGame()
	delete(lobbies, lobby.ID) // Remove lobby once game starts
}

func handleMovement(client *Client, input wsMessage) {
	game, ok := games[client.GameID]
	if !ok {
		log.Println("Game not found for input handling")
		return
	}
	// Process input here (e.g., updating player's position based on the direction)
	// Assuming input.Payload contains movement information like {Direction: "up"}
	direction := input.Payload.(string)
	newPosition := game.processPlayerMovement(client.ID, direction)
	game.BroadcastPlayerMovement(client.Player.ID, newPosition)
}

func (gm *Game) processPlayerMovement(clientID string, direction string) Coordinates {
	// Retrieve player
	player, ok := gm.Players[clientID]
	if !ok {
		// Handle case where player is not found
		fmt.Println("Player not found: ", clientID)
		return Coordinates{} // or handle the error as appropriate
	}

	// Calculate new position based on the direction
	newPosition := player.Player.Position
	switch direction {
	case "up":
		newPosition.Y -= 1
	case "down":
		newPosition.Y += 1
	case "left":
		newPosition.X -= 1
	case "right":
		newPosition.X += 1
	}

	// Check if the new position is valid within the game constraints, e.g., not out of bounds
	if !gm.isValidPosition(newPosition) {
		return player.Player.Position // Return the old position if new position is invalid
	}
	if gm.Map.gameMap[newPosition.Y][newPosition.X] == 8 {
		player.Player.LoseLife(gm) // Return the old position if new position is invalid
	}

	// Update player's position in the game structure
	gm.Players[clientID].Player.Position = newPosition

	return newPosition
}

// Example of a validation function, assuming the game has bounds or other conditions
func (gm *Game) isValidPosition(pos Coordinates) bool {
	// Example checks, these will depend on your game's specific logic and boundaries
	return pos.X >= 0 && pos.X < gm.Map.mapWidth && pos.Y >= 0 && pos.Y < gm.Map.mapHeight && gm.Map.gameMap[pos.Y][pos.X] != 1 && gm.Map.gameMap[pos.Y][pos.X] != 2
}

func (gm *Game) isValidFlamePosition(pos Coordinates) bool {
	// Example checks, these will depend on your game's specific logic and boundaries
	return pos.X >= 0 && pos.X < gm.Map.mapWidth && pos.Y >= 0 && pos.Y < gm.Map.mapHeight && gm.Map.gameMap[pos.Y][pos.X] != 1
}

func (gm *Game) activateFlames(position Coordinates, flameRange int) {
	// Set initial flame at position
	gm.activateFlameAt(position)

	var flames []Coordinates
	directions := []Coordinates{{0, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}} // represents right, down, left, up

	for i := 1; i <= flameRange; i++ {
		for _, dir := range directions {
			newPos := Coordinates{X: position.X + i*dir.X, Y: position.Y + i*dir.Y}
			if gm.isValidFlamePosition(newPos) {
				gm.activateFlameAt(newPos)
				flames = append(flames, newPos)
			}
		}
	}

	gm.processFlameEffects(flames)
	gm.BroadcastFlames(flames)
}

func (gm *Game) activateFlameAt(position Coordinates) {
	gm.Map.gameMap[position.Y][position.X] = 8 // Assuming 8 represents an active flame
}

func (gm *Game) processFlameEffects(flames []Coordinates) {
	for _, flame := range flames {
		flameCode := gm.Map.gameMap[flame.Y][flame.X]
		if flameCode >= 3 && flameCode <= 6 { // Assuming these codes correspond to players
			gm.triggerPlayerDamage(flameCode)
		}
	}
}

func (gm *Game) triggerPlayerDamage(playerCode int) {
	for _, player := range gm.Players {
		if player.Player.ID == strconv.Itoa(playerCode-2) { // Assumes ID "1" for code 3, "2" for code 4, etc.
			player.Player.LoseLife(gm)
		}
	}
}

func (gm *Game) generatePowerUp(position Coordinates) {
	rand.Seed(time.Now().UnixNano())
	numbers := []int{9, 10, 10, 11, 11}
	number := numbers[rand.Intn(len(numbers))]
	if gm.Map.gameMap[position.Y][position.X] == 2 {
		gm.Map.gameMap[position.Y][position.X] = number
	}
}
