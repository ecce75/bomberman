package ws

import (
	"fmt"
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
	Name    string
	Conn    *websocket.Conn
	LobbyID string
}

type GameWs struct {
	ID      string
	Players map[string]*Client
	Timer   *time.Timer
}

type Lobby struct {
	ID       string
	Players  map[string]*Client
	TimeLeft int
}

var clients = make(map[string]*Client)
var lobbies = make(map[string]*Lobby)
var games = make(map[string]*GameWs)

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
			ID:       lobbyID,
			Players:  make(map[string]*Client),
			TimeLeft: 2,
		}
		lobbies[lobbyID] = lobby
	}
	client.LobbyID = lobby.ID
	lobby.Players[client.ID] = client
	if len(lobby.Players) == 2 { // Assuming countdown starts when the first player joins
		go lobbyCountdown(lobby)
	}
	broadcastLobbyStatus(lobby)

	if len(lobby.Players) == 4 {
		startGame(lobby)
	}
}

func lobbyCountdown(lobby *Lobby) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if len(lobby.Players) < 2 {
			ticker.Stop()
		}
		lobby.TimeLeft--
		broadcastTimeLeft(lobby)

		if lobby.TimeLeft <= 0 {
			ticker.Stop()
			startGame(lobby)
			return
		}
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

func broadcastTimeLeft(lobby *Lobby) {
	timeMsg := wsMessage{
		Type:    "updateTime",
		Payload: lobby.TimeLeft,
	}
	for _, client := range lobby.Players {
		if err := client.Conn.WriteJSON(timeMsg); err != nil {
			log.Printf("error: failed to send time update to client %s: %v", client.ID, err)
			handleDisconnect(client.ID)
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
