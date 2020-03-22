package main

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	cellWidth  = 20.0
	tickPeriod = 400
	N          = 50
	windowSize = cellWidth * float64(N)
)

func main() {
	game := NewGame(N)
	game.LoadLifeInput(gliderGun)

	pixelgl.Run(func() { run(game) })
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
				imd.Color = colornames.Black
				// imd.Color = colornames.Map[colornames.Names[rand.Intn(len(colornames.Names))]]
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
		default:
		}

	}
}
