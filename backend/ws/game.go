package ws

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

// func handleGameInput(client *Client, input wsMessage) {
// 	game, ok := games[client.GameID]
// 	if !ok {
// 		log.Println("Game not found for input handling")
// 		return
// 	}

// 	// Process input here (e.g., updating player's position based on the direction)
// 	// Assuming input.Payload contains movement information like {Direction: "up"}
// 	direction := input.Payload.(map[string]interface{})["Direction"].(string)
// 	processPlayerMovement(game, client.ID, direction)
// }
