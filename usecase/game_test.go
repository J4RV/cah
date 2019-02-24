package usecase

import (
	"github.com/j4rv/golang-stuff/cah"
	"github.com/j4rv/golang-stuff/cah/db/mem"
)

func getGameUsecase() cah.GameUsecases {
	store := mem.GetGameStore()
	return NewGameUsecase(store)
}
