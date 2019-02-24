package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

func handleGames(r *mux.Router) {
	s := r.PathPrefix("/game").Subrouter()
	s.Handle("/{gameID}/room-state", srvHandler(roomState)).Methods("GET")
	s.Handle("/list-open", srvHandler(openGames)).Methods("GET")
	s.Handle("/list-in-progress", srvHandler(inProgressGames)).Methods("GET")
	s.Handle("/create", srvHandler(createGame)).Methods("POST")
	s.Handle("/join", srvHandler(joinGame)).Methods("POST")
	//s.Handle("/Leave", srvHandler(playCards)).Methods("POST")
	s.Handle("/start", srvHandler(startGame)).Methods("POST")
	s.Handle("/available-expansions", srvHandler(availableExpansions)).Methods("GET")
}

const minWhites = 34
const minBlacks = 8
const minHandSize = 5
const maxHandSize = 30

/*
OPEN GAMES LIST
*/

type gameRoomResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	HasPassword bool     `json:"hasPassword"`
	Players     []string `json:"players"`
	Phase       string   `json:"phase"`
	StateID     int      `json:"stateID"`
}

func roomState(w http.ResponseWriter, req *http.Request) error {
	g, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	writeResponse(w, gameToResponse(g))
	return nil
}

func openGames(w http.ResponseWriter, req *http.Request) error {
	response := []gameRoomResponse{}
	for _, g := range usecase.Game.AllOpen() {
		response = append(response, gameToResponse(g))
	}
	writeResponse(w, response)
	return nil
}

func inProgressGames(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	response := []gameRoomResponse{}
	for _, g := range usecase.Game.InProgressForUser(u) {
		response = append(response, gameToResponse(g))
	}
	writeResponse(w, response)
	return nil
}

func gameToResponse(g cah.Game) gameRoomResponse {
	players := make([]string, len(g.Users))
	for i := range g.Users {
		players[i] = g.Users[i].Username
	}
	return gameRoomResponse{
		ID:          g.ID,
		Owner:       g.Owner.Username,
		Name:        g.Name,
		HasPassword: g.Password != "",
		Players:     players,
		Phase:       g.State.Phase.String(),
		StateID:     g.State.ID,
	}
}

/*
CREATE GAME
*/

type createGamePayload struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

func createGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload createGamePayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	err = usecase.Game.Create(u, payload.Name, payload.Password)
	if err != nil {
		return err
	}
	return nil
}

/*
JOIN GAME
*/

type joinGamePayload struct {
	ID int `json:"id"`
}

func joinGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload joinGamePayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	g, err := usecase.Game.ByID(payload.ID)
	if err != nil {
		return err
	}
	err = usecase.Game.UserJoins(u, g)
	if err != nil {
		return err
	}
	return nil
}

/*
JOIN GAME
*/

type startGamePayload struct {
	GameID          int      `json:"gameID"`
	Expansions      []string `json:"expansions"`
	HandSize        int      `json:"handSize"`
	RandomFirstCzar bool     `json:"randomFirstCzar,omitempty"`
}

func startGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	var payload startGamePayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return err
	}
	log.Println("Starting game with options:", payload)
	if err != nil {
		return err
	}
	g, err := usecase.Game.ByID(payload.GameID)
	if err != nil {
		return err
	}
	if g.Owner != u {
		return errors.New("Only the game owner can start the game")
	}
	opts, err := optionsFromCreateRequest(payload)
	if err != nil {
		return err
	}
	state := usecase.GameState.Create()
	err = usecase.Game.Start(g, state, opts...)
	if err != nil {
		return err
	}
	return nil
}

func optionsFromCreateRequest(payload startGamePayload) ([]cah.Option, error) {
	ret := []cah.Option{}
	// EXPANSIONS
	exps := payload.Expansions
	blacks := usecase.Card.ExpansionBlacks(exps...)
	whites := usecase.Card.ExpansionWhites(exps...)
	if len(blacks) < minBlacks {
		return ret, fmt.Errorf("Not enough black cards to play a game. Please select more expansions. The amount of Black cards in selected expansions is %d, but the minimum is %d", len(blacks), minBlacks)
	}
	if len(whites) < minWhites {
		return ret, fmt.Errorf("Not enough white cards to play a game. Please select more expansions. The amount of White cards in selected expansions is %d, but the minimum is %d", len(whites), minWhites)
	}
	ret = append(ret, usecase.Game.Options().BlackDeck(blacks))
	ret = append(ret, usecase.Game.Options().WhiteDeck(whites))
	// HAND SIZE
	handS := payload.HandSize
	if handS < minHandSize || handS > maxHandSize {
		return ret, fmt.Errorf("Hand size needs to be a number between %d and %d (both included).", minHandSize, maxHandSize)
	}
	ret = append(ret, usecase.Game.Options().HandSize(handS))
	// RANDOM FIRST CZAR?
	if payload.RandomFirstCzar {
		ret = append(ret, usecase.Game.Options().RandomStartingCzar())
	}
	return ret, nil
}

func availableExpansions(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	_, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	exps := usecase.Card.AvailableExpansions()
	writeResponse(w, exps)
	return nil
}

// Utils

func gameFromRequest(req *http.Request) (cah.Game, error) {
	strID := mux.Vars(req)["gameID"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return cah.Game{}, err
	}
	g, err := usecase.Game.ByID(id)
	if err != nil {
		return g, fmt.Errorf("Could not get game with id %d", id)
	}
	return g, nil
}
