package main

// The following are the rules for Conway's game of life
// * Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// * Any live cell with two or three live neighbours lives on to the next generation.
// * Any live cell with more than three live neighbours dies, as if by overpopulation.
// * Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	tickRate   = 120
	windowSize = 1000
)

type Matrix [][]bool

var lifeInput = [][]bool{
	[]bool{false, false, false, false, false},
	[]bool{false, false, false, false, false},
	[]bool{false, true, true, true, false},
	[]bool{false, false, false, false, false},
	[]bool{false, false, false, false, false},
}

type Game struct {
	Input  Matrix
	Output Matrix
}

func (g *Game) Swap() {
	g.Input = g.Output
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, windowSize, windowSize),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Color = colornames.Black

	point1 := pixel.Vec{X: 0, Y: 0}

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Clear()

		imd.Push(point1, pixel.V(point1.X+10, point1.Y+10))
		imd.Rectangle(0)

		imd.Draw(win)
		win.Update()
		point1.X++
		point1.Y++
		<-time.Tick(1000 / tickRate * time.Millisecond)
	}
}

func main() {
	game := Game{}
	game.Input = lifeInput
	game.Output = lifeInput
	game.Input.Print()
	for {
		for r, cells := range game.Input {
			for c := range cells {
				game.RunRules(r, c)
			}
		}
	}
	// pixelgl.Run(run)
}

/*
x x x x
o o x o
x x o o
x x x o
*/
func countNeighbours(input Matrix, r, c int) int {
	alive := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i < 0 || j < 0) || (i >= len(input) || j >= len(input[r])) ||
				(i == r && j == c) {
				continue
			}

			if input[i][j] {
				alive++
			}
		}
	}

	return alive
}

func (g *Game) RunRules(r, c int) {
	n := countNeighbours(g.Input, r, c)
	fmt.Printf("Row = %d, Column = %d, N = %d\n", r, c, n)
	if g.Input[r][c] {
		//alive
		if n < 2 || n > 3 {
			// * Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			// * Any live cell with more than three live neighbours dies, as if by overpopulation.
			g.Output[r][c] = false
		} else {
			// * Any live cell with two or three live neighbours lives on to the next generation.
			g.Output[r][c] = true
		}
	} else {
		//dead
		if n == 3 {
			g.Output[r][c] = true
		}
	}
}

func (m Matrix) Print() {
	for _, cells := range m {
		for _, cell := range cells {
			if cell == true {
				fmt.Print("■")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
