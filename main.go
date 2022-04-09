package main

import (
	"github.com/Claude47/poker/pkg/components"
)

func main() {
	// unshuffled deck
	deck := components.CreateNewDeck()

	// print
	deck.Print()

	// shuffle deck
	deck.Shuffle()

	// print
	deck.Print()
}
