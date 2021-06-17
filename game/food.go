package snakegame

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Variable insideborderWGlobal and insideborderHGlobal are variables consisting of the arena width and height
var insideborderWGlobal int
var insideborderHGlobal int

// NewFood will create a new piece of food, this will only happen once when the game has started.
func NewFood(insideborderW int ,insideborderH int ) *Food {
	food := new(Food)

	// To create food which don't appear on arena boundary
	insideborderWGlobal = insideborderW - 1
	insideborderHGlobal = insideborderH - 1
	// Create a new entity food with a standard position and 1x1 size
	food.Entity = tl.NewEntity(1, 1, 1, 1)
	// Call function MoveFood to move the food to a random position.
	food.MoveFood()

	return food
}

// MoveFood moves the food into a new random position.
func (food *Food) MoveFood() {

	// Calls the RandomInsideArena function to make sure that the foods spawns inside the arena.
	NewX := RandomInsideArena(insideborderWGlobal, 1)
	NewY := RandomInsideArena(insideborderHGlobal, 1)

	// Changes the X and Y coordinates of the food.
	food.Foodposition.X = NewX
	food.Foodposition.Y = NewY
	food.Emoji = '#'
	
	// Set the new position of the food.
	food.SetPosition(food.Foodposition.X, food.Foodposition.Y)
}


// Draw will print out the food on the screen.
func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Foodposition.X, food.Foodposition.Y, &tl.Cell{
		Ch: food.Emoji,
	})
}

// Contains checks if food contains the coordinates, if so this will return a bool.
func (food *Food) Contains(c Coordinates) bool {
	return c.X == food.Foodposition.X && c.Y == food.Foodposition.Y
}

// RandomInsideArena will the minimal, which is just inside the border and the maximal, being the arena width or height.
func RandomInsideArena(iMax int, iMin int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(iMax-iMin) + iMin
}
