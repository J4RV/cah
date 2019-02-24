package usecase

import (
	"testing"

	"github.com/j4rv/golang-stuff/cah"
	"github.com/j4rv/golang-stuff/cah/db/mem"
	"github.com/stretchr/testify/assert"
)

var control = stateController{}

func getStateUsecase() cah.GameStateUsecases {
	store := mem.GetGameStateStore()
	usecase := NewGameStateUsecase(store)
	return usecase
}

func TestNextCzar(t *testing.T) {
	assert := assert.New(t)
	s := getStateFixture()
	s.Phase = cah.CzarChoosingWinner
	assert.Equal(s.CurrCzarIndex, 0, "Unexpected first czar")

	s, err := control.nextCzar(s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 1, "Unexpected second czar")

	s, err = control.nextCzar(s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 2, "Unexpected third czar")

	s, err = control.nextCzar(s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 0, "Unexpected fourth czar")
}

func TestNextCzar_errors(t *testing.T) {
	assert := assert.New(t)
	s := getStateFixture()

	s.BlackCardInPlay = s.BlackDeck[0]
	s.Phase = cah.CzarChoosingWinner
	s, err := control.nextCzar(s)
	assert.NotEqual(err, nil, "Expected 'black card in play' error but found nil")

	s.Phase = cah.Finished
	s.BlackCardInPlay = nilBlackCard
	s, err = control.nextCzar(s)
	assert.NotEqual(err, nil, "Expected 'incorrect phase' error but found nil")
}
