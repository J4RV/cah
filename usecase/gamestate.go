package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/j4rv/golang-stuff/cah"
)

var nilBlackCard = &cah.BlackCard{}

type errorEmptyBlackDeck struct{}

func (e errorEmptyBlackDeck) Error() string {
	return "Zero cards left in black deck"
}

type stateController struct {
	store cah.GameStateStore
}

func NewGameStateUsecase(store cah.GameStateStore) *stateController {
	return &stateController{store: store}
}

func (control stateController) Create() cah.GameState {
	ret := cah.GameState{
		Players:         []*cah.Player{},
		HandSize:        10,
		DiscardPile:     []*cah.WhiteCard{},
		WhiteDeck:       []*cah.WhiteCard{},
		BlackDeck:       []*cah.BlackCard{},
		BlackCardInPlay: nilBlackCard,
	}
	ret, err := control.store.Create(ret)
	if err != nil {
		log.Println("Error while creating a new game:", err)
	}
	return ret
}

func (control stateController) ByID(id int) (cah.GameState, error) {
	return control.store.ByID(id)
}

func (control stateController) End(g cah.GameState) (cah.GameState, error) {
	if g.Phase == cah.Finished {
		return g, errors.New("Tried to end a game but it has already finished")
	}
	ret := g.Clone()
	ret.Phase = cah.Finished
	err := control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
}

func (control stateController) GiveBlackCardToWinner(wID int, g cah.GameState) (cah.GameState, error) {
	err := giveBlackCardToWinnerChecks(wID, g)
	if err != nil {
		return g, err
	}
	ret := g.Clone()
	var winner *cah.Player
	for _, p := range ret.Players {
		if p.User.ID == wID {
			winner = p
		}
	}
	if winner == nil {
		return g, fmt.Errorf("Invalid winner id %d", wID)
	}
	winner.Points = append(winner.Points, ret.BlackCardInPlay)
	ret.BlackCardInPlay = nilBlackCard
	for _, p := range g.Players {
		p.WhiteCardsInPlay = []*cah.WhiteCard{}
	}
	ret, _ = control.nextCzar(ret)
	if (len(ret.BlackDeck)) == 0 {
		return control.End(ret)
	}
	ret, err = putBlackCardInPlay(ret)
	if err != nil {
		return g, err
	}
	playersDraw(&ret)
	err = control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
}

func giveBlackCardToWinnerChecks(w int, s cah.GameState) error {
	if s.Phase != cah.CzarChoosingWinner {
		return fmt.Errorf("Tried to choose a winner in a non valid phase '%d'", s.Phase)
	}
	for i, p := range s.Players {
		if i == s.CurrCzarIndex {
			continue
		}
		if len(p.WhiteCardsInPlay) != s.BlackCardInPlay.Blanks {
			return errors.New("Not all sinners have played their cards")
		}
	}
	return nil
}

func (control stateController) PlayWhiteCards(p int, cs []int, g cah.GameState) (cah.GameState, error) {
	if p < 0 || p >= len(g.Players) {
		return g, errors.New("Non valid sinner index")
	}
	if p == g.CurrCzarIndex {
		return g, errors.New("The Czar cannot play white cards")
	}
	if len(g.Players[p].WhiteCardsInPlay) != 0 {
		return g, errors.New("You played your card(s) already")
	}
	if len(cs) != g.BlackCardInPlay.Blanks {
		return g, fmt.Errorf("Invalid amount of white cards to play, expected %d but got %d",
			g.BlackCardInPlay.Blanks,
			len(cs))
	}
	ret := g.Clone()
	player := ret.Players[p]
	newCardsPlayed, err := player.ExtractCardsFromHand(cs)
	if err != nil {
		return ret, err
	}
	player.WhiteCardsInPlay = append(player.WhiteCardsInPlay, newCardsPlayed...)
	if control.AllSinnersPlayedTheirCards(ret) {
		ret.Phase = cah.CzarChoosingWinner
	}
	err = control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
}

func (_ stateController) AllSinnersPlayedTheirCards(s cah.GameState) bool {
	for i, p := range s.Players {
		if i == s.CurrCzarIndex {
			continue
		}
		if len(p.WhiteCardsInPlay) != s.BlackCardInPlay.Blanks {
			return false
		}
	}
	return true
}

func playersDraw(s *cah.GameState) {
	for _, p := range s.Players {
		for len(p.Hand) < s.HandSize {
			p.Hand = append(p.Hand, s.DrawWhite())
		}
	}
}

func putBlackCardInPlay(g cah.GameState) (cah.GameState, error) {
	if g.BlackCardInPlay != nilBlackCard {
		return g, errors.New("Tried to put a black card in play but there is already a black card in play")
	}
	if g.Phase == cah.Finished {
		return g, errors.New("Tried to put a black card in play but the game has already finished")
	}
	if len(g.BlackDeck) == 0 {
		return g, errorEmptyBlackDeck{}
	}
	ret := g.Clone()
	ret.BlackCardInPlay = ret.BlackDeck[0]
	ret.BlackDeck = ret.BlackDeck[1:]
	ret.Phase = cah.SinnersPlaying
	return ret, nil
}

func (_ stateController) nextCzar(s cah.GameState) (cah.GameState, error) {
	if s.BlackCardInPlay != nilBlackCard {
		return s, errors.New("Tried to rotate to the next Czar but there is still a black card in play")
	}
	if s.Phase == cah.Finished {
		return s, errors.New("Tried to rotate to the next Czar but the game has already finished")
	}
	res := s.Clone()
	res.CurrCzarIndex++
	if res.CurrCzarIndex == len(s.Players) {
		res.CurrCzarIndex = 0
	}
	return res, nil
}
