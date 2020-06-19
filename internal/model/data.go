package cah

type DataStore struct {
	Game      GameStore
	GameState GameStateStore
	Card      CardStore
	User      UserStore
}
