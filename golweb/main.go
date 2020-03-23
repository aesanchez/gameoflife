package main

import (
	"gameoflife/golweb/components"

	"github.com/gopherjs/vecty"
)

func main() {
	vecty.SetTitle("Game of Life")
	vecty.AddStylesheet("assets/bulma.min.css")
	vecty.AddStylesheet("assets/main.css")

	p := &components.Page{}

	vecty.RenderBody(p)
}
