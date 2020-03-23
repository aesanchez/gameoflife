package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Page is the main page
type Page struct {
	vecty.Core
}

// Render implements the vecty renderer interface.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(vecty.Class("Site")),
		p.renderHeader(),
		elem.Main(
			vecty.Markup(vecty.Class("Site-content")),
			elem.Section(
				vecty.Markup(
					vecty.Class("section"),
					prop.ID("formulario"),
				),
			),
		),
		p.renderFooter(),
	)
}

func (p *Page) renderHeader() *vecty.HTML {
	return elem.Section(
		vecty.Markup(
			vecty.Class("hero"),
			vecty.Class("is-info"),
		),
		elem.Div(
			vecty.Markup(vecty.Class("hero-body")),
			elem.Div(
				vecty.Markup(vecty.Class("container")),
				elem.Heading1(
					vecty.Markup(vecty.Class("title")),
					vecty.Text("Pinturito - Antonio Pan & Hijos"),
				),
			),
		),
	)
}

func (p *Page) renderFooter() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(vecty.Class("footer")),
		elem.Div(
			vecty.Markup(
				vecty.Class("content"),
				vecty.Class("has-text-centered"),
			),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://antoniopan.com.ar"),
				),
				vecty.Text("Antonio Pan & Hijos"),
			),
		),
	)
}
