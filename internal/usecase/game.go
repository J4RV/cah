package usecase

import (
	"errors"
	"fmt"
	"log"
	"strings"

	cah "github.com/j4rv/cah/internal/model"
)

type gameController struct {
	store   cah.GameStore
	options Options
}

// NewGameUsecase returns a cah.GameUsecases
func NewGameUsecase(uc cah.Usecases, store cah.GameStore) cah.GameUsecases {
	return &gameController{
		store: store,
	}
}

func (control gameController) Create(owner cah.User, name, pass string) (cah.Game, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return cah.Game{}, errors.New("A game name cannot be blank")
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

func (control gameController) ByGameStateID(id int) (cah.Game, error) {
	return control.store.ByGameStateID(id)
}

func (control gameController) AllOpen() ([]cah.Game, error) {
	return control.store.ByPhase(false, false)
}

func (control gameController) InProgressForUser(user cah.User) ([]cah.Game, error) {
	gamesInProgress, err := control.store.ByPhase(true, false)
	if err != nil {
		return gamesInProgress, err
	}
	ret := []cah.Game{}
	for _, ipg := range gamesInProgress {
		for _, u := range ipg.Users {
			if u.ID == user.ID {
				ret = append(ret, ipg)
				break
			}
		}
	}
	return ret, err
}

func (control gameController) Update(g cah.Game) error {
	return control.store.Update(g)
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

func (control gameController) UserLeaves(user cah.User, game cah.Game) error {
	log.Printf("User '%s' leaves game '%s'\n", user.Username, game.Name)
	for i, u := range game.Users {
		if u.ID == user.ID {
			game.Users = append(game.Users[:i], game.Users[i+1:]...)
			return control.store.Update(game)
		}
	}
	return errors.New("User not in game: " + user.Username)
}

func (control gameController) Start(g cah.Game, state *cah.GameState, opts ...cah.Option) error {
	if len(g.Users) < 3 {
		return fmt.Errorf("The minimum amount of players to start a game is 3, got: %d", len(g.Users))
	}
	if g.Started {
		return fmt.Errorf("Tried to start a game but it already started")
	}
	players := make([]*cah.Player, len(g.Users))
	for i, u := range g.Users {
		players[i] = cah.NewPlayer(u)
	}
	state.Players = players
	applyOptions(state, opts...)
	g.StateID = state.ID
	err := putBlackCardInPlay(state)
	if err != nil {
		return err
	}
	playersDraw(state)
	g.Started = true
	err = control.store.Update(g)
	if err != nil {
		return err
	}
	return nil
}

func (control gameController) Options() cah.GameOptions {
	return control.options
}
