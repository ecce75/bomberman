package ws

import "fmt"

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

func (gm *Game) BroadcastPlayerDisconnected(name string, playerID string) {
	for _, player := range gm.Players {
		fmt.Println("Broadcasting player left: "+ name)
		player.Conn.WriteJSON(wsMessage{Type: "playerLeft", Payload: map[string]interface{}{
			"playerID": playerID,
			"name":     name,
		}})
	}
}