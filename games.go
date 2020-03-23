package main

import "fmt"

type Matrix [][]int

type Game struct {
	Input  Matrix
	Output Matrix
}

func NewGame(w,h int) *Game {
	game := Game{}
	game.Input = make(Matrix, h)
	game.Output = make(Matrix, h)
	for i := 0; i < h; i++ {
		game.Output[i] = make([]int, w)
		game.Input[i] = make([]int, w)
	}
	return &game
}

func (g *Game) Swap() {
	aux := g.Output
	g.Output = g.Input
	g.Input = aux
}

func (g *Game) LoadLifeInput(i LifeInput) {
	for r, cells := range i.Cells {
		for c, cell := range cells {
			g.Input[i.RowOffset+r][i.ColumnOffset+c] = cell
		}
	}
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
