package ws

import (
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[string]*Client)
var lobbies = make(map[string]*Lobby)
var games = make(map[string]*Game)

// Define wsMessage struct as per your application's requirements
type wsMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Client struct {
	ID     string
	Name   string
	Conn   *websocket.Conn
	GameID string
	Player *GamePlayer
}

type Game struct {
	ID           string
	Players      map[string]*Client
	Timer        *time.Timer
	Map          *gameMap
	ChatMessages []ChatMessage
}

type Lobby struct {
	ID       string
	Players  map[string]*Client
	TimeLeft int
}

type ChatMessage struct {
	Username string
	Message  string
	TimeSent string
}

// GamePlayer holds data about a game player
type GamePlayer struct {
	ID                string
	Username          string
	Lives             int
	Powerups          Powerups
	Position          Coordinates
	ImmunityTimer     *time.Timer
	ActiveBombsPlaced int
}

// Powerups holds the power-up status of a player
type Powerups struct {
	Bomb   int
	Flames int
	Speed  int
}

// gameMap represents the structure of the game map
type gameMap struct {
	mapWidth     int
	mapHeight    int
	gameMap      [][]int
	corners      [][2]int
	activeFlames []Coordinates
	activePowerups []PostFlameCoordinates
}

// Coordinates define a pair of x, y coordinates
type Coordinates struct {
	X int
	Y int
}

type PostFlameCoordinates struct {
	Position  Coordinates
	FieldCode int
}
