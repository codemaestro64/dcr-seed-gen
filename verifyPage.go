package main

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/aarzilli/nucular"
	bip39 "github.com/tyler-smith/go-bip39"
)

func (handler *RenderHandler) renderVerify(window *nucular.Window) {
	window.Row(20).Dynamic(1)
	window.Label("Verify: ", "LC")

	window.Row(20).Ratio(0.04, 0.2, 0.04, 0.2, 0.04, 0.2, 0.04, 0.2)
	for i := range handler.wordInputs {
		window.Label(strconv.Itoa(i+1), "LC")
		handler.wordInputs[i].Edit(window)
	}

	if handler.verifyErr != "" {
		window.Row(25).Dynamic(1)
		window.Label(handler.verifyErr, "LC")
	}

	window.Row(30).Dynamic(3)
	if window.ButtonText("Submit") {
		handler.verifyErr = ""
		window.Master().Changed()

		if handler.verifyIfValid() {
			handler.generateSeed()
		} else {
			handler.verifyErr = "Verification error"
		}
		window.Master().Changed()
	}

	if handler.seed != "" {
		window.Row(20).Dynamic(1)
		window.Label("Hex Seed", "LC")

		window.Row(60).Dynamic(1)
		window.LabelWrap(handler.seed)
	}
}

func (handler *RenderHandler) verifyIfValid() bool {
	for index, word := range handler.words {
		if string(handler.wordInputs[index].Buffer) != word {
			return false
		}
	}

	return true
}

func (handler *RenderHandler) generateSeed() {
	str := strings.Join(handler.words, " ")

	bts, err := bip39.NewSeedWithErrorChecking(str, string(handler.passphraseInput.Buffer))
	if err != nil {
		handler.verifyErr = err.Error()
		return
	}

	handler.seed = hex.EncodeToString(bts)
	return
}
