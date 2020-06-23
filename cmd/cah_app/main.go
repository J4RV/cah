package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/j4rv/cah/internal/db/mem"
	"github.com/j4rv/cah/internal/db/sqlite"
	cah "github.com/j4rv/cah/internal/model"
	"github.com/j4rv/cah/internal/server"
	"github.com/j4rv/cah/internal/usecase"
	"github.com/j4rv/cah/internal/usecase/fixture"
	"gopkg.in/yaml.v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var cfg cah.Config
	cfgData, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("Config file could not be read: %v", err)
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		log.Fatalf("Could not unmarshal YAML config file: %v", err)
	}

	run(cfg)
}

func run(cfg cah.Config) {
	sqlite.InitDB(cfg.SQLiteDBPath)

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

	loadCards(cfg.ExpansionsPath, usecases.Card)

	fixture.PopulateUsers(usecases.User)
	createTestGames(usecases)

	server.Start(cfg, usecases)
}

func loadCards(expsPath string, cardUC cah.CardUsecases) {
	files, err := ioutil.ReadDir(expsPath)
	if err != nil {
		log.Fatal("error while loading expansions from path: "+expsPath, err)
	}
	for _, f := range files {
		if f.IsDir() {
			log.Println("Loading cards from", f.Name())
			err := cardUC.CreateFromFolder(expsPath+f.Name(), f.Name())
			if err != nil {
				fmt.Println("Got error while loading cards", err)
			}
		}
	}
}
