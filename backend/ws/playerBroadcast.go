package ws

// BroadcastPlayerDamage is a placeholder for broadcasting player damage
func BroadcastPlayerDamage(gameID string, playerIndex int, player *GamePlayer) {
	// Actual implementation needed here
}

// CheckGameOver is a placeholder to check if the game is over
func CheckGameOver(gameID string) {
	// Actual implementation needed here
}

// BroadcastImmunityEnd is a placeholder for broadcasting end of immunity
func BroadcastImmunityEnd(gameID string, playerIndex int) {
	// Actual implementation needed here
}

func BroadcastPlayerMovement(gameID string, playerIndex int, direction string) {
	for _, player := range games[gameID].Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"playerIndex": playerIndex,
			"coordinates": direction,
		}})
	}
}
