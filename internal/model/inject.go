package cah

type DataStore struct {
	Game      GameStore
	GameState GameStateStore
	Card      CardStore
	User      UserStore
}

type Usecases struct {
	Game      GameUsecases
	GameState GameStateUsecases
	Card      CardUsecases
	User      UserUsecases
}
