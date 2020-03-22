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
	cellWidth float64 = 20
)

var (
	tickPeriod = 1200000
)

type Matrix [][]int

var lifeInput = Matrix{
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
	[]int{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

//var lifeInput = [][]bool{
//	[]bool{0, 0, 0, 0},
//	[]bool{0, 1, 1, 0},
//	[]bool{0, 1, 1, 0},
//	[]bool{0, 0, 0, 0},
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
	windowSize := cellWidth * float64(len(g.Input))
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, windowSize, windowSize),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Color = colornames.Black

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Clear()

		for r, cells := range g.Input {
			for c, cell := range cells {
				// ignore "Dead" cells
				if cell == 0 {
					continue
				}

				start := pixel.V(float64(c)*cellWidth, windowSize-float64(r)*cellWidth)
				imd.Push(start, pixel.V(start.X+cellWidth, start.Y+cellWidth))
				imd.Rectangle(0)
			}
		}

		imd.Draw(win)
		win.Update()

		g.Tick()
		g.Swap()

		<-time.Tick(time.Millisecond * time.Duration(tickPeriod))
	}
}

func main() {
	game := &Game{}
	game.Input = lifeInput
	game.Output = make(Matrix, len(lifeInput))
	for i := 0; i < len(lifeInput); i++ {
		game.Output[i] = make([]int, len(lifeInput[i]))
	}
	//game.Input.Print()
	//for {
	//	game.Tick()
	//	<-time.Tick(time.Millisecond * 500)
	//	fmt.Println()
	//	game.Output.Print()
	//	game.Swap()
	//}
	pixelgl.Run(func() { run(game) })
}

func countNeighbours(input Matrix, r, c int) int {
	alive := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i < 0 || j < 0) || (i >= len(input) || j >= len(input[r])) ||
				(i == r && j == c) {
				continue
			}

			if input[i][j] == 1 {
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
	if g.Input[r][c] == 1 {
		//alive
		if n < 2 || n > 3 {
			// * Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			// * Any live cell with more than three live neighbours dies, as if by overpopulation.
			g.Output[r][c] = 0
		} else {
			// * Any live cell with two or three live neighbours lives on to the next generation.
			g.Output[r][c] = 1
		}
	} else {
		//dead
		if n == 3 {
			g.Output[r][c] = 1
		}
	}
}

func (m Matrix) Print() {
	for _, cells := range m {
		for _, cell := range cells {
			if cell == 1 {
				fmt.Print("■")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
