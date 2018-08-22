package game

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
	Players    map[string]Player
	Start      int64
	Treasures  []Pos
	GridHeight int
	GridWidth  int
	You        string
}

// newGameState is the constructor for the GameState struct
func newGameState() gameState {
	gs := GameState{
		Players:
		Start: 
		Treasures:
		GridHeight:
		GridWidth:
		You:
	}
	return gs
}