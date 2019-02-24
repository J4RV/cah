package usecase

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type gameController struct {
	store   cah.GameStore
	options Options
}

func NewGameUsecase(store cah.GameStore) *gameController {
	return &gameController{
		store: store,
	}
}

func (control gameController) Create(owner cah.User, name, pass string) error {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return errors.New("A game name cannot be blank")
	}
	game := cah.Game{
		Owner:  owner,
		UserID: owner.ID,
		Name:   trimmed,
		Users:  []cah.User{owner},
	}
	trimmedPass := strings.TrimSpace(pass)
	if trimmedPass != "" {
		game.Password = trimmedPass
	}
	return control.store.Create(game)
}

func (control gameController) ByID(id int) (cah.Game, error) {
	return control.store.ByID(id)
}

func (control gameController) AllOpen() []cah.Game {
	return control.store.ByStatePhase(cah.NotStarted)
}

func (control gameController) InProgressForUser(user cah.User) []cah.Game {
	gamesInProgress := control.store.ByStatePhase(cah.SinnersPlaying, cah.CzarChoosingWinner)
	ret := []cah.Game{}
	for _, ipg := range gamesInProgress {
		for _, u := range ipg.Users {
			if u.ID == user.ID {
				ret = append(ret, ipg)
				break
			}
		}
	}
	return ret
}

func (control gameController) UserJoins(user cah.User, game cah.Game) error {
	log.Printf("User '%s' joins game '%s'\n", user.Username, game.Name)
	for _, u := range game.Users {
		if u.ID == user.ID {
			return nil // don't add the user if they already joined
		}
	}
	game.Users = append(game.Users, user)
	return control.store.Update(game)
}

func (control gameController) Start(g cah.Game, state cah.GameState, opts ...cah.Option) error {
	if len(g.Users) < 3 {
		return fmt.Errorf("The minimum amount of players to start a game is 3, got: %d", len(g.Users))
	}
	s := g.State
	if s.ID != 0 {
		return fmt.Errorf("Tried to start a game but it already has a state. State ID '%d'", s.ID)
	}
	players := make([]*cah.Player, len(g.Users))
	for i, u := range g.Users {
		players[i] = cah.NewPlayer(u)
	}
	state.Players = players
	applyOptions(&state, opts...)
	state, err := putBlackCardInPlay(state)
	if err != nil {
		return err
	}
	playersDraw(&state)
	g.State = state
	err = control.store.Update(g)
	if err != nil {
		return err
	}
	return nil
}

func (control gameController) Options() cah.GameOptions {
	return control.options
}
