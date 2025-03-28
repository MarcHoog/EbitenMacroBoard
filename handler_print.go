package main

import "fmt"

type Handler interface {
	Execute(key *Key) // The function that gets executed when a key is pressed
}

type PrintHandler struct{}

// Execute prints the key's label
func (p *PrintHandler) Execute(key *Key) {
	fmt.Println("Key Pressed:", key.Label)
}
