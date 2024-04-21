package ws

func (gm *Game) BroadcastStartGame() {
	var players []*GamePlayer
	for _, player := range gm.Players {
		players = append(players, player.Player)
	}
	for _, player := range gm.Players {
		player.Conn.WriteJSON(wsMessage{Type: "gameStart", Payload: map[string]interface{}{
			"players": players,
			"map":     gm.Map.gameMap,
		}})
	}
}
