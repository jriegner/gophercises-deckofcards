package main

import "github.com/jriegner/gophercises-deckofcards/deck"

func main() {
	cards := deck.New()

	for _, c := range cards {
		c.Print()
	}
}
