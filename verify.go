package main

import (
	"strconv"

	"github.com/aarzilli/nucular"
)

func (handler *RenderHandler) renderVerify(window *nucular.Window) {
	window.Row(20).Dynamic(1)
	window.Label("Verify: ", "LC")

	window.Row(190).Dynamic(1)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(170).Dynamic(5)
		currentItem := 1
		for index, column := range handler.columns {
			if subgroup := group.GroupBegin(strconv.Itoa(index), 0); subgroup != nil {
				subgroup.Row(28).Ratio(0.22, 0.78)
				for inputIndex, word := range column.words {
					subgroup.Label(strconv.Itoa(currentItem), "RC")
					column.inputs[inputIndex].Buffer = []rune(word)
					column.inputs[inputIndex].Edit(subgroup)
					currentItem++
				}
				subgroup.GroupEnd()
			}
		}
		group.GroupEnd()
	}

	window.Row(40).Dynamic(4)
	if window.ButtonText("Verify") {

	}
	if window.ButtonText("Back") {

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
