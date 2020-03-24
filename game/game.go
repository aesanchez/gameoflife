package game

import (
	"fmt"
	"sync"
)

type Matrix [][]int

type Game struct {
	Input            Matrix
	Output           Matrix
	Width            int
	Height           int
	PeriodicBoundary bool
}

func InitMatrix(w, h int) Matrix {
	m := make(Matrix, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
	}
	return m
}

func NewGame(w, h int) *Game {
	game := Game{Width: w, Height: h}
	game.PeriodicBoundary = false
	game.Input = InitMatrix(w, h)
	game.Output = InitMatrix(w, h)
	return &game
}

func (g *Game) SetPeriodicBoundary(b bool) {
	g.PeriodicBoundary = b
}

func (g *Game) Swap() {
	aux := g.Output
	g.Output = g.Input
	g.Input = aux
}

func (g *Game) LoadLifeInput(i LifeInput, rowOffset, columnOffset int) {
	if i.Width+columnOffset > g.Width || i.Height+rowOffset > g.Height {
		panic("Input out of bounds")
	}
	for r, cells := range i.Cells {
		for c, cell := range cells {
			g.Input[rowOffset+r][columnOffset+c] = cell
		}
	}
}

func (g *Game) countNeighbours(r, c int) int {
	alive := 0
	for row := r - 1; row <= r+1; row++ {
		for col := c - 1; col <= c+1; col++ {
			if !g.PeriodicBoundary {
				if (row < 0 || col < 0) || (row >= g.Height || col >= g.Width) ||
					(row == r && col == c) {
					continue
				}
				if g.Input[row][col] == 1 {
					alive++
				}
			} else {
				if !(row == r && col == c) && g.Input[(row+g.Height)%g.Height][(col+g.Width)%g.Width] == 1 {
					alive++
				}
			}

		}
	}
	return alive
}

// ToDo: find the most efficient way to do this
func (g *Game) Tick() {
	var wg sync.WaitGroup
	for r, cells := range g.Input {
		wg.Add(1)
		go func(r int, cells []int, wg *sync.WaitGroup) {
			defer wg.Done()
			for c := range cells {
				g.RunRules(r, c)
			}
		}(r, cells, &wg)
	}
	wg.Wait()
}

func (g *Game) RunRules(r, c int) {
	n := g.countNeighbours(r, c)
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
