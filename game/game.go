package game

type Pos struct {
	X int
	Y int
}

type Size struct {
	Height int
	Width  int
}
type Player struct {
	Pos      Pos
	Name     string
	Id       string
	Cooldown int64
	Money    uint
}

type GameState struct {
	Players    map[string]Player
	Start      int64
	Treasures  []Pos
	GridHeight int
	GridWidth  int
	You        string
}
