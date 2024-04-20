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

// gameMap represents the structure of the game map
type gameMap struct {
	mapWidth     int
	mapHeight    int
	gameMap      [][]int
	corners      [][2]int
	activeFlames []Coordinates
}

// Coordinates define a pair of x, y coordinates
type Coordinates struct {
	X int
	Y int
}

func (p GamePlayer) SetPosition(x int, y int) {
	p.Position.X = x
	p.Position.Y = y
}

func initiatePlayers(players map[string]*Client) []*GamePlayer {
	gamePlayers := make([]*GamePlayer, 0)
	for _, player := range players {
		playerIDint, err := strconv.Atoi(player.ID)
		if err != nil {
			// handle error appropriately
		}
		gamePlayers = append(gamePlayers, NewGamePlayer(playerIDint, player.Name))
	}
	return gamePlayers
}

// NewGameMap initializes a new game map with default settings
func NewGameMap(players map[string]*Client) *gameMap {
	gm := &gameMap{
		mapWidth:  13,
		mapHeight: 13,
		gameMap:   make([][]int, 13),
		corners: [][2]int{
			{1, 1}, {1, 2}, {2, 1},
			{13, 1}, {12, 1}, {13, 2},
			{1, 11}, {1, 10}, {2, 11},
			{13, 11}, {12, 11}, {13, 10},
		},
		activeFlames: []Coordinates{},
	}

	for i := range gm.gameMap {
		gm.gameMap[i] = make([]int, gm.mapWidth)
	}

	gm.initMap()
	// gm.bookCorners()
	gm.placeDestructibleBlocks(50)
	gm.placePlayers(initiatePlayers(players))

	return gm
}

// initMap initializes the map with indestructible blocks on the edges and in a grid pattern
func (gm *gameMap) initMap() {
	for y := 0; y < gm.mapHeight; y++ {
		for x := 0; x < gm.mapWidth; x++ {
			// Set indestructible blocks on the edges
			if x == 0 || x == gm.mapWidth-1 || y == 0 || y == gm.mapHeight-1 {
				gm.gameMap[y][x] = 0
			} else if x%2 == 1 && y%2 == 1 { // Set indestructible blocks in a grid pattern inside
				gm.gameMap[y][x] = 1
			} else {
				gm.gameMap[y][x] = 0 // Set the rest as free blocks
			}
		}
	}
}

// bookCorners marks the corners of the map as booked
func (gm *gameMap) bookCorners() {
	for _, corner := range gm.corners {
		gm.gameMap[corner[1]][corner[0]] = 8
	}
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
			gm.gameMap[y][x] = 2
			placedBlocks++
		}
	}
}

// freeCorners clears the booked corners for use by players
func (gm *gameMap) freeCorners() {
	for _, corner := range gm.corners {
		gm.gameMap[corner[1]][corner[0]] = 0
	}
}

// placePlayers sets player positions on the map
func (gm *gameMap) placePlayers(players []*GamePlayer) {
	cornerCoordinates := [][2]int{
		{0, 0}, {0, 12}, {12, 0}, {12, 12},
	}
	player := 3
	for i, playerObj := range players {
		x, y := cornerCoordinates[i][0], cornerCoordinates[i][1]
		gm.gameMap[y][x] = player
		playerObj.SetPosition(x, y)
		player++
	}
}

// getFieldID returns the field ID at specified coordinates
func (gm *gameMap) getFieldID(x, y int) int {
	return gm.gameMap[y][x]
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
