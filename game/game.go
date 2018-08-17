package game

type Pos struct {
	x int
	y int
}

type Size struct {
	height int
	width  int
}
type Player struct {
	pos      Pos
	name     string
	id       string
	cooldown int64
	money    uint
}

type GameState struct {
	players    map[string]Player
	start      int64
	treasures  []Pos
	gridHeight int
	gridWidth  int
	you        string
}
