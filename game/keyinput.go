package snakegame

import (

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)



// Tick listens for a keypress and then returns a direction for the snake.
func (snake *Snake) Tick(event tl.Event) {
	// Checks if the event is a keyevent.
	if event.Type == tl.EventKey {
		switch event.Key {
		// Checks if the key is a → press.
		case tl.KeyArrowRight:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != left {
				// Changes the direction of the snake to right.
				snake.Direction = right
			}
		// Checks if the key is a ← press.
		case tl.KeyArrowLeft:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != right {
				// Changes the direction of the snake to left.
				snake.Direction = left
			}
		// Checks if the key is a ↑ press.
		case tl.KeyArrowUp:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != down {
				// Changes the direction of the snake to down.
				snake.Direction = up
			}
		// Checks if the key is a ↓ press.
		case tl.KeyArrowDown:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != up {
				// Changes the direction of the snake to up.
				snake.Direction = down
			}
		}
	}
}

// Tick is a method for the gameoverscreen which listens for either a restart or a quit input from the user.
func (gos *Gameoverscreen) Tick(event tl.Event) {
	// Check if the event is a key event.
	if event.Type == tl.EventKey {
		switch event.Key {
		// If the key pressed is backspace the game will restart!!
		case tl.KeyHome:
			// Will call the RestartGame function to restart the game.
			RestartGame()
		case tl.KeyDelete:
			// Will end the game using a fatal log. This uses the termbox package as termloop does not have a function like that.
			tb.Close()
		}
	}
}

// Tick will listen for a keypress to initiate the game.
func (ts *Titlescreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event and the key pressed is the enter key.
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen(ts.Width,ts.Height)
			sg.Screen().SetLevel(gs)
		}
	}
}

