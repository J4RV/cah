package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/j4rv/cah"
	"github.com/j4rv/cah/db/mem"
	"github.com/j4rv/cah/db/sqlite"
	"github.com/j4rv/cah/server"
	"github.com/j4rv/cah/usecase"
	"github.com/j4rv/cah/usecase/fixture"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	run()
}

func run() {
	printRunningDir()
	sqlite.InitDB("db/database.sqlite3")

	dataStore := cah.DataStore{}
	dataStore.GameState = mem.GetGameStateStore(dataStore)
	dataStore.Game = mem.GetGameStore(dataStore)
	dataStore.Card = mem.GetCardStore(dataStore)
	dataStore.User = sqlite.NewUserStore(dataStore)

	usecases := cah.Usecases{}
	usecases.GameState = usecase.NewGameStateUsecase(usecases, dataStore.GameState)
	usecases.Game = usecase.NewGameUsecase(usecases, dataStore.Game)
	usecases.Card = usecase.NewCardUsecase(usecases, dataStore.Card)
	usecases.User = usecase.NewUserUsecase(usecases, dataStore.User)

	populateCards(usecases.Card)

	fixture.PopulateUsers(usecases.User)
	createTestGames(usecases)

	server.Start(usecases)
}

func printRunningDir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WOKDING DIR", dir)
}

// For quick prototyping

func createTestGames(usecase cah.Usecases) {
	users := getTestUsers(usecase)
	usecase.Game.Create(users[1], "A long and descriptive game name", "")
	usecase.Game.Create(users[0], "Amo a juga", "1234")
	usecase.Game.Create(users[0], "Almost finished", "")

	// Start the Amo a juga game
	g, _ := usecase.Game.ByID(2)
	usecase.Game.UserJoins(users[1], g)
	g, _ = usecase.Game.ByID(2)
	usecase.Game.UserJoins(users[2], g)
	g, _ = usecase.Game.ByID(2)
	wd := usecase.Card.ExpansionWhites("Base UK")
	bd := usecase.Card.ExpansionBlacks("Base UK")
	state := usecase.GameState.Create()
	err := usecase.Game.Start(g, state,
		usecase.Game.Options().BlackDeck(bd),
		usecase.Game.Options().WhiteDeck(wd),
	)
	if err != nil {
		panic(err)
	}

	// Start the "Almost finished"
	g, _ = usecase.Game.ByID(3)
	usecase.Game.UserJoins(users[1], g)
	g, _ = usecase.Game.ByID(3)
	usecase.Game.UserJoins(users[2], g)
	g, _ = usecase.Game.ByID(3)
	wd = usecase.Card.ExpansionWhites("Base UK")
	bd = usecase.Card.ExpansionBlacks("Base UK")
	state = usecase.GameState.Create()
	err = usecase.Game.Start(g, state,
		usecase.Game.Options().BlackDeck(bd),
		usecase.Game.Options().WhiteDeck(wd),
		usecase.Game.Options().MaxRounds(1),
	)
	if err != nil {
		panic(err)
	}
	g, _ = usecase.Game.ByID(3)
	if err != nil {
		panic(err)
	}
	s, err := usecase.GameState.ByID(g.StateID)
	if err != nil {
		panic(err)
	}
	for i := range s.Players {
		if s.CurrCzarIndex == i {
			continue
		}
		err := usecase.GameState.PlayRandomWhiteCards(i, s)
		if err != nil {
			panic(err)
		}
	}
}

func getTestUsers(usecase cah.Usecases) []cah.User {
	users := make([]cah.User, 4)
	for i := 0; i < 4; i++ {
		u, _ := usecase.User.ByID(i + 1)
		users[i] = u
	}
	return users
}

func populateCards(cardUC cah.CardUsecases) {
	files, err := ioutil.ReadDir("expansions")
	if err != nil {
		log.Fatal("error while populating cards. Is there an expansions folder in the active dir?", err)
	}
	for _, f := range files {
		if f.IsDir() {
			log.Println("Loading cards from", f.Name())
			err := cardUC.CreateFromFolder("expansions/"+f.Name(), f.Name())
			if err != nil {
				fmt.Println("Got error while loading cards", err)
			}
		}
	}
}
