package main

import (
	"github.com/Claude47/poker/pkg/components"
	"github.com/Claude47/poker/pkg/player"
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

	player := player.CreateNewPlayer("Claude")
	player.Print()
}
