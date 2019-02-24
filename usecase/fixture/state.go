package fixture

import (
	"github.com/j4rv/golang-stuff/cah"
)

func getStateFixture(playernames []string) cah.GameState {
	p := make([]*cah.Player, len(playernames))
	for i := range playernames {
		p[i] = getPlayerFixture(playernames[i])
	}
	return cah.GameState{
		BlackDeck:   getBlackCardsFixture(20),
		WhiteDeck:   getWhiteCardsFixture(40),
		DiscardPile: []*cah.WhiteCard{},
		Players:     p,
	}
}

var states = []struct {
	whitesAmount, blacksAmount int
	players                    []string
}{
	{40, 20, []string{"Player1", "Player2", "Player3"}},
}
