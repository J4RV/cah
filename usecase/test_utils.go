package usecase

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

// consider using mockery or something similar
func getWhiteCardsFixture(amount int) []*cah.WhiteCard {
	ret := make([]*cah.WhiteCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = &cah.WhiteCard{Text: fmt.Sprintf("White card fixture (%d)", i)}
	}
	return ret
}

func getBlackCardsFixture(amount int) []*cah.BlackCard {
	ret := make([]*cah.BlackCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = &cah.BlackCard{Text: fmt.Sprintf("Black card fixture (%d)", i), Blanks: 1}
	}
	return ret
}

func getPlayerFixture(name string) *cah.Player {
	return &cah.Player{
		User:             cah.User{Username: name},
		Hand:             []*cah.WhiteCard{},
		WhiteCardsInPlay: []*cah.WhiteCard{},
		Points:           []*cah.BlackCard{},
	}
}

func getStateFixture() cah.GameState {
	return cah.GameState{
		BlackDeck:       getBlackCardsFixture(20),
		WhiteDeck:       getWhiteCardsFixture(40),
		DiscardPile:     []*cah.WhiteCard{},
		BlackCardInPlay: nilBlackCard,
		Players: []*cah.Player{
			getPlayerFixture("Player1"),
			getPlayerFixture("Player2"),
			getPlayerFixture("Player3"),
		},
	}
}
