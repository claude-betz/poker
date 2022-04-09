package components

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
