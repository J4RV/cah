package fixture

import "github.com/j4rv/golang-stuff/cah"

func getPlayerFixture(name string) *cah.Player {
	return &cah.Player{
		User:             cah.User{Username: name},
		Hand:             []*cah.WhiteCard{},
		WhiteCardsInPlay: []*cah.WhiteCard{},
		Points:           []*cah.BlackCard{},
	}
}
