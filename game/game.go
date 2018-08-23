package game

import (
	"encoding/json"
	"log"
	"math/rand"
)

// Pos determines the position on the grid using X & Y axis
type Pos struct {
	X int
	Y int
}

// Size determines the size of the grid
type Size struct {
	Height int
	Width  int
}

// Player defines the struct with player information
type Player struct {
	Pos      Pos
	Name     string
	ID       string
	Cooldown int
	Money    uint
}

//State is the struct which represents the game state
type State struct {
	Players    Player `json:"players"`
	Start      int64  `json:"start"`
	Treasures  []Pos  `json:"treasures"`
	GridHeight int    `json:"gridHeight"`
	GridWidth  int    `json:"gridWidth"`
	You        string `json:"you"`
}

func newGameState() State {
	gs := State{
		Players:    newPlayer(),
		Start:      0,
		Treasures:  treasures(),
		GridHeight: 50,
		GridWidth:  50,
		You:        "String?",
	}
	return gs
}

func treasures() []Pos {
	return []Pos{
		{X: 5, Y: 5},
		{X: 15, Y: 15},
		{X: 20, Y: 20},
		{X: 45, Y: 45},
	}
}

func newPlayer() Player {
	player := &Player{
		Pos:      Pos{X: rand.Intn(50), Y: rand.Intn(50)},
		Name:     NameGenerator(),
		ID:       "ID?",
		Cooldown: rand.Intn(20),
		Money:    0,
	}
	return *player
}

// StateToJSON marshals the State struct to JSON
func (s *State) StateToJSON() []byte {
	json, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return json
}
