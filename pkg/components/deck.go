package components

import (
	"fmt"
	"math/rand"
	"time"
)

var Suites = []Suite{
	Spades,
	Clubs,
	Hearts,
	Diamonds,
}

var Values = []Value{
	Ace,
	King,
	Queen,
	Jack,
	Ten,
	Nine,
	Eight,
	Seven,
	Six,
	Five,
	Four,
	Three,
	Two,
}

type Deck struct {
	cards []Card
}

func CreateNewDeck() *Deck {
	deck := &Deck{}

	cards := &[]Card{}
	for _, suite := range Suites {
		for _, value := range Values {
			*cards = append(*cards, Card{
				Suite: suite,
				Value: value,
			})
		}
	}
	deck.cards = *cards
	return deck
}

func (d *Deck) GetCard(i int) Card {
	return d.cards[i]
}

func (d *Deck) Print() {
	for i := 0; i < len(d.cards); i++ {
		card := d.GetCard(i)
		fmt.Print(card)
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().Unix())

	cards := d.cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	d.cards = cards
}
