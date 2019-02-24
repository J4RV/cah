package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

func handleGameStates(r *mux.Router) {
	s := r.PathPrefix("/gamestate/{gameStateID}").Subrouter()
	s.Handle("/state", srvHandler(gameStateForUser)).Methods("GET")
	s.Handle("/choose-winner", srvHandler(chooseWinner)).Methods("POST")
	s.Handle("/play-cards", srvHandler(playCards)).Methods("POST")
}

/*
GET GAME STATE
*/

type playerInfo struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	HandSize         int             `json:"handSize"`
	WhiteCardsInPlay int             `json:"whiteCardsInPlay"`
	Points           []cah.BlackCard `json:"points"`
}

type fullPlayerInfo struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	Hand             []cah.WhiteCard `json:"hand" db:"hand"`
	WhiteCardsInPlay []cah.WhiteCard `json:"whiteCardsInPlay"`
	Points           []cah.BlackCard `json:"points"`
}

type sinnerPlay struct {
	ID         int             `json:"id"`
	WhiteCards []cah.WhiteCard `json:"whiteCards"`
}

type gameStateResponse struct {
	ID              int            `json:"id"`
	Phase           string         `json:"phase"`
	Players         []playerInfo   `json:"players"`
	CurrCzarID      int            `json:"currentCzarID"`
	BlackCardInPlay cah.BlackCard  `json:"blackCardInPlay"`
	SinnerPlays     []sinnerPlay   `json:"sinnerPlays"`
	MyPlayer        fullPlayerInfo `json:"myPlayer"`
}

func gameStateForUser(w http.ResponseWriter, req *http.Request) error {
	u, err := userFromSession(req)
	if err != nil {
		return err
	}
	game, err := gameStateFromRequest(req)
	if err != nil {
		return err
	}
	p, err := player(game, u)
	if err != nil {
		return err
	}
	response := gameStateResponse{
		ID:              game.ID,
		Phase:           game.Phase.String(),
		Players:         playersInfoFromGame(game),
		CurrCzarID:      game.Players[game.CurrCzarIndex].User.ID,
		BlackCardInPlay: *game.BlackCardInPlay,
		SinnerPlays:     sinnerPlaysFromGame(game),
		MyPlayer:        newFullPlayerInfo(*p),
	}
	writeResponse(w, response)
	return nil
}

func playersInfoFromGame(game cah.GameState) []playerInfo {
	ret := make([]playerInfo, len(game.Players))
	for i, p := range game.Players {
		ret[i] = newPlayerInfo(*p)
	}
	return ret
}

func newPlayerInfo(p cah.Player) playerInfo {
	return playerInfo{
		ID:               p.User.ID,
		Name:             p.User.Username,
		HandSize:         len(p.Hand),
		WhiteCardsInPlay: len(p.WhiteCardsInPlay),
		Points:           dereferenceBlackCards(p.Points),
	}
}

func newFullPlayerInfo(player cah.Player) fullPlayerInfo {
	return fullPlayerInfo{
		ID:               player.User.ID,
		Name:             player.User.Username,
		Hand:             dereferenceWhiteCards(player.Hand),
		WhiteCardsInPlay: dereferenceWhiteCards(player.WhiteCardsInPlay),
	}
}

func sinnerPlaysFromGame(game cah.GameState) []sinnerPlay {
	if !usecase.GameState.AllSinnersPlayedTheirCards(game) {
		return []sinnerPlay{}
	}
	ret := make([]sinnerPlay, len(game.Players))
	for i, p := range game.Players {
		if game.IsCurrCzar(p.User) {
			continue
		}
		ret[i] = sinnerPlay{
			ID:         p.User.ID,
			WhiteCards: dereferenceWhiteCards(p.WhiteCardsInPlay),
		}
	}
	return ret
}

/*
CHOOSE WINNER
*/

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
}

func chooseWinner(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload chooseWinnerPayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	game, err := gameStateFromRequest(req)
	if err != nil {
		return err
	}
	pid, err := playerIndex(game, u)
	if err != nil {
		return err
	}
	if pid != game.CurrCzarIndex {
		return errors.New("Only the Czar can choose the winner")
	}
	_, err = usecase.GameState.GiveBlackCardToWinner(payload.Winner, game)
	if err != nil {
		return err
	}
	return nil
}

/*
PLAY CARDS
*/

type playCardsPayload struct {
	CardIndexes []int `json:"cardIndexes"`
}

func playCards(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload playCardsPayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	game, err := gameStateFromRequest(req)
	if err != nil {
		return err
	}
	pid, err := playerIndex(game, u)
	if err != nil {
		return err
	}
	_, err = usecase.GameState.PlayWhiteCards(pid, payload.CardIndexes, game)
	if err != nil {
		return err
	}
	return nil
}

// Utils

func playerIndex(g cah.GameState, u cah.User) (int, error) {
	for i, p := range g.Players {
		if p.User.ID == u.ID {
			return i, nil
		}
	}
	return -1, errors.New("You are not playing this game")
}

func player(g cah.GameState, u cah.User) (*cah.Player, error) {
	i, err := playerIndex(g, u)
	if err != nil {
		return &cah.Player{}, errors.New("You are not playing this game")
	}
	return g.Players[i], nil
}

func gameStateFromRequest(req *http.Request) (cah.GameState, error) {
	strID := mux.Vars(req)["gameStateID"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return cah.GameState{}, err
	}
	g, err := usecase.GameState.ByID(id)
	if err != nil {
		return g, errors.New("Could not get game state from request")
	}
	return g, nil
}
