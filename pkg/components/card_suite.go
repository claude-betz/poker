package components

type Suite int

const (
	Spades Suite = iota + 1
	Clubs
	Hearts
	Diamonds
)

// String - Creating common behavior - give the type a String function
func (s Suite) String() string {
	return [...]string{"Spades", "Clubs", "Hearts", "Diamonds"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (s Suite) EnumIndex() int {
	return int(s)
}
