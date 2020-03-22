package main

// The following are the rules for Conway's game of life
// * Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// * Any live cell with two or three live neighbours lives on to the next generation.
// * Any live cell with more than three live neighbours dies, as if by overpopulation.
// * Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	cellWidth  = 20.0
	tickPeriod = 400
	N          = 30
	windowSize = cellWidth * float64(N)
)

type Matrix [][]int

type LifeInput struct {
	Cells        Matrix
	RowOffset    int
	ColumnOffset int
}

var lifeInput1 = LifeInput{
	Cells: Matrix{
		[]int{1, 0, 0},
		[]int{0, 1, 1},
		[]int{1, 1, 0},
	},
	RowOffset:    5,
	ColumnOffset: 5,
}

var lifeInput2 = LifeInput{
	Cells: Matrix{
		[]int{1, 0, 0},
		[]int{0, 1, 1},
		[]int{1, 1, 0},
	},
	RowOffset:    15,
	ColumnOffset: 5,
}

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

	tickChan := time.Tick(time.Millisecond * time.Duration(tickPeriod))
	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Clear()

		for r, cells := range g.Input {
			for c, cell := range cells {
				// ignore "Dead" cells
				if cell == 0 {
					continue
				}
				imd.Color = colornames.Map[colornames.Names[rand.Intn(len(colornames.Names))]]
				start := pixel.V(float64(c)*cellWidth, windowSize-float64(r)*cellWidth)
				imd.Push(start, pixel.V(start.X+cellWidth, start.Y+cellWidth))
				imd.Rectangle(0)
			}
		}

		var index = 0.0
		imd.Color = color.RGBA{R: 125, G: 125, B: 125, A: 0x8F}
		for index < windowSize {
			imd.Push(pixel.V(0.0, index), pixel.V(windowSize, index))
			imd.Line(1.1)
			imd.Push(pixel.V(index, 0.0), pixel.V(index, windowSize))
			imd.Line(1)
			index += cellWidth
		}

		imd.Draw(win)
		win.Update()

		if win.Pressed(pixelgl.KeyUp) && tickPeriod > 20 {
			tickPeriod -= 20
		}
		if win.Pressed(pixelgl.KeyDown) {
			tickPeriod += 20
		}

		select {
		case <-tickChan:
			g.Tick()
			g.Swap()
			tickChan = time.Tick(time.Millisecond * time.Duration(tickPeriod))
		default:
		}

	}
}

func main() {
	game := &Game{}
	game.Input = make(Matrix, N)
	game.Output = make(Matrix, N)
	for i := 0; i < N; i++ {
		game.Output[i] = make([]int, N)
		game.Input[i] = make([]int, N)
	}

	for r, cells := range lifeInput1.Cells {
		for c, cell := range cells {
			game.Input[lifeInput1.RowOffset+r][lifeInput1.ColumnOffset+c] = cell
		}
	}

	for r, cells := range lifeInput2.Cells {
		for c, cell := range cells {
			game.Input[lifeInput2.RowOffset+r][lifeInput2.ColumnOffset+c] = cell
		}
	}

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
		} else {
			g.Output[r][c] = 0
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
