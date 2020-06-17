package server

import (
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

func lobbyState(w http.ResponseWriter, req *http.Request) error {
	g, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	writeResponse(w, gameToResponse(g))
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
	if err := requiredFormParams(name); err != nil {
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
JOIN AND LEAVE GAME
*/

func joinGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	game, err := gameFromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = usecase.Game.UserJoins(u, game)
	if err != nil {
		return err
	}
	http.Redirect(w, req, fmt.Sprint("/games/", game.ID), http.StatusFound)
	return nil
}

func leaveGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	game, err := gameFromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = usecase.Game.UserLeaves(u, game)
	if err != nil {
		return err
	}
	http.Redirect(w, req, fmt.Sprint("/games/", game.ID), http.StatusFound)
	return nil
}

/*
START GAME
*/

func startGame(w http.ResponseWriter, req *http.Request) error {
	u, err := userFromSession(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	req.ParseForm()
	expansions := req.Form["expansion"]
	handsize := req.Form["handsize"]
	maxrounds := req.Form["maxrounds"]
	randomfirstczar := req.Form["randomfirstczar"]
	if err := requiredFormParams(expansions, handsize, maxrounds); err != nil {
		return err
	}
	g, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	if g.Owner != u {
		return errors.New("Only the game owner can start the game")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	opts, err := optionsFromCreateRequest(expansions, handsize, maxrounds, randomfirstczar)
	if err != nil {
		return err
	}
	state := usecase.GameState.Create()
	err = usecase.Game.Start(g, state, opts...)
	if err != nil {
		return err
	}
	log.Println("User", u.Username, "started the game:", g.Name)
	return nil
}

func optionsFromCreateRequest(expansions, handsize, maxrounds, randomfirstczar []string) ([]cah.Option, error) {
	ret := []cah.Option{}

	// EXPANSIONS
	blacks := usecase.Card.ExpansionBlacks(expansions...)
	whites := usecase.Card.ExpansionWhites(expansions...)
	if len(blacks) < minBlacks {
		return ret, fmt.Errorf("Not enough black cards to play a game. Please select more expansions. The amount of Black cards in selected expansions is %d, but the minimum is %d", len(blacks), minBlacks)
	}
	if len(whites) < minWhites {
		return ret, fmt.Errorf("Not enough white cards to play a game. Please select more expansions. The amount of White cards in selected expansions is %d, but the minimum is %d", len(whites), minWhites)
	}
	ret = append(ret, usecase.Game.Options().BlackDeck(blacks))
	ret = append(ret, usecase.Game.Options().WhiteDeck(whites))

	// HAND SIZE
	handS, err := strconv.Atoi(handsize[0])
	if err != nil {
		return ret, fmt.Errorf("Hand size must be an int")
	}
	if handS < minHandSize || handS > maxHandSize {
		return ret, fmt.Errorf("Hand size needs to be a number between %d and %d (both included)", minHandSize, maxHandSize)
	}
	ret = append(ret, usecase.Game.Options().HandSize(handS))

	// RANDOM FIRST CZAR?
	if len(randomfirstczar) > 0 {
		ret = append(ret, usecase.Game.Options().RandomStartingCzar())
	}

	// MAX ROUNDS
	maxRounds, err := strconv.Atoi(maxrounds[0])
	if err != nil {
		return ret, fmt.Errorf("Hand size must be an int")
	}
	ret = append(ret, usecase.Game.Options().MaxRounds(maxRounds))

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
