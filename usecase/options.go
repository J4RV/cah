package usecase

import (
	"log"
	"math/rand"

	"github.com/j4rv/cah"
)

// Options contains methods to alter a gamestate
type Options struct {
	cards cah.CardUsecases
}

func applyOptions(s *cah.GameState, opts ...cah.Option) {
	for _, opt := range opts {
		opt(s)
	}
}

// HandSize changes the gamestate's hand size
func (Options) HandSize(size int) cah.Option {
	return func(s *cah.GameState) {
		s.HandSize = size
	}
}

// WhiteDeck changes the gamestate's white deck
func (Options) WhiteDeck(wd []*cah.WhiteCard) cah.Option {
	return func(s *cah.GameState) {
		s.WhiteDeck = wd
		shuffleW(&s.WhiteDeck)
	}
}

// BlackDeck changes the gamestate's black deck
func (Options) BlackDeck(bd []*cah.BlackCard) cah.Option {
	return func(s *cah.GameState) {
		s.BlackDeck = bd
		shuffleB(&s.BlackDeck)
	}
}

// RandomStartingCzar set's the gamestate to start with a random czar
func (Options) RandomStartingCzar() cah.Option {
	return func(s *cah.GameState) {
		if len(s.Players) == 0 {
			log.Println("WARNING Tried to call RandomStartingCzar using a game without players")
			return
		}
		s.CurrCzarIndex = rand.Intn(len(s.Players))
	}
}

// MaxRounds changes the gamestate's max rounds
func (Options) MaxRounds(max int) cah.Option {
	return func(s *cah.GameState) {
		s.MaxRounds = max
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
