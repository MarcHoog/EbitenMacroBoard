package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type KeyBoardOptions struct {
	Padding   int
	Spacing   int
	KeyWidth  int
	KeyHeight int
	FontScale float64
}

type KeyBoard struct {
	Options  KeyBoardOptions
	Keys     []*Key
	bpWidth  int
	bpHeight int
}

func (kb *KeyBoard) CalcBpSize() {
	kb.bpWidth = kb.Options.Padding*2 + len(kb.Keys)*kb.Options.KeyWidth + (len(kb.Keys)-1)*kb.Options.Spacing
	kb.bpHeight = kb.Options.Padding*2 + kb.Options.KeyHeight

}

func (kb *KeyBoard) RegisterKey(label, text string) {

	kb.Keys = append(kb.Keys, &Key{Label: label,
		X:       kb.Options.Padding + len(kb.Keys)*(kb.Options.KeyWidth+kb.Options.Spacing),
		Y:       kb.Options.Padding,
		Height:  kb.Options.KeyHeight,
		Width:   kb.Options.KeyWidth,
		Handler: &EmulateKeyHandler{},
	})

	kb.CalcBpSize()

}

func (kb *KeyBoard) Draw(screen *ebiten.Image) (err error) {
	// Draw the Backplate

	backPlate := ebiten.NewImage(kb.bpWidth, kb.bpHeight)
	backPlate.Fill(color.RGBA{R: 100, G: 100, B: 100, A: 255})

	// Decide where to position the backplate on the screen.
	// Here, we use a fixed offset (e.g., 50, 50).
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
