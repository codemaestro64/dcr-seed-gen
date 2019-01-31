package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
)

type App struct {
	name          string
	hasRendered   bool
	currentPage   string
	renderhandler *RenderHandler
}

const (
	scaling = 1.3
)

func main() {
	app := &App{
		hasRendered:   false,
		name:          "DCR Seed Generator",
		currentPage:   "home",
		renderhandler: &RenderHandler{},
	}

	window := nucular.NewMasterWindow(0, app.name, app.render)
	window.SetStyle(style.FromTheme(style.DefaultTheme, scaling))
	window.Main()
}

func (app *App) render(window *nucular.Window) {
	if !app.hasRendered {
		app.hasRendered = true
		app.renderhandler.beforeRender(&app.currentPage)
	}

	if app.currentPage == "home" {
		app.renderhandler.renderHome(window)
	} else {
		app.renderhandler.renderVerify(window)
	}
}
