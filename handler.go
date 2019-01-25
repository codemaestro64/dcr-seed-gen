package main

import (
	"encoding/hex"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/rect"
	"github.com/tyler-smith/go-bip39"
)

type Handler struct {
	errorPopupBounds  rect.Rect
	mnemonicSeedInput nucular.TextEditor
	passInput         nucular.TextEditor
	dcrSeed           string
	masterWindow      nucular.MasterWindow
}

func Newhandler() *Handler {
	h := &Handler{}
	h.mnemonicSeedInput.Flags = nucular.EditMultiline | nucular.EditClipboard
	h.passInput.Flags = nucular.EditField
	h.passInput.PasswordChar = '*'

	h.errorPopupBounds = rect.Rect{
		X: 400,
		Y: 180,
		W: 500,
		H: 180,
	}

	return h
}

func (h *Handler) SetMasterWindow(masterWindow nucular.MasterWindow) {
	h.masterWindow = masterWindow
}

func (h *Handler) Render(window *nucular.Window) {
	window.Row(30).Ratio(0.4, 0.6)
	window.Label("BIP39 Mnenomic: ", "LC")
	h.mnemonicSeedInput.Edit(window)

	window.Row(30).Ratio(0.4, 0.6)
	window.Label("BIP39 Passphrase (Optional): ", "LC")
	h.passInput.Edit(window)

	window.Row(30).Ratio(0.4, 0.6)
	window.Label("", "LC")
	if window.Button(label.T("Generate"), false) {
		h.generate()
	}

	if h.dcrSeed != "" {
		window.Row(20).Dynamic(1)
		window.Label("", "LC")

		window.Row(30).Ratio(0.4, 0.6)
		window.Label("BIP39 Seed:", "LC")

		window.Row(100).Dynamic(2)
		window.LabelWrap(h.dcrSeed)
	}
}

func (h *Handler) generate() {
	mnemonicSeed := string(h.mnemonicSeedInput.Buffer)
	password := string(h.passInput.Buffer)

	bts, err := bip39.NewSeedWithErrorChecking(mnemonicSeed, password)
	if err != nil {
		h.displayError(err)
		return
	}

	defer func() {
		for i := range bts {
			bts[i] = 0
		}
	}()

	h.dcrSeed = hex.EncodeToString(bts)
	h.masterWindow.Changed()
}

func (h *Handler) displayError(err error) {
	popup := func(window *nucular.Window) {
		window.Row(25).Dynamic(1)
		window.Label(err.Error(), "LC")

		window.Row(25).Dynamic(3)
		if window.Button(label.T("Close"), false) {
			window.Close()
		}
	}

	h.masterWindow.PopupOpen("Error", nucular.WindowTitle, h.errorPopupBounds, false, popup)
}
