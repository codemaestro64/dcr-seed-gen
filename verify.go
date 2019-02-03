package main

import (
	"strconv"

	"github.com/aarzilli/nucular"
)

func (handler *RenderHandler) renderVerify(window *nucular.Window) {
	window.Row(20).Dynamic(1)
	window.Label("Verify: ", "LC")

	window.Row(20).Ratio(0.04, 0.2, 0.04, 0.2, 0.04, 0.2, 0.04, 0.2)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(120).Dynamic(5)
		//num := 1
		for index, column := range handler.wordInputs {
			if subgroup := group.GroupBegin(strconv.Itoa(index), 0); subgroup != nil {
				subgroup.Row(20).Dynamic(1)
				for _, input := range column {
					input.Edit(subgroup)
				}
				subgroup.GroupEnd()
			}
		}
		group.GroupEnd()
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
			//handler.generateSeed()
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
	/**for index, word := range handler.words {
		if string(handler.wordInputs[index].Buffer) != word {
			return false
		}
	}

	return true**/
	return false
}
