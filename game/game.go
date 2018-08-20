package game

//This file is defining the structs we will use. Structs are similar to objects.
//Capitalization of the struct name means it is public and can be used outside of it's particular package
// The fields must be capitalized otherwise there will be errors.
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
