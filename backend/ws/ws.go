package ws

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[string]*Client)
var lobbies = make(map[string]*Lobby)
var games = make(map[string]*Game)

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
		fmt.Println(name, ok, msg.Payload)
		if !ok {
			client.Conn.WriteJSON(wsMessage{Type: "invalidUsername", Payload: "Invalid username"})
			return
		}
		client.Name = name
	case "chatMessage":
		// Handle chat message
		handleChatMessage(client, msg.Payload)
	case "gameInput":
		handleGameInput(client, *msg)
	case "restartGame":
		// Handle restart game
	case "bombPlaced":
		// Handle bomb placement
	}
}

func handleDisconnect(clientID string) {
	if client, ok := clients[clientID]; ok {
		lobby := lobbies[client.GameID]
		delete(lobby.Players, clientID)
		delete(clients, clientID)
		if len(lobby.Players) == 0 {
			delete(lobbies, lobby.ID)
		} else {
			broadcastLobbyStatus(lobby)
		}
	}
}
