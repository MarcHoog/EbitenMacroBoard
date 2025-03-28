package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type Key struct {
	Label         string
	RuneValue     rune // Only support 16 bit unicodes for a now
	X, Y          int
	Width, Height int
	Hovered       bool
	Clicked       bool
	Handler       Handler
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

func (k *Key) Draw(screen *ebiten.Image, fontScale float64) error {
	// Create a key background image.
	keyImg := ebiten.NewImage(k.Width, k.Height)

	if k.Hovered && !k.Clicked {
		keyImg.Fill(color.RGBA{
			R: 36,
			G: 36,
			B: 36,
			A: 255,
		})
	} else if k.Hovered && k.Clicked {
		keyImg.Fill(color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		})

	} else {
		keyImg.Fill(color.RGBA{
			R: 72,
			G: 72,
			B: 72,
			A: 255,
		})
	}
	// Draw the key background onto the screen.
	opKey := &ebiten.DrawImageOptions{}
	opKey.GeoM.Translate(float64(k.X), float64(k.Y))
	screen.DrawImage(keyImg, opKey)

	// Use basicfont.Face7x13 as the base font.
	// Base dimensions for one character:
	baseCharWidth, baseCharHeight := 7, 13

	// Calculate text dimensions.
	textWidth := len(k.Label) * baseCharWidth
	textHeight := baseCharHeight

	// Create an offscreen image for the letter.
	letterImg := ebiten.NewImage(textWidth, textHeight)
	letterImg.Fill(color.Transparent)
	// Render the label onto the offscreen image.
	text.Draw(letterImg, k.Label, basicfont.Face7x13, 0, textHeight, color.Black)

	// Prepare transformation to scale the letter.
	opText := &ebiten.DrawImageOptions{}
	opText.GeoM.Scale(fontScale, fontScale)

	// Calculate the scaled dimensions.
	scaledWidth := float64(textWidth) * fontScale
	scaledHeight := float64(textHeight) * fontScale

	// Center the scaled text within the key square.
	textX := float64(k.X) + float64(k.Width)/2 - scaledWidth/2
	textY := float64(k.Y) + float64(k.Height)/2 - scaledHeight/2
	opText.GeoM.Translate(textX, textY)

	// Draw the scaled letter onto the screen.
	screen.DrawImage(letterImg, opText)

	return nil
}
