package main

import (
	"gameoflife/game"
	"strconv"
	"syscall/js"
)

var g game.Game

func main() {
	//register callback
	js.Global().Set("run", js.FuncOf(run))
	js.Global().Set("speed", js.FuncOf(speed))

	// so it doesn't die
	select {}
}

func run(this js.Value, i []js.Value) interface{} {
	width := js.Global().Get("document").Call("getElementById", "width").Get("value").String()
	height := js.Global().Get("document").Call("getElementById", "height").Get("value").String()
	input := js.Global().Get("document").Call("getElementById", "input").Get("value").String()
	w, _ := strconv.Atoi(width)
	h, _ := strconv.Atoi(height)
	var inputL game.LifeInput
	switch input {
	case "glider":
		inputL = game.Glider
	case "gliderGun":
		inputL = game.GliderGun
	case "spiral":
		inputL = game.Spiral
	default:
		inputL = game.Glider
	}
	g = *game.NewGame(w, h)
	g.LoadLifeInput(inputL, 5, 5)
	g.SetPeriodicBoundary(false)

	// set interval
	js.Global().Call("setMyInterval", js.FuncOf(update), 60)
	return nil
}
func speed(this js.Value, i []js.Value) interface{} {
	speed := js.Global().Get("document").Call("getElementById", "speed").Get("value").String()
	s, _ := strconv.Atoi(speed)

	// set interval
	js.Global().Call("setMyInterval", js.FuncOf(update), s)
	return nil
}

func update(this js.Value, i []js.Value) interface{} {
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	canvas.Set("innerText", g.Input.ToString())
	g.Tick()
	g.Swap()
	return nil
}
