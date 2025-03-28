package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type Key struct {
	Label         string
	RuneValue     rune
	X, Y          int
	Width, Height int
	Hovered       bool
	Clicked       bool
	Handler       Handler
	Font          font.Face
}

func (k *Key) Update(mouseX, mouseY int) {

	k.Hovered = mouseX >= k.X && mouseX < k.X+k.Width &&
		mouseY >= k.Y && mouseY < k.Y+k.Height

	if k.Hovered {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !k.Clicked {
			k.Clicked = true
		} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && k.Clicked {
			k.Handler.Execute(k)
			k.Clicked = false

		}
	} else {
		k.Clicked = false
	}

}

func (k *Key) Draw(screen *ebiten.Image, fontSize float64) error {
	// Determine key color based on state
	var keyColor color.RGBA
	switch {
	case k.Hovered && k.Clicked:
		keyColor = color.RGBA{R: 0, G: 0, B: 0, A: 255} // Black when clicked
	case k.Hovered:
		keyColor = color.RGBA{R: 36, G: 36, B: 36, A: 255} // Dark gray when hovered
	default:
		keyColor = color.RGBA{R: 72, G: 72, B: 72, A: 255} // Medium gray by default
	}

	// Draw key background
	keyImg := ebiten.NewImage(k.Width, k.Height)
	keyImg.Fill(keyColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(k.X), float64(k.Y))
	screen.DrawImage(keyImg, op)

	// Get font face

	// Calculate text position
	bounds := text.BoundString(k.Font, k.Label)
	textWidth := bounds.Max.X - bounds.Min.X
	textHeight := bounds.Max.Y - bounds.Min.Y

	textX := k.X + (k.Width-textWidth)/2
	textY := k.Y + (k.Height+textHeight)/2

	text.Draw(screen, k.Label, k.Font, textX, textY, color.White)

	return nil
}
