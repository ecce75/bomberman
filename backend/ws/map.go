package ws

import (
	"math/rand"
	"strconv"
)

// Documentation for map field codes: {
//     0: "free",
//     1: "indestructible",
//     2: "destructible",
//     3: "player1",
//     4: "player2",
//     5: "player3",
//     6: "player4",
//     7: "bomb",
//     8: "booked" // for development purposes
//     9: "powerup: speed"
//     10: "powerup: explosion length"
//     11: "powerup: bombCount"
//     9: "flame"
// }

// NewGameMap initializes a new game map with default settings
func NewGameMap(players map[string]*Client) *gameMap {
	gm := &gameMap{
		mapWidth:  13,
		mapHeight: 13,
		gameMap:   make([][]int, 13), // Initializes the rows of the map
		corners: [][2]int{
			{0, 0}, {0, 1}, {1, 0},
			{12, 0}, {11, 0}, {12, 1},
			{0, 12}, {0, 11}, {1, 12},
			{12, 12}, {12, 11}, {11, 12},
		},
		activeFlames: []Coordinates{},
	}

	// Initializes the columns of the map
	for i := range gm.gameMap {
		gm.gameMap[i] = make([]int, gm.mapWidth)
	}

	gm.initMap()
	// gm.bookCorners()
	gm.placeDestructibleBlocks(100)
	gm.placePlayers(initiatePlayers(players))

	return gm
}

// initMap initializes the map with indestructible blocks on the edges and in a grid pattern
func (gm *gameMap) initMap() {
	for y := 0; y < gm.mapHeight; y++ {
		for x := 0; x < gm.mapWidth; x++ {
			// Set open field on the edges
			if x == 0 || x == gm.mapWidth-1 || y == 0 || y == gm.mapHeight-1 {
				gm.setFieldID(x, y, 0)
			} else if x%2 == 1 && y%2 == 1 { // Set indestructible blocks in a grid pattern inside
				gm.setFieldID(x, y, 1)
			} else {
				gm.setFieldID(x, y, 0)
			}
		}
	}
}

func initiatePlayers(players map[string]*Client) map[string]*Client {
	counter := 1
	for _, player := range players {
		newPlayer := NewGamePlayer(strconv.Itoa(counter), player.Name)
		player.Player = newPlayer
		counter++
	}
	return players
}

// isInCorners checks if the given coordinates are in the list of corner coordinates.
func (gm *gameMap) isInCorners(x, y int) bool {
	for _, corner := range gm.corners {
		if corner[0] == x && corner[1] == y {
			return true
		}
	}
	return false
}

// placeIndestructibleBlocks randomly places a specified number of indestructible blocks on the map
func (gm *gameMap) placeDestructibleBlocks(count int) {
	placedBlocks := 0
	for placedBlocks < count {
		x := rand.Intn(gm.mapWidth)
		y := rand.Intn(gm.mapHeight)
		if gm.gameMap[y][x] != 1 && !gm.isInCorners(x, y) {
			gm.setFieldID(x, y, 2)
			placedBlocks++
		}
	}
}

// placePlayers sets player positions on the map
func (gm *gameMap) placePlayers(players map[string]*Client) {
	cornerCoordinates := [][2]int{
		{0, 0}, {0, 12}, {12, 0}, {12, 12},
	}
	player := 3
	for _, playerObj := range players {
		x, y := cornerCoordinates[player-2][0], cornerCoordinates[player-2][1]
		gm.setFieldID(x, y, player)
		playerObj.Player.SetPosition(x, y)
		player++
	}
}

// setFieldID sets a new ID at the specified field coordinates
func (gm *gameMap) setFieldID(x, y, newID int) {
	gm.gameMap[y][x] = newID
}

// isActiveFlameOnCell checks if there is an active flame on the cell
func (gm *gameMap) isActiveFlameOnCell(coordinates Coordinates) bool {
	for _, flame := range gm.activeFlames {
		if flame.X == coordinates.X && flame.Y == coordinates.Y {
			return true
		}
	}
	return false
}

// addActiveFlames adds new active flames to the map
func (gm *gameMap) addActiveFlames(newFlameCoordinates []Coordinates) {
	gm.activeFlames = append(gm.activeFlames, newFlameCoordinates...)
}

// removeActiveFlames removes an active flame from the map
func (gm *gameMap) removeActiveFlames(flameToRemove Coordinates) bool {
	for i, flame := range gm.activeFlames {
		if flame.X == flameToRemove.X && flame.Y == flameToRemove.Y {
			gm.activeFlames = append(gm.activeFlames[:i], gm.activeFlames[i+1:]...)
			return true
		}
	}
	return false
}

func (p *GamePlayer) SetPosition(x int, y int) {
	p.Position.X = x
	p.Position.Y = y
}
