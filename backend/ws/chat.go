package ws

import "log"

func handleChatMessage(client *Client, payload interface{}) {
	lobby, ok := lobbies[client.LobbyID]
	if !ok {
		log.Printf("Lobby not found for client %s", client.ID)
		return
	}

	chatMsg := wsMessage{
		Type:    "chatMessage",
		Payload: payload,
	}

	// Broadcasting chat message to all clients in the same lobby
	for _, cl := range lobby.Players {
		if err := cl.Conn.WriteJSON(chatMsg); err != nil {
			log.Printf("error: failed to send chat message to client %s", cl.ID)
			handleDisconnect(cl.ID)
		}
	}
}
