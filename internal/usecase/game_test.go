package usecase

import (
	"github.com/j4rv/cah/internal/db/mem"
	cah "github.com/j4rv/cah/internal/model"
)

func getGameUsecase() cah.GameUsecases {
	store := mem.GetGameStore(cah.DataStore{})
	return NewGameUsecase(cah.Usecases{}, store)
}
