package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	cah "github.com/j4rv/cah/internal/model"
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

func gamesPageHandler(user cah.User, w http.ResponseWriter, req *http.Request) error {
	inProgress, err := usecase.Game.InProgressForUser(user)
	if err != nil {
		return err
	}
	openGames, err := usecase.Game.AllOpen()
	if err != nil {
		return err
	}
	execTemplate(gamesPageTmpl, w, http.StatusOK, gamesPageCtx{
		LoggedUser:      user,
		InProgressGames: inProgress,
		OpenGames:       openGames,
	})
	return nil
}

type createGamePageCtx struct {
	LoggedUser cah.User
	Flashes    []interface{}
}

func createGamePageHandler(user cah.User, w http.ResponseWriter, req *http.Request) error {
	execTemplate(createGamePageTmpl, w, http.StatusOK, lobbyPageCtx{
		LoggedUser: user,
		Flashes:    getFlashes(w, req),
	})
	return nil
}

type lobbyPageCtx struct {
	LoggedUser          cah.User
	Game                cah.Game
	AvailableExpansions []string
	Flashes             []interface{}
}

func lobbyPageHandler(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	execTemplate(lobbyPageTmpl, w, http.StatusOK, lobbyPageCtx{
		LoggedUser:          user,
		Game:                game,
		AvailableExpansions: usecase.Card.AvailableExpansions(),
		Flashes:             getFlashes(w, req),
	})
	return nil
}

type ingamePageCtx struct {
	LoggedUser cah.User
	Game       cah.Game
	Flashes    []interface{}
}

func ingamePageHandler(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	if game.Finished {
		http.Redirect(w, req, fmt.Sprint("/games/", game.ID, "/final-stats"), http.StatusFound)
		return nil
	}
	execTemplate(ingamePageTmpl, w, http.StatusOK, lobbyPageCtx{
		LoggedUser: user,
		Game:       game,
		Flashes:    getFlashes(w, req),
	})
	return nil
}

func finalStatsPageHandler(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	execTemplate(finalStatsPageTmpl, w, http.StatusOK, nil)
	return nil
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

func lobbyState(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	writeJSONResponse(w, gameToResponse(game))
	return nil
}

func gameToResponse(g cah.Game) gameRoomResponse {
	players := make([]string, len(g.Users))
	for i := range g.Users {
		players[i] = g.Users[i].Username
	}
	state, err := usecase.GameState.ByID(g.StateID)
	phase := ""
	if err == nil {
		phase = state.Phase.String()
	}
	return gameRoomResponse{
		ID:          g.ID,
		Owner:       g.Owner.Username,
		Name:        g.Name,
		HasPassword: g.Password != "",
		Players:     players,
		Phase:       phase,
		StateID:     g.StateID,
	}
}

/*
CREATE GAME
*/

func createGame(user cah.User, w http.ResponseWriter, req *http.Request) error {
	req.ParseForm()
	name := req.Form["name"]
	password := optionalSingleFormParam(req.Form["password"])
	if err := requiredFormParams(name); err != nil {
		return err
	}
	g, err := usecase.Game.Create(user, name[0], password)
	if err != nil {
		return err
	}
	http.Redirect(w, req, fmt.Sprint("/games/", g.ID), http.StatusFound)
	return nil
}

/*
JOIN AND LEAVE GAME
*/

func joinGame(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	err := usecase.Game.UserJoins(user, game)
	if err != nil {
		return err
	}
	http.Redirect(w, req, fmt.Sprint("/games/", game.ID), http.StatusFound)
	return nil
}

func leaveGame(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	err := usecase.Game.UserLeaves(user, game)
	if err != nil {
		return err
	}
	http.Redirect(w, req, fmt.Sprint("/games/", game.ID), http.StatusFound)
	return nil
}

/*
START GAME
*/

func startGame(user cah.User, w http.ResponseWriter, req *http.Request) error {
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
	if g.Owner.ID != user.ID {
		return errors.New("Only the game owner can start the game")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	opts, err := optionsFromCreateRequest(
		expansions,
		handsize[0],
		maxrounds[0],
		len(randomfirstczar) > 0,
	)
	if err != nil {
		return err
	}
	state := usecase.GameState.Create()
	err = usecase.Game.Start(g, state, opts...)
	if err != nil {
		return err
	}
	log.Println("User", user.Username, "started the game:", g.Name)
	http.Redirect(w, req, fmt.Sprintf("/games/%d/ingame", g.ID), http.StatusFound)
	return nil
}

func optionsFromCreateRequest(expansions []string, handsize, maxrounds string, randomfirstczar bool) ([]cah.Option, error) {
	ret := []cah.Option{}

	// EXPANSIONS
	blacks := usecase.Card.BlacksByExpansion(expansions...)
	whites := usecase.Card.WhitesByExpansion(expansions...)
	if len(blacks) < minBlacks {
		return ret, fmt.Errorf("Not enough black cards to play a game. Please select more expansions. The amount of Black cards in selected expansions is %d, but the minimum is %d", len(blacks), minBlacks)
	}
	if len(whites) < minWhites {
		return ret, fmt.Errorf("Not enough white cards to play a game. Please select more expansions. The amount of White cards in selected expansions is %d, but the minimum is %d", len(whites), minWhites)
	}
	ret = append(ret, usecase.Game.Options().BlackDeck(blacks))
	ret = append(ret, usecase.Game.Options().WhiteDeck(whites))

	// HAND SIZE
	handS, err := strconv.Atoi(handsize)
	if err != nil {
		return ret, fmt.Errorf("Hand size must be an int")
	}
	if handS < minHandSize || handS > maxHandSize {
		return ret, fmt.Errorf("Hand size needs to be a number between %d and %d (both included)", minHandSize, maxHandSize)
	}
	ret = append(ret, usecase.Game.Options().HandSize(handS))

	// RANDOM FIRST CZAR?
	if randomfirstczar {
		ret = append(ret, usecase.Game.Options().RandomStartingCzar())
	}

	// MAX ROUNDS
	maxRounds, err := strconv.Atoi(maxrounds)
	if err != nil {
		return ret, fmt.Errorf("Hand size must be an int")
	}
	ret = append(ret, usecase.Game.Options().MaxRounds(maxRounds))

	return ret, nil
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
