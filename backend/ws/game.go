package ws

import (
	"github.com/google/uuid"
	"log"
)

func initializeGame(lobby *Lobby) {
	gameID := uuid.New().String()
	newGame := &Game{
		ID:      gameID,
		Players: make(map[string]*Client),
	}

	for _, client := range lobby.Players {
		newGame.Players[client.ID] = client
	}

	games[gameID] = newGame
	delete(lobbies, lobby.ID) // Optionally remove lobby if no longer needed

	// Broadcast game start
	startGameMsg := wsMessage{
		Type:    "startGame",
		Payload: "Game Starting",
	}
	for _, cl := range newGame.Players {
		if err := cl.Conn.WriteJSON(startGameMsg); err != nil {
			log.Printf("error: failed to start game for client %s", cl.ID)
			handleDisconnect(cl.ID)
		}
	}
}

func handleRestartGame(client *Client) {
	lobby, ok := lobbies[client.LobbyID]
	if !ok {
		log.Printf("Lobby not found for client %s", client.ID)
		return
	}

	// Restart game logic
	initializeGame(lobby)
}

func startGame(lobby *Lobby) {
	gameID := uuid.New().String()
	newGame := &Game{
		ID:      gameID,
		Players: lobby.Players,
		Timer:   nil, // No timer for now
	}
	games[gameID] = newGame
	for _, client := range newGame.Players {
		client.Conn.WriteJSON(wsMessage{Type: "gameStart", Payload: "The game has started!"})
	}
	delete(lobbies, lobby.ID) // Remove lobby once game starts
}

func handleGameInput(client *Client, input wsMessage) {
	game, ok := games[client.LobbyID]
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

func placeBomb(game *Game, clientID string, position map[string]int) {
	player, ok := game.Players[clientID]
	if !ok {
		log.Println("Player not found in the game for bomb placement")
		return
	}
	println(player)

	// Assume we get x, y coordinates for bomb placement
	x, y := position["x"], position["y"]

	// Validate bomb placement, update game state, start explosion timer etc.
	// Bomb logic here
	log.Printf("Bomb placed by %s at (%d, %d)", clientID, x, y)
	// Implement bomb explosion timer and effect
}

func handleBombPlaced(client *Client, payload interface{}) {
	game, ok := games[client.LobbyID]
	if !ok {
		log.Println("Game not found for bomb placement")
		return
	}

	// Assuming payload contains bomb placement coordinates
	position := payload.(map[string]int)
	placeBomb(game, client.ID, position)
}
