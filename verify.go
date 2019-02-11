package main

import (
	"image/color"
	"strconv"

	"github.com/aarzilli/nucular"
)

func (handler *RenderHandler) renderVerify(window *nucular.Window) {
	window.Row(20).Dynamic(1)
	SetFont(window, boldFont)
	window.Label("Verify: ", "LC")

	SetFont(window, normalFont)
	window.Row(235).Dynamic(1)
	if group := window.GroupBegin("", 0); group != nil {
		group.Row(220).Dynamic(5)
		currentItem := 0
		for index := range handler.columns {
			newInputColumn(group, handler.columns[index].inputs, &currentItem)
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

	window.Row(40).Ratio(0.5, 0.25, 0.25)
	window.Label("", "LC")
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
		for index := range inputs {
			group.Row(25).Ratio(0.2, 0.8)
			group.Label(strconv.Itoa(*currentItem+1)+". ", "LC")
			inputs[index].Edit(group)

			*currentItem++
		}
		group.GroupEnd()
	}
}

func (handler *RenderHandler) doVerify(window *nucular.Window) bool {
	for _ = range handler.columns {
		for columnIndex := range handler.columns {
			for itemIndex := range handler.columns[columnIndex].words {
				if handler.columns[columnIndex].words[itemIndex] != string(handler.columns[columnIndex].inputs[itemIndex].Buffer) {
					return false
				}
			}
		}
	}
	return true
}
