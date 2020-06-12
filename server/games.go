package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/cah"
)

const minWhites = 34
const minBlacks = 8
const minHandSize = 5
const maxHandSize = 30
const gameDoesntExistMsg = "That game does not exist."
const gamesFlashKey = "games-flash"

/*
	TEMPLATE HANDLERS
*/

type gamesPageCtx struct {
	LoggedUser      cah.User
	InProgressGames []cah.Game
	OpenGames       []cah.Game
}

func gamesPageHandler(w http.ResponseWriter, req *http.Request) {
	user, err := userFromSession(w, req)
	if err != nil {
		addFlashMsg(notLoggedInMsg, loginFlashKey, w, req)
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}
	execTemplate(gamesPageTmpl, w, gamesPageCtx{
		LoggedUser:      user,
		InProgressGames: usecase.Game.InProgressForUser(user),
		OpenGames:       usecase.Game.AllOpen(),
	})
}

type createGamePageCtx struct {
	LoggedUser cah.User
	Flashes    []interface{}
}

func createGamePageHandler(w http.ResponseWriter, req *http.Request) {
	user, err := userFromSession(w, req)
	if err != nil {
		addFlashMsg(notLoggedInMsg, loginFlashKey, w, req)
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}
	execTemplate(createGamePageTmpl, w, lobbyPageCtx{
		LoggedUser: user,
		Flashes:    getFlashes(gamesFlashKey, w, req),
	})
}

type lobbyPageCtx struct {
	LoggedUser          cah.User
	Game                cah.Game
	AvailableExpansions []string
	Flashes             []interface{}
}

func lobbyPageHandler(w http.ResponseWriter, req *http.Request) {
	user, err := userFromSession(w, req)
	if err != nil {
		addFlashMsg(notLoggedInMsg, loginFlashKey, w, req)
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}
	game, err := gameFromRequest(req)
	if err != nil {
		addFlashMsg(gameDoesntExistMsg, gamesFlashKey, w, req)
		http.Redirect(w, req, "/games", http.StatusFound)
		return
	}
	execTemplate(lobbyPageTmpl, w, lobbyPageCtx{
		LoggedUser:          user,
		Game:                game,
		AvailableExpansions: usecase.Card.AvailableExpansions(),
		Flashes:             getFlashes(gamesFlashKey, w, req),
	})
}

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
	u, err := userFromSession(w, req)
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

func createGame(w http.ResponseWriter, req *http.Request) error {
	req.ParseForm()
	name := req.Form["name"]
	password := optionalSingleFormParam(req.Form["password"])
	if err := requiredSingleFormParams(name); err != nil {
		return err
	}
	// User is logged
	u, err := userFromSession(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	g, err := usecase.Game.Create(u, name[0], password)
	if err != nil {
		return err
	}
	http.Redirect(w, req, fmt.Sprint("/games/", g.ID), http.StatusFound)
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
	u, err := userFromSession(w, req)
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
	MaxRounds       int      `json:"maxRounds"`
}

func startGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(w, req)
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
	ret = append(ret, usecase.Game.Options().MaxRounds(payload.MaxRounds))
	return ret, nil
}

func availableExpansions(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	_, err := userFromSession(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	exps := usecase.Card.AvailableExpansions()
	sort.Strings(exps)
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
