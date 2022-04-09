package main

import (
	"github.com/Claude47/poker/pkg/components"
)

func main() {
	deck := components.CreateNewDeck()
	deck.PrintDeck()

}
