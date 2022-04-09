package components

type Card struct {
	Suite Suite
	Value Value
}

func CreateCard(s Suite, v Value) *Card {
	return &Card{
		s, v,
	}
}
