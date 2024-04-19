package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	ID      string
	Conn    *websocket.Conn
	LobbyID string
}

type Game struct {
	ID      string
	Players map[string]*Client
	Timer   *time.Timer
}

type Lobby struct {
	ID      string
	Players map[string]*Client
	Timer   *time.Timer
}

var clients = make(map[string]*Client)
var lobbies = make(map[string]*Lobby)
var games = make(map[string]*Game)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	clientID := uuid.New().String()
	newClient := &Client{ID: clientID, Conn: ws}
	clients[clientID] = newClient

	// Example of adding to lobby
	// This needs actual implementation
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
	case "chatMessage":
		// Handle chat message
	case "gameInput":
		// Handle game input
	case "restartGame":
		// Handle restart game
	case "bombPlaced":
		// Handle bomb placement
	}
}

func addToLobby(client *Client) {
	var lobby *Lobby
	for _, l := range lobbies {
		if len(l.Players) < 4 { // Assuming max players per lobby is 4
			lobby = l
			break
		}
	}

	if lobby == nil {
		lobbyID := uuid.New().String()
		lobby = &Lobby{
			ID:      lobbyID,
			Players: make(map[string]*Client),
			Timer:   time.NewTimer(5 * time.Minute), // Auto-start game after 5 minutes if not full
		}
		// lobby.Timer.C = make(chan<- time.Time, 1)
		go lobbyTimeout(lobby)
		lobbies[lobbyID] = lobby
	}

	client.LobbyID = lobby.ID
	lobby.Players[client.ID] = client
	broadcastLobbyStatus(lobby)

	if len(lobby.Players) == 4 {
		startGame(lobby)
	}
}

func lobbyTimeout(lobby *Lobby) {
	<-lobby.Timer.C
	if len(lobby.Players) > 0 && len(lobby.Players) < 4 {
		startGame(lobby) // Start the game even if not full
	}
}

func broadcastLobbyStatus(lobby *Lobby) {
	countMsg := wsMessage{
		Type:    "updateCounter",
		Payload: len(lobby.Players),
	}
	for _, cl := range lobby.Players {
		err := cl.Conn.WriteJSON(countMsg)
		if err != nil {
			log.Printf("error: failed to send lobby status to client %s", cl.ID)
			handleDisconnect(cl.ID)
		}
	}
}

func handleDisconnect(clientID string) {
	if client, ok := clients[clientID]; ok {
		lobby := lobbies[client.LobbyID]
		delete(lobby.Players, clientID)
		delete(clients, clientID)
		if len(lobby.Players) == 0 {
			delete(lobbies, lobby.ID)
		} else {
			broadcastLobbyStatus(lobby)
		}
	}
}

// Define wsMessage struct as per your application's requirements
type wsMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
