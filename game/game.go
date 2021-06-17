package snakegame

import (
	"fmt"
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

var widthGlobal , heightGlobal int

// StartGame will start the game with the tilescreen.
func StartGame(width int, height int) {
	sg = tl.NewGame()
	widthGlobal = width
	heightGlobal = height
	// Create titlescreen.
	ts := NewTitleScreen(width,height)

	ts.AddEntity(ts.Logo)

	// Range options and add entities.
	for _, v := range ts.OptionsText {
		ts.AddEntity(v)
	}

	// Set FPS and start game.
	sg.Screen().SetFps(10)
	sg.Screen().SetLevel(ts)
	sg.Start()
}

// NewTitleScreen will create a new titlescreen and return it.
func NewTitleScreen( width int, height int) *Titlescreen {
	// Create a title screen and its objects.
	ts = new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
	})

	logofile, _ := ioutil.ReadFile("util/titlescreen-logo.txt")
	ts.Logo = tl.NewEntityFromCanvas(38, 2, tl.CanvasFromString(string(logofile)))

	ts.Width = width
	ts.Height = height
	ts.OptionsText = []*tl.Text{
		tl.NewText(50, 20, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack),
	}

	return ts
}


func NewGamescreen(width int, height int) *Gamescreen {
	// Creates the gamescreen level and create the entities
	gs = new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gs.Background = tl.NewRectangle(1, 1, width-1, height-1, tl.ColorGreen)
	gs.FPS = 10
	gs.Score = 0
	gs.Length = 3
	gs.SnakeEntity = NewSnake()
	gs.ArenaEntity = NewArena(width, height)
	gs.FoodEntity = NewFood(width, height)
	gs.SidepanelObject = NewSidepanel(width, height)

	// Add entities for the game level.
	gs.AddEntity(gs.Background)
	gs.AddEntity(gs.FoodEntity)
	gs.AddEntity(gs.SidepanelObject.Background)
	gs.AddEntity(gs.SidepanelObject.ScoreText)
	gs.AddEntity(gs.SidepanelObject.DimensionText)
	gs.AddEntity(gs.SidepanelObject.LengthText)
	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.ArenaEntity)

	

	// Range over the instructions and add them to the entities
	y := 7
	for _, v := range sp.Instructions {
		var i *tl.Text
		y += 2
		i = tl.NewText(width+2, y, v, tl.ColorBlack, tl.ColorWhite)
		gs.AddEntity(i)
	}

	// Set Fps and return the screen
	sg.Screen().SetFps(gs.FPS)

	return gs
}

// NewSidepanel will create a new sidepanel given the arena height and width.
func NewSidepanel(width int,height int) *Sidepanel {
	// Create a sidepanel and its objects and return it
	sp = new(Sidepanel)
	sp.Instructions = []string{
		"Instructions:",
		"Use ← → ↑ ↓ to move the snake around",
		"Pick up the food to grow bigger",
		"#: 1 point/growth",
	}

	sp.Background = tl.NewRectangle(width+1, 0,40, height, tl.ColorWhite)
	sp.DimensionText = tl.NewText(width+2, 1, fmt.Sprintf("Board Dimension (W X H): %d X %d", widthGlobal,heightGlobal), tl.ColorBlack, tl.ColorWhite)
	sp.ScoreText = tl.NewText(width+2, 3, fmt.Sprintf("Score: %d", gs.Score), tl.ColorBlack, tl.ColorWhite)
	sp.LengthText = tl.NewText(width+2, 5, fmt.Sprintf("Length of Snake: %d", len(gs.SnakeEntity.Bodylength)), tl.ColorBlack, tl.ColorWhite)

	return sp
}

func Gameover() {
	// Create a new gameover screen and its content.
	gos := new(Gameoverscreen)
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	logofile, _ := ioutil.ReadFile("util/gameover-logo.txt")
	gos.Logo = tl.NewEntityFromCanvas(40, 2, tl.CanvasFromString(string(logofile)))

	gos.Finalstats = []*tl.Text{
		tl.NewText(50, 15, fmt.Sprintf("Score: %d", gs.Score), tl.ColorWhite, tl.ColorBlack),
		tl.NewText(50, 17, fmt.Sprintf("Length of Snake: %d", len(gs.SnakeEntity.Bodylength)), tl.ColorWhite, tl.ColorBlack),

	}
	gos.OptionsBackground = tl.NewRectangle(45, 20, 45, 7, tl.ColorWhite)
	gos.OptionsText = []*tl.Text{
		tl.NewText(50, 22, "Press \"Home\" key to restart!", tl.ColorBlack, tl.ColorWhite),
		tl.NewText(50, 24, "Press \"Delete\" key to quit!", tl.ColorBlack, tl.ColorWhite),
	}

	// Add all of the entities to the screen
	for _, v := range gos.Finalstats {
		gos.AddEntity(v)
	}
	gos.AddEntity(gos.Logo)
	gos.AddEntity(gos.OptionsBackground)

	for _, vv := range gos.OptionsText {
		gos.AddEntity(vv)
	}
	// Set the screen
	sg.Screen().SetLevel(gos)
}

// UpdateScore updates the score with the given amount of points.
func UpdateScore(amount int) {
	gs.Score += amount
	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
	sp.LengthText.SetText(fmt.Sprintf("Length of Snake: %d", len(gs.SnakeEntity.Bodylength)))
}


// RestartGame will restart the game and reset the position of the food and the snake to prevent collision issues.
func RestartGame() {
	// Removes the current snake and food from the level.
	gs.RemoveEntity(gs.SnakeEntity)
	gs.RemoveEntity(gs.FoodEntity)

	// Generate a new snake and food.
	gs.SnakeEntity = NewSnake()
	gs.FoodEntity = NewFood(widthGlobal,heightGlobal)	

	// Revert the score and fps to the standard.
	gs.FPS = 10
	gs.Score = 0

	// Update the score and fps text.
	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
	sp.LengthText.SetText(fmt.Sprintf("Length of Snake: %d", len(gs.SnakeEntity.Bodylength)))

	// Adds the snake and food back and sets them to the standard position.
	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.FoodEntity)
	sg.Screen().SetFps(gs.FPS)
	sg.Screen().SetLevel(gs)
}

