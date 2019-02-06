package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/aarzilli/nucular"
	bip39 "github.com/tyler-smith/go-bip39"
)

type verifyMessage struct {
	message     string
	messageType string
}

type RenderHandler struct {
	words        string
	wordSlice    []string
	inputs       []nucular.TextEditor
	inputColumns [][]nucular.TextEditor
	err          error

	passphraseInput nucular.TextEditor
	livePassword    string
	seed            string
	currentPage     *string

	verifyMessage *verifyMessage
}

const (
	entropyBitSize   = 256 // will produce 24 words
	entropyNoOfWords = 24
	noColumns        = 5
)

func (h *RenderHandler) beforeRender(currentPage *string) {
	h.passphraseInput.Flags = nucular.EditSimple
	h.passphraseInput.PasswordChar = '*'
	h.livePassword = ""
	h.currentPage = currentPage
	h.generate()
}

func (h *RenderHandler) generate() {
	h.words, h.err = h.generateWords()
	if h.err != nil {
		return
	}

	h.err = h.generateSeed()
}

func (h *RenderHandler) generateWords() (string, error) {
	// get bip39 mnemonic words
	entropy, err := bip39.NewEntropy(entropyBitSize)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(entropy)
}

func (h *RenderHandler) generateSeed() error {
	// generate seed
	bts, err := bip39.NewSeedWithErrorChecking(h.words, string(h.passphraseInput.Buffer))
	if err != nil {
		return err
	}
	h.seed = hex.EncodeToString(bts)
	h.wordSlice = strings.Split(h.words, " ")
	h.inputs = make([]nucular.TextEditor, 5)

	for _ = range h.inputs {
		h.inputs = append(h.inputs, nucular.TextEditor{})
	}

	h.inputColumns = make([][]nucular.TextEditor, noColumns)

	currentItem := 0
	for index, input := range h.inputs {
		h.inputColumns[currentItem] = append(h.inputColumns[currentItem], input)
		if index > 0 && (index+1)%5 == 0 {
			currentItem++
		}
	}

	return nil
}

func (h *RenderHandler) renderHome(window *nucular.Window) {
	if h.err != nil {
		window.Row(20).Dynamic(1)
		window.Label(h.err.Error(), "LC")
	}

	window.Row(20).Dynamic(1)
	window.Label("Mnemonic Words:", "LC")

	window.Row(140).Dynamic(1)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(120).Dynamic(5)
		SetFont(group, boldFont)

		currentItem := 0
		columns := make([][]string, noColumns)
		for index, word := range h.wordSlice {
			columns[currentItem] = append(columns[currentItem], word)
			if index > 0 && (index+1)%noColumns == 0 {
				currentItem++
			}
		}

		currentItem = 0
		for _, column := range columns {
			newWordColumn(group, column, &currentItem)
		}
		group.GroupEnd()
	}

	window.Row(30).Dynamic(1)
	window.Label("Passphrase (Optional): ", "LC")
	h.passphraseInput.Edit(window)

	currentPassword := string(h.passphraseInput.Buffer)
	if h.livePassword != currentPassword {
		h.err = h.generateSeed()
		h.livePassword = currentPassword
	}

	window.Row(30).Dynamic(1)
	if h.err != nil {
		window.Label(fmt.Sprintf("error generating seed: %s", h.err.Error()), "LC")
	} else {
		window.Label("", "LC")
		window.Label("Hex Seed", "LC")
		window.Row(70).Dynamic(1)
		window.LabelWrap(h.seed)
	}

	window.Row(40).Ratio(0.5, 0.25, 0.25)
	window.Label("", "LC")

	if window.ButtonText("Verify") {
		h.verifyMessage = &verifyMessage{}
		*h.currentPage = "verify"
		//h.words.inputs[0].Active = true
		window.Master().Changed()
	}

	if window.ButtonText("Regenerate") {
		h.generate()
		window.Master().Changed()
	}
}

func newWordColumn(window *nucular.Window, words []string, currentItem *int) {
	if group := window.GroupBegin(words[0], 0); group != nil {
		for _, word := range words {
			group.Row(20).Dynamic(1)
			group.Label(strconv.Itoa(*currentItem+1)+". "+word, "LC")
			*currentItem++
		}
		group.GroupEnd()
	}
}
