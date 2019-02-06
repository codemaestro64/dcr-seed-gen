package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/aarzilli/nucular"
	nstyle "github.com/aarzilli/nucular/style"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	lightFont font.Face
	boldFont  font.Face

	buttonTextColor = color.RGBA{70, 127, 207, 255}
	whiteColor      = color.RGBA{255, 255, 255, 255}
	inputTextColor  = color.RGBA{73, 80, 87, 255}

	colorDanger  = color.RGBA{215, 58, 73, 255}
	colorSuccess = color.RGBA{227, 98, 9, 255}
)

var colorTable = nstyle.ColorTable{
	ColorText:                  whiteColor,
	ColorWindow:                color.RGBA{9, 20, 64, 255},
	ColorHeader:                color.RGBA{175, 175, 175, 255},
	ColorBorder:                color.RGBA{0, 0, 0, 255},
	ColorButton:                buttonTextColor,
	ColorButtonHover:           whiteColor,
	ColorButtonActive:          color.RGBA{0, 153, 204, 255},
	ColorToggle:                color.RGBA{150, 150, 150, 255},
	ColorToggleHover:           color.RGBA{120, 120, 120, 255},
	ColorToggleCursor:          color.RGBA{175, 175, 175, 255},
	ColorSelect:                color.RGBA{175, 175, 175, 255},
	ColorSelectActive:          color.RGBA{190, 190, 190, 255},
	ColorSlider:                color.RGBA{190, 190, 190, 255},
	ColorSliderCursor:          color.RGBA{80, 80, 80, 255},
	ColorSliderCursorHover:     color.RGBA{70, 70, 70, 255},
	ColorSliderCursorActive:    color.RGBA{60, 60, 60, 255},
	ColorProperty:              color.RGBA{175, 175, 175, 255},
	ColorEdit:                  color.RGBA{150, 150, 150, 255},
	ColorEditCursor:            color.RGBA{0, 0, 0, 255},
	ColorCombo:                 color.RGBA{175, 175, 175, 255},
	ColorChart:                 color.RGBA{160, 160, 160, 255},
	ColorChartColor:            color.RGBA{45, 45, 45, 255},
	ColorChartColorHighlight:   color.RGBA{255, 0, 0, 255},
	ColorScrollbar:             color.RGBA{180, 180, 180, 255},
	ColorScrollbarCursor:       color.RGBA{140, 140, 140, 255},
	ColorScrollbarCursorHover:  color.RGBA{150, 150, 150, 255},
	ColorScrollbarCursorActive: color.RGBA{160, 160, 160, 255},
	ColorTabHeader:             color.RGBA{0x89, 0x89, 0x89, 0xff},
}

func loadFonts() error {
	lightFontData, err := ioutil.ReadFile("fonts/SourceSansPro-Light.ttf")
	if err != nil {
		return err
	}

	boldFontData, err := ioutil.ReadFile("fonts/SourceSansPro-Regular.ttf")
	if err != nil {
		return err
	}

	lightFont, err = getFont(13, 72, lightFontData)
	if err != nil {
		return err
	}

	boldFont, err = getFont(13, 72, boldFontData)
	if err != nil {
		return err
	}

	return nil
}

func getFont(fontSize, DPI int, fontData []byte) (font.Face, error) {
	ttfont, err := freetype.ParseFont(fontData)
	if err != nil {
		return nil, err
	}

	size := int(float64(fontSize) * scaling)
	options := &truetype.Options{
		Size:    float64(size),
		Hinting: font.HintingFull,
		DPI:     float64(DPI),
	}

	return truetype.NewFace(ttfont, options), nil
}

func SetFont(window *nucular.Window, font font.Face) {
	masterWindow := window.Master()
	style := masterWindow.Style()
	style.Font = font
	masterWindow.SetStyle(style)
}

func setStyle(window nucular.MasterWindow) error {
	err := loadFonts()
	if err != nil {
		return fmt.Errorf("error loading font: %s", err.Error())
	}

	style := nstyle.FromTable(colorTable, scaling)
	style.Font = lightFont

	// window
	style.NormalWindow.Padding = image.Point{20, 0}

	// buttons
	style.Button.Rounding = 0
	style.Button.TextHover = inputTextColor

	// text input
	style.Edit.Normal.Data.Color = whiteColor
	style.Edit.Hover.Data.Color = whiteColor
	style.Edit.Active.Data.Color = whiteColor
	style.Edit.TextActive = inputTextColor
	style.Edit.TextNormal = inputTextColor
	style.Edit.TextHover = inputTextColor
	style.Edit.CursorHover = inputTextColor

	window.SetStyle(style)

	return nil
}
