package components

type Value int

const (
	Ace Value = iota + 1
	King
	Queen
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
)

// String - Creating common behavior - give the type a String function
func (v Value) String() string {
	return [...]string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}[v-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (v Value) EnumIndex() int {
	return int(v)
}
