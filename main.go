package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

type pageHandler struct {
	beforeRender func()
	render       func(*nucular.Window)
}

type Handler struct {
	currentPage string
	pages       map[string]pageHandler

	errorPopupBounds rect.Rect
	hasPageChanged   bool

	masterWindow nucular.MasterWindow
}

const (
	scaling = 1.3
)

func NewHandler() *Handler {
	handler := &Handler{
		currentPage:    "form",
		hasPageChanged: false,
		pages:          make(map[string]pageHandler, 2), // currently has only two pages
	}

	handler.registerPageHandlers()

	handler.errorPopupBounds = rect.Rect{
		X: 400,
		Y: 180,
		W: 500,
		H: 180,
	}

	return handler
}

func (h *Handler) registerPageHandlers() {
	h.pages["form"] = pageHandler{
		beforeRender: h.formBeforeRender,
		render:       h.render,
	}
}

func (h *Handler) setMasterWindow(window nucular.MasterWindow) {
	h.masterWindow = window
}

func (h *Handler) changePage(page string) {
	h.currentPage = page
	h.hasPageChanged = true
	h.masterWindow.Changed()
}

func (h *Handler) render(window *nucular.Window) {
	page := h.pages[h.currentPage]
	// ensure that the handler's BeforeRender function is called only once per page call
	if h.hasPageChanged {
		page.beforeRender()
		h.hasPageChanged = false
	}

	page.render(window)
}

func main() {
	handler := NewHandler()

	window := nucular.NewMasterWindow(0, "DCR SEED", handler.Render)
	setStyle(window)

	handler.SetMasterWindow(window)
	window.Main()
}
