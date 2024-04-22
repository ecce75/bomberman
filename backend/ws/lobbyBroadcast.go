package ws

import (
	"log"
	"time"

	"github.com/google/uuid"
)

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
	client.GameID = lobby.ID
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
