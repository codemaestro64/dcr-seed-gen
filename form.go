package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/aarzilli/nucular"
	bip39 "github.com/tyler-smith/go-bip39"
)

type column struct {
	words  []string
	inputs []nucular.TextEditor
}

type RenderHandler struct {
	columns []column
	err     error

	passphraseInput nucular.TextEditor
	seed            string
	currentPage     *string
}

const (
	entropyBitSize   = 256 // will produce 24 words
	entropyNoOfWords = 24
	columns          = 5
)

func (h *RenderHandler) beforeRender(currentPage *string) {
	h.passphraseInput.Flags = nucular.EditSimple
	h.passphraseInput.PasswordChar = '*'

	h.currentPage = currentPage

	// get bip39 mnemonic words
	entropy, err := bip39.NewEntropy(entropyBitSize)
	if err != nil {
		h.err = err
		return
	}

	words, err := bip39.NewMnemonic(entropy)
	if err != nil {
		h.err = err
		return
	}

	// generate seed
	h.generateSeed(words)

	wordSlice := strings.Split(words, " ")
	h.columns = make([]column, 5)
	currentColumnIndex := 0

	for index, word := range wordSlice {
		editor := nucular.TextEditor{}
		h.columns[currentColumnIndex].inputs = append(h.columns[currentColumnIndex].inputs, editor)
		h.columns[currentColumnIndex].words = append(h.columns[currentColumnIndex].words, word)

		if index > 0 && (index+1)%5 == 0 {
			currentColumnIndex++
		}
	}
}

func (handler *RenderHandler) generateSeed(words string) {
	bts, err := bip39.NewSeedWithErrorChecking(words, string(handler.passphraseInput.Buffer))
	if err != nil {
		handler.err = err
		return
	}
	handler.seed = hex.EncodeToString(bts)
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
		group.Row(120).Dynamic(5)

		currentItem := 1
		for index, column := range h.columns {
			if subgroup := group.GroupBegin(strconv.Itoa(index), 0); subgroup != nil {
				subgroup.Row(20).Dynamic(1)
				for _, word := range column.words {
					subgroup.Label(strconv.Itoa(currentItem)+". "+word, "LC")
					currentItem++
				}
				subgroup.GroupEnd()
			}
		}
		group.GroupEnd()
	}

	window.Row(30).Ratio(0.2, 0.8)
	window.Label("Passphrase (Optional): ", "LC")
	h.passphraseInput.Edit(window)

	window.Row(40).Ratio(0.2, 0.4, 0.4)
	window.Label("", "LC")
	if window.ButtonText("Next") {
		*h.currentPage = "verify"
		window.Master().Changed()
	}

	if window.ButtonText("Regenerate") {
		window.Master().Changed()
	}

	window.Row(30).Dynamic(1)
	if h.err != nil {
		window.Label(fmt.Sprintf("error generating seed: %s", h.err.Error()), "LC")
	} else {
		window.Label("Hex Seed", "LC")
		window.Row(70).Dynamic(1)
		window.LabelWrap(h.seed)
	}

}
