package ws

import (
	"log"
	"time"
)

// NewGamePlayer creates a new player instance
func NewGamePlayer(id string, username string) *GamePlayer {
	return &GamePlayer{
		ID:                id,
		Username:          username,
		Lives:             3,
		Powerups:          Powerups{Bomb: 1, Flames: 1, Speed: 1},
		Position:          Coordinates{X: 0, Y: 0},
		ActiveBombsPlaced: 0,
	}
}

func handlePlaceBomb(client *Client) {
	if client.Player.CanPlaceBomb() {
		game := games[client.GameID]
		client.Player.PlaceBomb(game)
		bombPos := client.Player.Position
		// Decrease the count of active bombs after 3 seconds
		time.AfterFunc(3*time.Second, func() {
			client.Player.ActiveBombsPlaced--
			game.activateFlames(bombPos, client.Player.Powerups.Flames)
		})
	}
}

// CanPlaceBomb checks if a player can place another bomb
func (p *GamePlayer) CanPlaceBomb() bool {
	return p.Powerups.Bomb > p.ActiveBombsPlaced
}

// StartImmunityTimer starts an immunity timer for the player
func (p *GamePlayer) StartImmunityTimer(game *Game, playerIndex int) {
	if p.ImmunityTimer == nil {
		p.ImmunityTimer = time.AfterFunc(2*time.Second, func() {
			p.ImmunityTimer.Stop()
			p.ImmunityTimer = nil
			// Trigger event for immunity end
			game.BroadcastImmunityEnd(playerIndex)
		})
	}
}

// LoseLife processes the player losing a life
func (p *GamePlayer) LoseLife(game *Game, playerIndex int) {
	if p.Lives > 0 && p.ImmunityTimer == nil {
		p.Lives--
		if p.Lives == 0 {
			// Check if game is over
			game.CheckGameOver()
		}
		game.BroadcastPlayerDamage(playerIndex, p)
		p.StartImmunityTimer(game, playerIndex)
	}
}

func (p *GamePlayer) PlaceBomb(game *Game) {
	p.ActiveBombsPlaced++
	game.Map.gameMap[p.Position.Y][p.Position.X] = 7

	game.BroadCastBombPlacement(p.Position)
}
