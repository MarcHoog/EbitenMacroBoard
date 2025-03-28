package main

import (
	"fmt"
	"macroboard/internal/windows_api"
	"unicode"
)

type EmulateKeyHandler struct{}

func (h *EmulateKeyHandler) Execute(key *Key) {

	title := windows_api.GetCurrentWindowTitle()
	fmt.Printf("Current Windows Title: %s\n", title)
	if unicode.IsPrint(key.RuneValue) {
		windows_api.SendTextUniversal(key.RuneValue)
		return
	}

	fmt.Printf("UnPrintable Unicode: %b\n", key.RuneValue)

}
