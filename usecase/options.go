package usecase

import (
	"log"
	"math/rand"

	"github.com/j4rv/golang-stuff/cah"
)

type Options struct {
	cards cah.CardUsecases
}

func applyOptions(s *cah.GameState, opts ...cah.Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (_ Options) HandSize(size int) cah.Option {
	return func(s *cah.GameState) {
		s.HandSize = size
	}
}

func (_ Options) WhiteDeck(wd []*cah.WhiteCard) cah.Option {
	return func(s *cah.GameState) {
		s.WhiteDeck = wd
		shuffleW(&s.WhiteDeck)
	}
}

func (_ Options) BlackDeck(bd []*cah.BlackCard) cah.Option {
	return func(s *cah.GameState) {
		s.BlackDeck = bd
		shuffleB(&s.BlackDeck)
	}
}

func (_ Options) RandomStartingCzar() cah.Option {
	return func(s *cah.GameState) {
		if len(s.Players) == 0 {
			log.Println("WARNING Tried to call RandomStartingCzar using a game without players")
			return
		}
		s.CurrCzarIndex = rand.Intn(len(s.Players))
	}
}

func shuffleB(cards *[]*cah.BlackCard) {
	if cards == nil {
		return
	}
	for i, j := range rand.Perm(len(*cards)) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	}
}

func shuffleW(cards *[]*cah.WhiteCard) {
	if cards == nil {
		return
	}
	for i, j := range rand.Perm(len(*cards)) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	}
}
