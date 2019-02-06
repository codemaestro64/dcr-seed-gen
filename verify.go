package main

import (
	"image/color"
	"strconv"

	"github.com/aarzilli/nucular"
)

type s struct {
	add nucular.TextEditor
}

func (handler *RenderHandler) renderVerify(window *nucular.Window) {
	window.Row(20).Dynamic(1)
	window.Label("Verify: ", "LC")

	window.Row(165).Dynamic(1)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(150).Dynamic(5)
		//currentItem := 0
		for _, i := range handler.inputs {
			i.Edit(group)
		}
		group.GroupEnd()
	}

	if handler.verifyMessage.message != "" {
		var color color.RGBA

		switch handler.verifyMessage.messageType {
		case "error":
			color = colorDanger
		case "success":
			color = colorSuccess
		}

		window.Row(20).Dynamic(1)
		window.LabelColored(handler.verifyMessage.message, "LC", color)
	}

	window.Row(40).Dynamic(4)
	if window.ButtonText("Verify") {
		msg := &verifyMessage{}
		if handler.doVerify(window) {
			msg.message = "Verification successfull !!"
			msg.messageType = "success"
		} else {
			msg.message = "Invalid mnemonic"
			msg.messageType = "error"
		}
		handler.verifyMessage = msg
	}
	if window.ButtonText("Back") {
		*handler.currentPage = "home"
		window.Master().Changed()
	}
}

func newInputColumn(window *nucular.Window, inputs []nucular.TextEditor, currentItem *int) {
	if group := window.GroupBegin(strconv.Itoa(*currentItem), 0); group != nil {
		for _, input := range inputs {
			group.Row(25).Ratio(0.2, 0.8)
			group.Label(strconv.Itoa(*currentItem+1)+". ", "LC")
			input.Edit(group)
			*currentItem++
		}
		group.GroupEnd()
	}
}

func (handler *RenderHandler) doVerify(window *nucular.Window) bool {
	for index, word := range handler.wordSlice {
		if string(handler.inputs[index].Buffer) != word {
			return false
		}
	}
	return true
}

/**
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

	return true

}
**/
