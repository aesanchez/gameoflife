package main

import (
	"gameoflife/game"
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//change this values as needed
const (
	// board size. It must be bigger than the input
	cellsWidth  int = 80
	cellsHeight int = 50

	// max window size
	windowMaxWidth  float64 = 1600
	windowMaxHeight float64 = 1000

	// multicolor cells
	strobeMode = false

	// cyclic boundaries
	periodicBoundary = true
)

var (
	tickPeriod   = 60
	cellSize     float64
	windowWidth  = windowMaxWidth
	windowHeight = windowMaxHeight
)

func main() {
	// this should be called always to fit the game to the window max size
	calculateResolution()

	// input can be loaded from files
	// input := game.ReadInputFile("res/cambrian-explosion.rle")

	// input can be also loaded from pre-loaded shapes in life_input.go
	input := game.GliderGun

	g := game.NewGame(cellsWidth, cellsHeight)
	g.LoadLifeInput(input, 10, 10)
	g.SetPeriodicBoundary(periodicBoundary)

	pixelgl.Run(func() { run(g) })
}

func calculateResolution() {
	cellWidth := windowMaxWidth / float64(cellsWidth)
	cellHeight := windowMaxHeight / float64(cellsHeight)
	if cellWidth < cellHeight {
		cellSize = cellWidth
	} else {
		cellSize = cellHeight
	}
	windowWidth = cellSize * float64(cellsWidth)
	windowHeight = cellSize * float64(cellsHeight)
}

func run(g *game.Game) {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
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

		printGrid(imd)
		printCells(imd, g)

		imd.Draw(win)
		win.Update()

		if win.Pressed(pixelgl.KeyUp) && tickPeriod > 10 {
			tickPeriod -= 10
		}
		if win.Pressed(pixelgl.KeyDown) {
			tickPeriod += 10
		}

		select {
		case <-tickChan:
			g.Tick()
			g.Swap()
			tickChan = time.Tick(time.Millisecond * time.Duration(tickPeriod))
			if strobeMode {
				imd.Color = newColor()
			} else {
				imd.Color = colornames.Black
			}
		default:
		}

	}
}

// returns a non light color
func newColor() color.Color {
	for {
		c := colornames.Map[colornames.Names[rand.Intn(len(colornames.Names))]]
		c.A = 0xFF
		aux := int32(c.R) + int32(c.G) + int32(c.B)
		if aux < 500 {
			return c
		}
	}
}

func printGrid(imd *imdraw.IMDraw) {
	prevColor := imd.Color
	imd.Color = color.RGBA{R: 125, G: 125, B: 125, A: 0x8F}
	index := 0.0
	for index < windowHeight {
		imd.Push(pixel.V(0.0, index), pixel.V(windowWidth, index))
		imd.Line(1.1)
		index += cellSize
	}
	index = 0.0
	for index < windowWidth {
		imd.Push(pixel.V(index, 0.0), pixel.V(index, windowHeight))
		imd.Line(1.1)
		index += cellSize
	}
	imd.Color = prevColor
}

func printCells(imd *imdraw.IMDraw, g *game.Game) {
	for r, cells := range g.Input {
		for c, cell := range cells {
			// ignore "Dead" cells
			if cell == 0 {
				continue
			}
			start := pixel.V(float64(c)*cellSize, windowHeight-float64(r)*cellSize)
			imd.Push(start, pixel.V(start.X+cellSize, start.Y-cellSize))
			imd.Rectangle(0)
		}
	}
}
