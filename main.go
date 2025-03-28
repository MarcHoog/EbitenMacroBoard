package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"macroboard/internal/windows_api"
)

type Game struct {
	keyboard KeyBoard
}

func NewGame() *Game {

	kb := KeyBoard{
		Options: KeyBoardOptions{
			Padding:   20,
			Spacing:   5,
			KeyWidth:  100,
			KeyHeight: 100,
			FontScale: 2.0,
		},
	}
	kb.CalcBpSize()

	kb.RegisterKey("A", "A")
	kb.RegisterKey("B", "B")
	kb.RegisterKey("C", "C")
	return &Game{keyboard: kb}
}

func (g *Game) Update() (err error) {
	g.keyboard.Update()
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.keyboard.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.keyboard.bpWidth, g.keyboard.bpHeight
}

func main() {

	go windows_api.TrackActiveWindow()

	g := NewGame()
	ebiten.SetWindowSize(g.keyboard.bpWidth, g.keyboard.bpHeight)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowDecorated(true)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowTitle("MacroBoard")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
