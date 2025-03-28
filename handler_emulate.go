package main

import (
	"fmt"
	"macroboard/internal/windows_api"
)

type EmulateKeyHandler struct{}

func (h *EmulateKeyHandler) Execute(key *Key) {

	title := windows_api.GetCurrentWindowTitle()
	fmt.Printf("Current Windows Title: %s\n", title)
	windows_api.SendTextUniversal()
	fmt.Printf("lol\n")

}
