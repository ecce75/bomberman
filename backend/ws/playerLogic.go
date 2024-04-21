package ws

import (
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

// CanPlaceBomb checks if a player can place another bomb
func (p *GamePlayer) CanPlaceBomb() bool {
	return p.Powerups.Bomb > p.ActiveBombsPlaced
}

// IncreaseActiveBombs increases the count of active bombs placed by the player
func (p *GamePlayer) IncreaseActiveBombs() {
	if p.CanPlaceBomb() {
		p.ActiveBombsPlaced++
	}
}

// DecreaseActiveBombs decreases the count of active bombs placed by the player
func (p *GamePlayer) DecreaseActiveBombs() {
	if p.ActiveBombsPlaced > 0 {
		p.ActiveBombsPlaced--
	}
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
