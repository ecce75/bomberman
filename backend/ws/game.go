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



func handleGameInput(client *Client, input wsMessage) {
	game, ok := games[client.GameID]
	if !ok {
		log.Println("Game not found for input handling")
		return
	}
	// Process input here (e.g., updating player's position based on the direction)
	// Assuming input.Payload contains movement information like {Direction: "up"}
	direction := input.Payload.(string)
	newPosition := processPlayerMovement(game, client.ID, direction)
	game.BroadcastPlayerMovement(client.Player.ID, newPosition)
}

// func processPlayerMovement(game *Game, clientID string, direction string) Coordinates {
// 	// Retrieve player
// 	player := game.Players[clientID]
// 	currentPosition := player.Player.Position
//
// 	switch direction {
// 	case "up":
// 		return Coordinates{currentPosition.X, currentPosition.Y - 1}
// 	case "down":
// 		return Coordinates{currentPosition.X, currentPosition.Y + 1}
// 	case "left":
// 		return Coordinates{currentPosition.X - 1, currentPosition.Y}
// 	case "right":
// 		return Coordinates{currentPosition.X + 1, currentPosition.Y}
// 	}
// 	return currentPosition
// }

func processPlayerMovement(game *Game, clientID string, direction string) Coordinates {
	// Retrieve player
	player, ok := game.Players[clientID]
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
	if !isValidPosition(game, newPosition) {
		return player.Player.Position // Return the old position if new position is invalid
	}

	// Update player's position in the game structure
	game.Players[clientID].Player.Position = newPosition

	return newPosition
}

// Example of a validation function, assuming the game has bounds or other conditions
func isValidPosition(game *Game, pos Coordinates) bool {
	// Example checks, these will depend on your game's specific logic and boundaries
	return pos.X >= 0 && pos.X < game.Map.mapWidth && pos.Y >= 0 && pos.Y < game.Map.mapHeight && game.Map.gameMap[pos.Y][pos.X] != 1 && game.Map.gameMap[pos.Y][pos.X] != 2
}

