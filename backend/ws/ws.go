package ws

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	clientID := uuid.New().String()
	newClient := &Client{ID: clientID, Conn: ws}
	clients[clientID] = newClient

	ws.SetCloseHandler(func(code int, text string) error {
		log.Printf("WebSocket closed for client %s, code: %d, reason: %s", clientID, code, text)
		handleDisconnect(clientID)
		return nil
	})

	addToLobby(newClient)

	for {
		var msg wsMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, clientID)
			break
		}
		handleMessages(&msg, newClient)
	}
}

func handleMessages(msg *wsMessage, client *Client) {
	switch msg.Type {
	case "setUsername":
		name, ok := msg.Payload.(string)
		if !ok {
			client.Conn.WriteJSON(wsMessage{Type: "invalidUsername", Payload: "Invalid username"})
			return
		}
		client.Name = name
	case "chatMessage":
		// Handle chat message
		handleChatMessage(client, msg.Payload)
	case "move":
		handleMovement(client, *msg)
	case "restartGame":
		// Handle restart game
	case "bomb":
		handlePlaceBomb(client)
	}
}

func handleDisconnect(clientID string) {
	if client, ok := clients[clientID]; ok {
		lobby, ok1 := lobbies[client.GameID]
		game, ok2 := games[client.GameID]

		if !ok1 {
			log.Printf("Lobby not found for client %s", clientID)
			if !ok2 {
				log.Printf("Game not found for client %s", clientID)
				// No lobby or game found so remove player from clients
				delete(clients, clientID)
			}
			playerID := client.Player.ID
			clientName := client.Name
			delete(game.Players, clientID)
			delete(clients, clientID)
			if len(game.Players) == 0 {
				delete(games, game.ID)
			} else {
				game.BroadcastPlayerDisconnected(clientName, playerID)
			}
		} else {
			// Lobby found so remove player from lobby and clients
			delete(lobby.Players, clientID)
			delete(clients, clientID)
			if len(lobby.Players) == 0 {
				delete(lobbies, lobby.ID)
			} else {
				broadcastLobbyStatus(lobby)
			}
		}
	} else {
		log.Printf("Client not found: %s", clientID)
	}
}
