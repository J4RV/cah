package usecase

import (
	"testing"

	"github.com/j4rv/cah/internal/db/mem"
	cah "github.com/j4rv/cah/internal/model"
	"github.com/stretchr/testify/assert"
)

var control = stateController{}

func getStateUsecase() cah.GameStateUsecases {
	store := mem.GetGameStateStore(cah.DataStore{})
	usecase := NewGameStateUsecase(nil, store)
	return usecase
}

func TestNextCzar(t *testing.T) {
	assert := assert.New(t)
	s := getStateFixture()
	s.Phase = cah.CzarChoosingWinner
	assert.Equal(s.CurrCzarIndex, 0, "Unexpected first czar")

	err := control.nextCzar(&s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 1, "Unexpected second czar")

	err = control.nextCzar(&s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 2, "Unexpected third czar")

	err = control.nextCzar(&s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 0, "Unexpected fourth czar")
}

func TestNextCzar_errors(t *testing.T) {
	assert := assert.New(t)
	s := getStateFixture()

	s.BlackCardInPlay = s.BlackDeck[0]
	s.Phase = cah.CzarChoosingWinner
	err := control.nextCzar(&s)
	assert.NotEqual(err, nil, "Expected 'black card in play' error but found nil")

	s.Phase = cah.Finished
	err = control.nextCzar(&s)
	assert.NotEqual(err, nil, "Expected 'incorrect phase' error but found nil")
}
