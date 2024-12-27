package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func main() {
	var actor = NewActor("Alice", 100, 100.0)
	fmt.Println(actor)
}
