package snakegame

import tl "github.com/JoelOtter/termloop"

//Game Object Variables.
var sg *tl.Game
var sp *Sidepanel
var gs *Gamescreen
var ts *Titlescreen

// Own created types.
type direction int

const (
	up direction = iota
	down
	left
	right
)

type Titlescreen struct {
	tl.Level
	Logo           *tl.Entity
	OptionsText    []*tl.Text
	Width int
	Height int
}

type Gameoverscreen struct {
	tl.Level
	Logo              *tl.Entity
	Finalstats        []*tl.Text
	OptionsBackground *tl.Rectangle
	OptionsText       []*tl.Text
}


type Gamescreen struct {
	tl.Level
	FPS             float64
	Score           int
	Length          int
	SnakeEntity     *Snake
	FoodEntity      *Food
	ArenaEntity     *Arena
	SidepanelObject *Sidepanel
	Background      *tl.Rectangle
}

type Sidepanel struct {
	Background     *tl.Rectangle
	Instructions   []string
	ScoreText      *tl.Text
	SpeedText      *tl.Text
	DimensionText     *tl.Text
	LengthText        *tl.Text
}

type Arena struct {
	*tl.Entity
	Width       int
	Height      int
	ArenaBorder map[Coordinates]int
}

type Snake struct {
	*tl.Entity
	Direction  direction
	Length     int
	Bodylength []Coordinates
	Speed      int
}

type Food struct {
	*tl.Entity
	Foodposition Coordinates
	Emoji        rune
}

type Coordinates struct {
	X int
	Y int
}
