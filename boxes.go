package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

const litBox = '█'

var cursorX, cursorY = 0, 0 // Track the cursor's position

var moveHistory [][]int

var maxDecay = 15

func main() {
	err := termbox.Init()
	// fullcolor :=
	// termbox.SetOutputMode(termbox.Output256)
	termbox.SetOutputMode(termbox.OutputRGB)

	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// set up initial grid
	drawGrid()

	// Main event loop, create an eventQueue to respond to
	eventQueue := make(chan termbox.Event)

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// respond to each event as ev
	for {
		ev := <-eventQueue
		if ev.Type == termbox.EventKey {
			// quit when an exit key is entered
			if ev.Key == termbox.KeyCtrlC || ev.Ch == 'q' || ev.Ch == 'Q' {
				break
			}
			// can handle other events here

			// handle cursor movement
			switch ev.Ch {
			case 'w', 'W': // up
				if cursorY > 0 {
					cursorY -= 2 // move up, adjust depending on grid vertical spcing
				}
			case 's', 'S': // down
				_, h := termbox.Size()
				if cursorY < h-2 {
					cursorY += 2 // move up, adjust depending on grid vertical spcing
				}
			case 'a', 'A': // left
				if cursorX > 0 {
					cursorX -= 2 // move up, adjust depending on grid vertical spcing
				}

			case 'd', 'D': // right
				_, w := termbox.Size()
				if cursorX < w*2 {
					cursorX += 2 // move up, adjust depending on grid vertical spcing
				}
			}
		}
		// Redraw the grid on every event
		drawGrid()
	}
}

/*
ideas:
- [] random movement of cursor
- [X] leave a trail of previously visited boxes that decays over time

issues:
- [] width is not being calculated properly
*/

func drawGrid() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	for y := 0; y < h; y += 2 {
		// rows
		for x := 0; x < w; x += 2 {
			// columns
			drawBox(x, y)
		}
	}

	// draw trail
	for i := 0; i < len(moveHistory); i++ {
		// init err
		var err error
		prevX, prevY, decay := moveHistory[i][0], moveHistory[i][1], moveHistory[i][2]
		// decay := moveHistory[i][2]
		// fmt.Print(prevX, " ", prevY, " ", decay)
		if decay > 0 {
			// box is still decaying, should be drawn as part of trail
			drawTrail(prevX, prevY, decay)
			moveHistory[i][2] -= 1
		} else {
			// remove element from moveHistory
			moveHistory, err = removeFromMoveHistory(moveHistory, i)
			if err != nil {
				panic(err)
			}
			// box is fully decayed, should be drawn white default
			drawBox(prevX, prevY)
		}
	}

	termbox.Flush()
}

func drawBox(x, y int) {
	// box drawing
	// modify depending on state of box
	if x == cursorX && y == cursorY {
		// this is the box where the cursor is
		red := termbox.RGBToAttribute(uint8(255), uint8(0), uint8(0))
		termbox.SetCell(x, y, litBox, red, termbox.ColorDefault)
		// add x,y to moveHistory
		moveHistory = append(moveHistory, []int{x, y, maxDecay})
	} else {
		// draw default box
		white := termbox.RGBToAttribute(uint8(255), uint8(255), uint8(255))

		termbox.SetCell(x, y, litBox, white, termbox.ColorDefault)
	}
}

func drawTrail(x, y, decay int) {

	r, g, b := calcRGB(decay) // get rgb by decay value
	termbox.SetCell(x, y, litBox, termbox.RGBToAttribute(r, g, b), termbox.ColorDefault)
	// switch decay {
	// case 3:
	// 	// draw box with different color
	// 	termbox.SetCell(x, y, litBox, termbox.ColorLightRed, termbox.ColorDefault)
	// case 2:
	// 	// draw box with different color
	// 	termbox.SetCell(x, y, litBox, termbox.ColorLightMagenta, termbox.ColorDefault)
	// case 1:
	// 	// draw box with different color
	// 	termbox.SetCell(x, y, litBox, termbox.ColorMagenta, termbox.ColorDefault)
	// case 0:
	// 	// reset box to default
	// 	termbox.SetCell(x, y, litBox, termbox.ColorWhite, termbox.ColorDefault)
	// }

}

func calcRGB(decay int) (r, g, b uint8) {
	decayFraction := float64(maxDecay-decay) / float64(maxDecay-1)
	r = uint8(255)
	g = uint8(decayFraction * 255.0)
	b = uint8(decayFraction * 255.0)
	return r, g, b
}

func removeFromMoveHistory(slice [][]int, i int) ([][]int, error) {

	if i >= len(slice) || i < 0 {
		return nil, fmt.Errorf("Index %d is out of bounds of slice length %d", i, len(slice))
	}
	return append(slice[:i], slice[i+1:]...), nil
}
