package ws

// BroadcastPlayerDamage is a placeholder for broadcasting player damage
func (gm *Game) BroadcastPlayerDamage(damagedPlayer *GamePlayer) {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerLoseLife", Payload: map[string]interface{}{
			"playerID": damagedPlayer.ID,
			"lives":    damagedPlayer.Lives,
		}})
	}
}

// CheckGameOver is a placeholder to check if the game is over
func (gm *Game) CheckGameOver() {
	// check if there are more than 1 players left with lives
	alivePlayers := 0
	// get last alive player name
	lastAlivePlayer := ""
	for _, player := range gm.Players {
		if player.Player.Lives > 0 {
			alivePlayers++
			lastAlivePlayer = player.Name
		}

	}

	if alivePlayers == 1 {
		for _, player := range gm.Players {
			player.Conn.WriteJSON(wsMessage{Type: "gameOver", Payload: map[string]interface{}{
				"gameOver": true,
				"winner":   lastAlivePlayer,
			}})
		}
	}
}

// BroadcastImmunityEnd is a placeholder for broadcasting end of immunity
func (gm *Game) BroadcastImmunityEnd(playerID string) {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"playerID":    playerID,
			"immunityEnd": true,
		}})
	}
}

func (gm *Game) BroadcastPlayerMovement(playerID string, coordinates Coordinates) {
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "playerMovement", Payload: map[string]interface{}{
			"playerID":    playerID,
			"newPosition": coordinates,
		}})
	}
}
