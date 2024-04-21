package ws

import (
	"log"
	"time"
)

func handleChatMessage(client *Client, payload interface{}) {
	game, ok := games[client.GameID]
	if !ok {
		log.Printf("Game not found for client %s", client.ID)
		return
	}

	messagePayload := ChatMessage{Username: client.Name, Message: payload.(string), TimeSent: time.Now().Format("15:04:05")}

	chatMsg := wsMessage{
		Type:    "chatMessage",
		Payload: messagePayload,
	}

	// Add the chat message to the game chat history
	game.ChatMessages = append(game.ChatMessages, messagePayload)

	// Broadcasting chat message to all clients in the same lobby
	for _, cl := range game.Players {
		if err := cl.Conn.WriteJSON(chatMsg); err != nil {
			log.Printf("error: failed to send chat message to client %s", cl.ID)
			handleDisconnect(cl.ID)
		}
	}
}
