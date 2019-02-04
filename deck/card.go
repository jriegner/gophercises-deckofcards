package deck

import "fmt"

type CardColor int

const (
	Diamond CardColor = 1
	Heart             = 2
	Spade             = 3
	Cross             = 4
)

type Card struct {
	Color CardColor
	Value int
}

var colors = map[CardColor]string{
	Diamond: "Diamond",
	Heart:   "Heart",
	Spade:   "Spade",
	Cross:   "Cross",
}

func New() []Card {
	cards := make([]Card, 0)
	for i := 1; i < 14; i++ {
		cards = append(cards, Card{Color: Diamond, Value: i})
	}
	for i := 1; i < 14; i++ {
		cards = append(cards, Card{Color: Heart, Value: i})
	}
	for i := 1; i < 14; i++ {
		cards = append(cards, Card{Color: Spade, Value: i})
	}
	for i := 1; i < 14; i++ {
		cards = append(cards, Card{Color: Cross, Value: i})
	}
	return cards
}

func (c *Card) Print() {
	fmt.Printf("Color: %s Value:%d\n", colors[c.Color], c.Value)
}
