package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"macroboard/internal/windows_api"
)

type Game struct {
	keyboard *KeyBoard
}

func NewGame() *Game {

	fp, err := getConfigFilePath()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config Path:", fp)
	config, err := NewConfigFromFile(fp)
	if err != nil {
		panic(err)
	}
	kb := NewKeyBoard(&config.KeyBoardOptions)
	for _, k := range config.Keys {
		if len(k.RuneValue) > 1 {
			// DO SOMETHING?
		}
		r := []rune(k.RuneValue)

		kb.RegisterKey(k.Label, r[0])
	}

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
