package main

import "fmt"

type Handler interface {
	Execute(key *Key)
}

type PrintHandler struct{}

func (p *PrintHandler) Execute(key *Key) {
	fmt.Println("Key Pressed:", key.Label)
}
