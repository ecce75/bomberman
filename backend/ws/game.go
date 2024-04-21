package ws

import (
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
	direction := input.Payload.(map[string]interface{})["Direction"].(string)
	processPlayerMovement(game, client.ID, direction)
}

func processPlayerMovement(game *Game, clientID string, direction string) {
	// Retrieve player
	player, ok := game.Players[clientID]
	if !ok {
		log.Println("Player not found in the game")
		return
	}
	println(player)

	// Example of moving up
	if direction == "up" {
		// Assuming we have coordinates for players
		// Update player coordinates
		// Check if movement is valid
		// player.X, player.Y = new coordinates after moving up
		// Update game state and broadcast to all players
	}

	// Implement other directions and validate the moves
}
