package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"image/color"
	"log"
)

type KeyBoardOptions struct {
	Padding   int     `yaml:"padding"`
	Spacing   int     `yaml:"spacing"`
	KeyWidth  int     `yaml:"keyWidth"`
	KeyHeight int     `yaml:"KeyHeight"`
	FontScale float64 `yaml:"FontScale"`
}

var DefaultKeyBoardOptions = KeyBoardOptions{
	Padding:   2,
	Spacing:   2,
	KeyWidth:  100,
	KeyHeight: 100,
	FontScale: 5,
}

type KeyBoard struct {
	Options  KeyBoardOptions
	Keys     []*Key
	FontFace font.Face
	bpWidth  int
	bpHeight int
}

func NewKeyBoard(options *KeyBoardOptions) *KeyBoard {

	if options != nil {

		if options.Padding < 1 {
			options.Padding = DefaultKeyBoardOptions.Padding
		}
		if options.Spacing < 1 {
			options.Spacing = DefaultKeyBoardOptions.Spacing
		}
		if options.KeyWidth < 1 {
			options.KeyWidth = DefaultKeyBoardOptions.KeyWidth
		}
		if options.KeyHeight < 1 {
			options.KeyHeight = DefaultKeyBoardOptions.KeyHeight
		}
		if options.FontScale < 1 {
			options.FontScale = DefaultKeyBoardOptions.FontScale
		}
	} else {
		options = &DefaultKeyBoardOptions
	}

	fontFace, err := CreateFontFace(options.FontScale)
	if err != nil {
		log.Printf("Failed to load font: %v", err)
		fontFace = FallbackFontFace(options.FontScale)
	}

	keyboard := &KeyBoard{Options: *options, FontFace: fontFace}
	return keyboard

}

func (kb *KeyBoard) CalcBpSize() {
	kb.bpWidth = kb.Options.Padding*2 + len(kb.Keys)*kb.Options.KeyWidth + (len(kb.Keys)-1)*kb.Options.Spacing
	kb.bpHeight = kb.Options.Padding*2 + kb.Options.KeyHeight

}

func (kb *KeyBoard) RegisterKey(label string, runeValue rune) {

	kb.Keys = append(kb.Keys, &Key{
		Label:     label,
		RuneValue: runeValue,
		X:         kb.Options.Padding + len(kb.Keys)*(kb.Options.KeyWidth+kb.Options.Spacing),
		Y:         kb.Options.Padding,
		Height:    kb.Options.KeyHeight,
		Width:     kb.Options.KeyWidth,
		Font:      kb.FontFace,
		Handler:   &EmulateKeyHandler{},
	})

	kb.CalcBpSize()

}

func (kb *KeyBoard) Draw(screen *ebiten.Image) (err error) {
	backPlate := ebiten.NewImage(kb.bpWidth, kb.bpHeight)
	backPlate.Fill(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	backPlateX, backPlateY := 0.0, 0.0
	opBP := &ebiten.DrawImageOptions{}
	opBP.GeoM.Translate(backPlateX, backPlateY)
	screen.DrawImage(backPlate, opBP)
	for _, k := range kb.Keys {
		k.Draw(screen, kb.Options.FontScale)
	}

	return
}
func (kb *KeyBoard) Update() (err error) {
	mouseX, mouseY := ebiten.CursorPosition()

	for _, k := range kb.Keys {
		k.Update(mouseX, mouseY)

	}
	return
}
