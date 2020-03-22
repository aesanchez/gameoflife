package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	tickRate = 120
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1000, 1000),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Color = colornames.Black

	point1 := pixel.Vec{X: 0, Y: 0}
	point2 := pixel.Vec{X: 50, Y: 50}

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Clear()

		imd.Push(point1, pixel.V(point1.X+10, point1.Y+10))
		imd.Rectangle(0)

		imd.Push(point2, pixel.V(point2.X+10, point2.Y+10))
		imd.Rectangle(0)

		imd.Draw(win)
		win.Update()
		point1.X++
		point1.Y++
		<-time.Tick(1000 / tickRate * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}
