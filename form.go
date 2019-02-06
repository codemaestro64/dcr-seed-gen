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

type wordInputColumn struct {
	words  []string
	inputs []nucular.TextEditor
}

type RenderHandler struct {
	words   string
	columns []wordInputColumn
	err     error

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
	// generate mnemonic words
	words, err := generateWords()
	if err != nil {
		h.err = err
		return
	}
	h.words = words

	// generate hex seed
	err = h.generateSeed()
	if err != nil {
		h.err = err
	}

	h.buildColumns(words)
}

func (h *RenderHandler) buildColumns(words string) {
	wordSlice := strings.Split(words, " ")
	h.columns = make([]wordInputColumn, noColumns)

	currentColumn := 0
	for index, word := range wordSlice {
		h.columns[currentColumn].words = append(h.columns[currentColumn].words, word)
		h.columns[currentColumn].inputs = append(h.columns[currentColumn].inputs, nucular.TextEditor{})

		if index > 0 && (index+1)%5 == 0 {
			currentColumn++
		}
	}
}

func generateWords() (string, error) {
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
	return nil
}

func (h *RenderHandler) renderHome(window *nucular.Window) {
	if h.err != nil {
		window.Row(20).Dynamic(1)
		window.Label(h.err.Error(), "LC")
	}

	window.Row(20).Dynamic(1)
	SetFont(window, boldFont)
	window.Label("Mnemonic Words:", "LC")

	window.Row(137).Dynamic(1)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(121).Dynamic(5)
		SetFont(group, boldFont)

		currentItem := 0
		for _, column := range h.columns {
			newWordColumn(group, column.words, &currentItem)
		}
		group.GroupEnd()
	}

	if h.err != nil {
		window.Row(30).Dynamic(1)
		window.Label(fmt.Sprintf("error generating seed: %s", h.err.Error()), "LC")
	} else {
		window.Row(1).Dynamic(1)
		window.Label("", "LC")

		window.Row(20).Dynamic(1)
		SetFont(window, boldFont)
		window.Label("Hex Seed", "LC")
		window.Row(60).Dynamic(1)
		SetFont(window, normalFont)
		window.LabelWrap(h.seed)
	}

	window.Row(30).Dynamic(1)
	SetFont(window, normalFont)
	window.Label("Passphrase (Optional): ", "LC")

	SetFont(window, normalFont)
	h.passphraseInput.Edit(window)

	currentPassword := string(h.passphraseInput.Buffer)
	if h.livePassword != currentPassword {
		h.err = h.generateSeed()
		h.livePassword = currentPassword
	}

	window.Row(40).Ratio(0.5, 0.25, 0.25)
	window.Label("", "LC")

	if window.ButtonText("Verify") {
		h.verifyMessage = &verifyMessage{}
		*h.currentPage = "verify"
		window.Master().Changed()
	}

	if window.ButtonText("Regenerate") {
		h.generate()
		window.Master().Changed()
	}
}

func newWordColumn(window *nucular.Window, words []string, currentItem *int) {
	if group := window.GroupBegin(words[0], 0); group != nil {
		SetFont(group, normalFont)
		for _, word := range words {
			group.Row(20).Dynamic(1)
			group.Label(strconv.Itoa(*currentItem+1)+". "+word, "LC")
			*currentItem++
		}
		group.GroupEnd()
	}
}
