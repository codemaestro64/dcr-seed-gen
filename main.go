package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
)

const (
	scaling = 1.9
)

func main() {
	handler := Newhandler()

	window := nucular.NewMasterWindow(0, "DCR SEED", handler.Render)
	window.SetStyle(style.FromTheme(style.DefaultTheme, scaling))

	handler.SetMasterWindow(window)

	window.Main()
}
