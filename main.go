package main

import (
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

	cellsWidth  int = 50
	cellsHeight int = 30

	windowWidth  = cellWidth * float64(cellsWidth)
	windowHeight = cellWidth * float64(cellsHeight)
)

func main() {
	game := NewGame(cellsWidth, cellsHeight)
	game.LoadLifeInput(gliderGun)

	pixelgl.Run(func() { run(game) })
}

func run(g *Game) {
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

		// cells
		for r, cells := range g.Input {
			for c, cell := range cells {
				// ignore "Dead" cells
				if cell == 0 {
					continue
				}
				// imd.Color = colornames.Black
				// imd.Color = colornames.Map[colornames.Names[rand.Intn(len(colornames.Names))]]
				start := pixel.V(float64(c)*cellWidth, windowHeight-float64(r)*cellWidth)
				imd.Push(start, pixel.V(start.X+cellWidth, start.Y+cellWidth))
				imd.Rectangle(0)
			}
		}

		// grid
		prevColor := imd.Color
		imd.Color = color.RGBA{R: 125, G: 125, B: 125, A: 0x8F}
		index := 0.0
		for index < windowHeight {
			imd.Push(pixel.V(0.0, index), pixel.V(windowWidth, index))
			imd.Line(1.1)
			index += cellWidth
		}
		index = 0.0
		for index < windowWidth {
			imd.Push(pixel.V(index, 0.0), pixel.V(index, windowHeight))
			imd.Line(1.1)
			index += cellWidth
		}
		imd.Color = prevColor

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
			imd.Color = newColor()

		default:
		}

	}
}

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
