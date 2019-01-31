package main

import (
	"strconv"
	"strings"

	"github.com/aarzilli/nucular"
	bip39 "github.com/tyler-smith/go-bip39"
)

type RenderHandler struct {
	words     []string
	err       error
	verifyErr string

	wordInputs []nucular.TextEditor

	passphraseInput nucular.TextEditor
	seed            string
	currentPage     *string
}

func (h *RenderHandler) beforeRender(currentPage *string) {
	h.passphraseInput.Flags = nucular.EditSimple
	h.passphraseInput.PasswordChar = '*'

	h.currentPage = currentPage

	// get bip39 mnemonic words
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		h.err = err
		return
	}

	words, err := bip39.NewMnemonic(entropy)
	if err != nil {
		h.err = err
		return
	}

	h.words = strings.Split(words, " ")
	h.wordInputs = make([]nucular.TextEditor, len(h.words))
	for index := range h.words {
		editor := nucular.TextEditor{}
		editor.Flags = nucular.EditSimple
		h.wordInputs[index] = editor
	}
}

func (h *RenderHandler) renderHome(window *nucular.Window) {
	if h.err != nil {
		window.Row(20).Dynamic(1)
		window.Label(h.err.Error(), "LC")
	}

	window.Row(140).Ratio(0.2, 0.8)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(20).Dynamic(1)
		group.Label("Mnemonic Words:", "LC")
		group.GroupEnd()
	}

	if group := window.GroupBegin("", 0); group != nil {
		group.Row(20).Dynamic(5)
		for index, value := range h.words {
			word := strconv.Itoa(index+1) + ". " + value
			group.Label(word, "LC")
		}
		group.GroupEnd()
	}

	window.Row(30).Ratio(0.2, 0.8)
	window.Label("Passphrase (Optional): ", "LC")
	h.passphraseInput.Edit(window)

	window.Row(30).Ratio(0.3, 0.7)
	window.Label("", "LC")
	if window.ButtonText("Next") {
		*h.currentPage = "verify"
		window.Master().Changed()
	}
}
