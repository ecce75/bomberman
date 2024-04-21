package ws

// BroadcastPlayerDamage is a placeholder for broadcasting player damage
func (gm *Game) BroadcastPlayerDamage(playerIndex int, damagedPlayer *GamePlayer) {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"playerIndex": playerIndex,
			"lives":       damagedPlayer.Lives,
		}})
	}
}

// CheckGameOver is a placeholder to check if the game is over
func (gm *Game) CheckGameOver() {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"gameOver": true,
		}})
	}
}

// BroadcastImmunityEnd is a placeholder for broadcasting end of immunity
func (gm *Game) BroadcastImmunityEnd(playerIndex int) {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"playerIndex": playerIndex,
			"immunityEnd": true,
		}})
	}
}

func (gm *Game) BroadcastPlayerMovement(playerIndex int, coordinates Coordinates) {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"playerIndex": playerIndex,
			"coordinates": coordinates,
		}})
	}
}
