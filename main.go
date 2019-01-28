package main

import (
	"github.com/aarzilli/nucular"
)

const (
	scaling = 1.3
)

func main() {
	handler := Newhandler()

	window := nucular.NewMasterWindow(0, "DCR SEED", handler.Render)
	setStyle(window)

	handler.SetMasterWindow(window)
	window.Main()
}
