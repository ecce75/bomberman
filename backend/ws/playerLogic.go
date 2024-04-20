package ws

import (
	"time"
)

type Game struct {
	ID      int
	Players []*GamePlayer
}

// GamePlayer holds data about a game player
type GamePlayer struct {
	ID                int
	Username          string
	PlayerNumber      int
	Lives             int
	Powerups          Powerups
	Position          Coordinates
	ImmunityTimer     *time.Timer
	ActiveBombsPlaced int
}

// Powerups holds the power-up status of a player
type Powerups struct {
	MaxBombCount   int
	ExplosionRange int
	Speed          int
}

// NewGamePlayer creates a new player instance
func NewGamePlayer(id int, username string) *GamePlayer {
	return &GamePlayer{
		ID:                id,
		Username:          username,
		Lives:             3,
		Powerups:          Powerups{MaxBombCount: 1, ExplosionRange: 1, Speed: 1},
		Position:          Coordinates{X: 0, Y: 0},
		ActiveBombsPlaced: 0,
	}
}

// CanPlaceBomb checks if a player can place another bomb
func (p *GamePlayer) CanPlaceBomb() bool {
	return p.Powerups.MaxBombCount > p.ActiveBombsPlaced
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
			BroadcastImmunityEnd(game.ID, playerIndex)
		})
	}
}

// LoseLife processes the player losing a life
func (p *GamePlayer) LoseLife(game *Game, playerIndex int) {
	if p.Lives > 0 && p.ImmunityTimer == nil {
		p.Lives--
		if p.Lives == 0 {
			// Check if game is over
			CheckGameOver(game.ID)
		}
		BroadcastPlayerDamage(game.ID, playerIndex, p)
		p.StartImmunityTimer(game, playerIndex)
	}
}
