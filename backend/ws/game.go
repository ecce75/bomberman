package ws

import (
	"fmt"
	"log"
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
	// Activate flames at the given position for n steps
	gm.Map.gameMap[position.Y][position.X] = 3 // Assuming 3 represents an active flame
	var flames []Coordinates
	for i := 1; i <= flameRange; i++ {
		// Check if the flame can propagate in each direction
		if gm.isValidFlamePosition(Coordinates{X: position.X, Y: position.Y}) {
			gm.Map.gameMap[position.Y][position.X] = 8
			flames = append(flames, Coordinates{X: position.X, Y: position.Y})
		}
		if gm.isValidFlamePosition(Coordinates{X: position.X + i, Y: position.Y}) {
			gm.Map.gameMap[position.Y][position.X+i] = 8
			flames = append(flames, Coordinates{X: position.X + i, Y: position.Y})
		}
		if gm.isValidFlamePosition(Coordinates{X: position.X - i, Y: position.Y}) {
			gm.Map.gameMap[position.Y][position.X-i] = 8
			flames = append(flames, Coordinates{X: position.X - i, Y: position.Y})
		}
		if gm.isValidFlamePosition(Coordinates{X: position.X, Y: position.Y + i}) {
			gm.Map.gameMap[position.Y+i][position.X] = 8
			flames = append(flames, Coordinates{X: position.X, Y: position.Y + i})
		}
		if gm.isValidFlamePosition(Coordinates{X: position.X, Y: position.Y - i}) {
			gm.Map.gameMap[position.Y-i][position.X] = 8
			flames = append(flames, Coordinates{X: position.X, Y: position.Y - i})
		}
	}
	gm.BroadcastFlames(flames)

}
