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
	stateStore := mem.GetGameStateStore()
	gameStore := mem.GetGameStore()
	cardStore := mem.GetCardStore()
	userStore := sqlite.NewUserStore()
	usecases := cah.Usecases{
		GameState: usecase.NewGameStateUsecase(stateStore),
		Card:      usecase.NewCardUsecase(cardStore),
		User:      usecase.NewUserUsecase(userStore),
		Game:      usecase.NewGameUsecase(gameStore),
	}
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
	// Start the 	usecase.Game.Create(users[2], "Finished", "")
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
	g, _ = usecase.Game.ByID(3)
	if err != nil {
		panic(err)
	}
	//winnerIndex := 0
	for i := range g.State.Players {
		if g.State.CurrCzarIndex == i {
			continue
		}
		g.State, _ = usecase.GameState.PlayRandomWhiteCards(i, g.State)
		//winnerIndex = i // last player to play its cards, we dont really care about the winner
	}
	//g.State, _ = usecase.GameState.GiveBlackCardToWinner(winnerIndex, g.State)
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
