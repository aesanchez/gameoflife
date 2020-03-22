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

var (
	tickPeriod = 120
	cellWidth  = 20
)

type Matrix [][]bool

var lifeInput = Matrix{
	[]bool{false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false},
	[]bool{false, false, true, true, true, false},
	[]bool{false, true, true, true, false, false},
	[]bool{false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false},
}

//var lifeInput = [][]bool{
//	[]bool{false, false, false, false},
//	[]bool{false, true, true, false},
//	[]bool{false, true, true, false},
//	[]bool{false, false, false, false},
//}

type Game struct {
	Input  Matrix
	Output Matrix
}

func (g *Game) Swap() {
	aux := g.Output
	g.Output = g.Input
	g.Input = aux
}

func run(g *Game) {

	windowSize := cellWidth * len(g.Input)
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

	for r, cells := range g.Input {
		for c := range cells {
			g.RunRules(r, c)
		}
	}

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Clear()

		g.Tick()
		g.Swap()

		imd.Push(point1, pixel.V(point1.X+10, point1.Y+10))
		imd.Rectangle(0)

		imd.Draw(win)
		win.Update()
		point1.X++
		point1.Y++
		<-time.Tick(tickPeriod * time.Millisecond)
	}
}

func main() {
	game := Game{}
	game.Input = lifeInput
	game.Output = make(Matrix, len(lifeInput))
	for i := 0; i < len(lifeInput); i++ {
		game.Output[i] = make([]bool, len(lifeInput[i]))
	}
	game.Input.Print()
	for {
		game.Tick()
		<-time.Tick(time.Millisecond * 500)
		fmt.Println()
		game.Output.Print()
		game.Swap()
	}
	// pixelgl.Run(run(game))
}

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

func (g *Game) Tick() {
	for r, cells := range g.Input {
		for c := range cells {
			g.RunRules(r, c)
		}
	}
}

func (g *Game) RunRules(r, c int) {
	n := countNeighbours(g.Input, r, c)
	// fmt.Printf("Row = %d, Column = %d, N = %d\n", r, c, n)
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
