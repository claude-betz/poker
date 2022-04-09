package main

import (
	"fmt"

	"github.com/Claude47/poker/pkg/components"
)

func main() {
	deck := components.CreateNewDeck()

	for i := 0; i < 52; i++ {
		card := deck.GetCard(i)
		fmt.Print(card)
		fmt.Print("\n")
	}

}
