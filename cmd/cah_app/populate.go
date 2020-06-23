package main

import cah "github.com/j4rv/cah/internal/model"

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
	wd := usecase.Card.WhitesByExpansion("Base UK")
	bd := usecase.Card.BlacksByExpansion("Base UK")
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
	wd = usecase.Card.WhitesByExpansion("Base UK")
	bd = usecase.Card.BlacksByExpansion("Base UK")
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
		u, _ := usecase.User.ByID(uint(i + 1))
		users[i] = u
	}
	return users
}
